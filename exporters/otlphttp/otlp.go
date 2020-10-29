// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otlphttp

// This code was based on
// contrib.go.opencensus.io/exporter/ocagent/connection.go

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"unsafe"

	"google.golang.org/grpc/metadata"

	"go.opentelemetry.io/otel"
	colmetricpb "go.opentelemetry.io/otel/exporters/otlphttp/internal/opentelemetry-proto-gen/collector/metrics/v1"
	coltracepb "go.opentelemetry.io/otel/exporters/otlphttp/internal/opentelemetry-proto-gen/collector/trace/v1"

	"go.opentelemetry.io/otel/exporters/otlphttp/internal/transform"
	metricsdk "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/export/metric/aggregation"
	tracesdk "go.opentelemetry.io/otel/sdk/export/trace"
)

// Exporter is an OpenTelemetry exporter. It exports both traces and metrics
// from OpenTelemetry instrumented to code using OpenTelemetry protocol
// buffers to a configurable receiver.
type Exporter struct {
	// mu protects the non-atomic and non-channel variables
	mu sync.RWMutex
	// senderMu protects the concurrent unsafe sends on the shared client connection.
	senderMu      sync.Mutex
	started       bool
	traceExporter coltracepb.TraceServiceClient // TODO Sam fix - still has references to gRPC within
	//metricExporter colmetricpb.MetricsServiceClient // TODO Sam remove not needed..?
	// can't use ^ exporters, need to define our own

	// grpcClientConn    *grpc.ClientConn
	httpClient        *http.Client // using this rather than gRPC as above
	lastConnectErrPtr unsafe.Pointer

	startOnce      sync.Once
	stopCh         chan bool
	disconnectedCh chan bool

	backgroundConnectionDoneCh chan bool

	c        config
	metadata metadata.MD
}

var _ tracesdk.SpanExporter = (*Exporter)(nil)
var _ metricsdk.Exporter = (*Exporter)(nil) // TODO Sam not sure why this is erroring, is this a compile time assertion?

// newConfig initializes a config struct with default values and applies
// any ExporterOptions provided.
func newConfig(opts ...ExporterOption) config {
	cfg := config{
		numWorkers: DefaultNumWorkers, // default = 1
		// grpcServiceConfig: DefaultGRPCServiceConfig,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

// NewExporter constructs a new Exporter and starts it.
func NewExporter(opts ...ExporterOption) (*Exporter, error) {
	exp := NewUnstartedExporter(opts...)
	if err := exp.Start(); err != nil {
		return nil, err
	}
	return exp, nil
}

// NewUnstartedExporter constructs a new Exporter and does not start it.
func NewUnstartedExporter(opts ...ExporterOption) *Exporter {
	e := new(Exporter)
	e.c = newConfig(opts...)
	if len(e.c.headers) > 0 {
		e.metadata = metadata.New(e.c.headers)
	}

	// TODO (rghetia): add resources

	return e
}

var (
	errAlreadyStarted  = errors.New("already started")
	errNotStarted      = errors.New("not started")
	errDisconnected    = errors.New("exporter disconnected")
	errStopped         = errors.New("exporter stopped")
	errContextCanceled = errors.New("context canceled")
)

// Start dials to the collector, establishing a connection to it. It also
// initiates the Config and Trace services by sending over the initial
// messages that consist of the node identifier. Start invokes a background
// connector that will reattempt connections to the collector periodically
// if the connection dies.
func (e *Exporter) Start() error {
	var err = errAlreadyStarted
	e.startOnce.Do(func() {
		e.mu.Lock()
		e.started = true
		e.disconnectedCh = make(chan bool, 1)
		e.stopCh = make(chan bool)
		e.backgroundConnectionDoneCh = make(chan bool)
		e.mu.Unlock()

		// An optimistic first connection attempt to ensure that
		// applications under heavy load can immediately process
		// data. See https://github.com/census-ecosystem/opencensus-go-exporter-ocagent/pull/63
		if err := e.connect(); err == nil {
			e.setStateConnected()
		} else {
			e.setStateDisconnected(err)
		}
		go e.indefiniteBackgroundConnection()

		err = nil
	})

	return err
}

func (e *Exporter) prepareCollectorAddress() string {
	if e.c.collectorAddr != "" {
		return e.c.collectorAddr
	}
	return fmt.Sprintf("%s:%d", DefaultCollectorHost, DefaultCollectorPort)
}

func (e *Exporter) enableConnections(c *http.Client) error {
	e.mu.RLock()
	started := e.started
	e.mu.RUnlock()

	if !started {
		return errNotStarted
	}

	e.mu.Lock()
	// If previous clientConn is same as the current then just return.
	// This doesn't happen right now as this func is only called with new ClientConn.
	// It is more about future-proofing.
	// if e.grpcClientConn == cc {
	// 	e.mu.Unlock()
	// 	return nil
	// }
	if e.httpClient == c {
		e.mu.Unlock()
		return nil
	}
	// // If the previous clientConn was non-nil, close it
	// if e.grpcClientConn != nil {
	// 	_ = e.grpcClientConn.Close()
	// }
	// e.grpcClientConn = cc
	e.httpClient = c

	// set up traceExporter and metric Exporter, that are
	// compatible with http clients! TODO SAM
	// e.traceExporter = coltracepb.NewTraceServiceClient(c)
	// e.metricExporter = colmetricpb.NewMetricsServiceClient(c)

	e.mu.Unlock()

	return nil
}

func (e *Exporter) contextWithMetadata(ctx context.Context) context.Context {
	if e.metadata.Len() > 0 {
		return metadata.NewOutgoingContext(ctx, e.metadata)
	}
	return ctx
}

// func (e *Exporter) dialToCollector() (*grpc.ClientConn, error) {
// 	addr := e.prepareCollectorAddress()

// 	dialOpts := []grpc.DialOption{}
// 	if e.c.grpcServiceConfig != "" {
// 		dialOpts = append(dialOpts, grpc.WithDefaultServiceConfig(e.c.grpcServiceConfig))
// 	}
// 	if e.c.clientCredentials != nil {
// 		dialOpts = append(dialOpts, grpc.WithTransportCredentials(e.c.clientCredentials))
// 	} else if e.c.canDialInsecure {
// 		dialOpts = append(dialOpts, grpc.WithInsecure())
// 	}
// 	if e.c.compressor != "" {
// 		dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(grpc.UseCompressor(e.c.compressor)))
// 	}
// 	if len(e.c.grpcDialOptions) != 0 {
// 		dialOpts = append(dialOpts, e.c.grpcDialOptions...)
// 	}

// 	ctx := e.contextWithMetadata(context.Background())
// 	return grpc.DialContext(ctx, addr, dialOpts...)
// }

func (e *Exporter) dialToCollector() (*http.Client, error) {
	//addr := e.prepareCollectorAddress() TODO Sam -- use address to set up network connection each time
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial:                (&net.Dialer{Timeout: 5 * time.Second}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}, nil
}

// closeStopCh is used to wrap the exporters stopCh channel closing for testing.
var closeStopCh = func(stopCh chan bool) {
	close(stopCh)
}

// Shutdown closes all connections and releases resources currently being used
// by the exporter. If the exporter is not started this does nothing.
func (e *Exporter) Shutdown(ctx context.Context) error {
	e.mu.RLock()
	c := e.httpClient
	started := e.started
	e.mu.RUnlock()

	if !started {
		return nil
	}

	var err error
	if c != nil {
		// Clean things up before checking this error.
		//err = c.Close() TODO SAM close the connections
	}

	// At this point we can change the state variable started
	e.mu.Lock()
	e.started = false
	e.mu.Unlock()
	closeStopCh(e.stopCh)

	// Ensure that the backgroundConnector returns
	select {
	case <-e.backgroundConnectionDoneCh:
	case <-ctx.Done():
		return ctx.Err()
	}

	return err
}

// Export implements the "go.opentelemetry.io/otel/sdk/export/metric".Exporter
// interface. It transforms and batches metric Records into OTLP Metrics and
// transmits them to the configured collector.
// func (e *Exporter) Export(parent context.Context, cps metricsdk.CheckpointSet) error {
// 	// Unify the parent context Done signal with the exporter stopCh.
// 	ctx, cancel := context.WithCancel(parent)
// 	defer cancel()
// 	go func(ctx context.Context, cancel context.CancelFunc) {
// 		select {
// 		case <-ctx.Done():
// 		case <-e.stopCh:
// 			cancel()
// 		}
// 	}(ctx, cancel)

// 	rms, err := transform.CheckpointSet(ctx, e, cps, e.c.numWorkers)
// 	if err != nil {
// 		return err
// 	}

// 	if !e.connected() {
// 		return errDisconnected
// 	}

// 	select {
// 	case <-e.stopCh:
// 		return errStopped
// 	case <-ctx.Done():
// 		return errContextCanceled
// 	default:
// 		e.senderMu.Lock()
// 		_, err := e.metricExporter.Export(e.contextWithMetadata(ctx), &colmetricpb.ExportMetricsServiceRequest{
// 			ResourceMetrics: rms,
// 		})
// 		e.senderMu.Unlock()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// Export implements the "go.opentelemetry.io/otel/sdk/export/metric".Exporter
// interface. It transforms and batches metric Records into OTLP Metrics and
// transmits them to the configured collector.
func (e *Exporter) Export(parent context.Context, cps metricsdk.CheckpointSet) error {
	// Unify the parent context Done signal with the exporter stopCh.
	ctx, cancel := context.WithCancel(parent)
	defer cancel()
	go func(ctx context.Context, cancel context.CancelFunc) {
		select {
		case <-ctx.Done():
		case <-e.stopCh:
			cancel()
		}
	}(ctx, cancel)

	rms, err := transform.CheckpointSet(ctx, e, cps, e.c.numWorkers)
	if err != nil {
		return err
	}

	if !e.connected() {
		return errDisconnected
	}

	select {
	case <-e.stopCh:
		return errStopped
	case <-ctx.Done():
		return errContextCanceled
	default:
		e.senderMu.Lock()

		msr := &colmetricpb.ExportMetricsServiceRequest{
			ResourceMetrics: rms,
		}

		msrbytes, err := msr.Marshal()
		if err != nil {
			fmt.Errorf("error marshalling request: %w", err)
			return err
		}

		//ctxmtd := e.contextWithMetadata(ctx)

		// create a client that exports using msr and ctxmtd
		// can i invoke the http client right here??? and send those pieces of data as proto-bytes
		//_, err := e.metricExporter.Export(ctxmtd, msr)

		req, err := http.NewRequest(http.MethodPost, e.c.collectorAddr, bytes.NewReader(msrbytes))
		if err != nil {
			fmt.Errorf("error creating request: %w", err)
			return err
		}
		req.Header.Set("Content-Type", "application/x-protobuf")

		resp, err := e.httpClient.Do(req)
		if err != nil {
			fmt.Errorf("failed to send metrics to ingestion %w", e)
			return err
		}

		fmt.Print("obsvs-ingestion/metrics response: %v\n", resp.Status)

		// return nil
		// in metrics_service.pb.go
		// func (c *metricsServiceClient) Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error) {
		// 	out := new(ExportMetricsServiceResponse)
		// 	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.metrics.v1.MetricsService/Export", in, out, opts...)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	return out, nil
		// }
		// create own exporter and do the export! TODO Sam

		e.senderMu.Unlock()
		if err != nil {
			return err
		}
	}
	return nil
}

// ExportKindFor reports back to the OpenTelemetry SDK sending this Exporter
// metric telemetry that it needs to be provided in a pass-through format.
func (e *Exporter) ExportKindFor(*otel.Descriptor, aggregation.Kind) metricsdk.ExportKind {
	return metricsdk.PassThroughExporter
}

// ExportSpans exports a batch of SpanData.
func (e *Exporter) ExportSpans(ctx context.Context, sds []*tracesdk.SpanData) error {
	return e.uploadTraces(ctx, sds)
}

func (e *Exporter) uploadTraces(ctx context.Context, sdl []*tracesdk.SpanData) error {
	select {
	case <-e.stopCh:
		return nil
	default:
		if !e.connected() {
			return nil
		}

		protoSpans := transform.SpanData(sdl)
		if len(protoSpans) == 0 {
			return nil
		}

		e.senderMu.Lock()
		// TODO Sam - fix this too
		_, err := e.traceExporter.Export(e.contextWithMetadata(ctx), &coltracepb.ExportTraceServiceRequest{
			ResourceSpans: protoSpans,
		})
		e.senderMu.Unlock()
		if err != nil {
			e.setStateDisconnected(err)
			return err
		}
	}
	return nil
}
