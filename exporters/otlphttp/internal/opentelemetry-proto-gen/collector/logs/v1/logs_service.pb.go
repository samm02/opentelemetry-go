// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/logs/v1/logs_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "go.opentelemetry.io/otel/exporters/otlphttp/internal/opentelemetry-proto-gen/logs/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExportLogsServiceRequest struct {
	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceLogs         []*v1.ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resource_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExportLogsServiceRequest) Reset()         { *m = ExportLogsServiceRequest{} }
func (m *ExportLogsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceRequest) ProtoMessage()    {}
func (*ExportLogsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{0}
}
func (m *ExportLogsServiceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportLogsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportLogsServiceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportLogsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceRequest.Merge(m, src)
}
func (m *ExportLogsServiceRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExportLogsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceRequest proto.InternalMessageInfo

func (m *ExportLogsServiceRequest) GetResourceLogs() []*v1.ResourceLogs {
	if m != nil {
		return m.ResourceLogs
	}
	return nil
}

type ExportLogsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportLogsServiceResponse) Reset()         { *m = ExportLogsServiceResponse{} }
func (m *ExportLogsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceResponse) ProtoMessage()    {}
func (*ExportLogsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{1}
}
func (m *ExportLogsServiceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportLogsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportLogsServiceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportLogsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceResponse.Merge(m, src)
}
func (m *ExportLogsServiceResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExportLogsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportLogsServiceRequest)(nil), "opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest")
	proto.RegisterType((*ExportLogsServiceResponse)(nil), "opentelemetry.proto.collector.logs.v1.ExportLogsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/logs/v1/logs_service.proto", fileDescriptor_8e3bf87aaa43acd4)
}

var fileDescriptor_8e3bf87aaa43acd4 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x0d, 0x42, 0x17, 0xa9, 0x82, 0xcc, 0xaa, 0x56, 0x18, 0x64, 0x40, 0xa9, 0x8b, 0x26,
	0xb4, 0x6e, 0xdc, 0x29, 0x05, 0x77, 0x22, 0x65, 0xdc, 0x75, 0x53, 0x74, 0x78, 0x0c, 0x23, 0x31,
	0x2f, 0xbe, 0xa4, 0x83, 0x1e, 0xc2, 0x23, 0xb8, 0xf5, 0x2c, 0x2e, 0x3d, 0x82, 0xcc, 0x49, 0x64,
	0x92, 0x22, 0x53, 0x1d, 0x61, 0x70, 0x15, 0xf2, 0xf2, 0x7f, 0xff, 0xff, 0x3f, 0x08, 0x3f, 0x43,
	0x03, 0xda, 0x81, 0x82, 0x07, 0x70, 0xf4, 0x2c, 0x0d, 0xa1, 0x43, 0x99, 0xa1, 0x52, 0x90, 0x39,
	0x24, 0xa9, 0x30, 0xb7, 0xb2, 0x9c, 0xf8, 0x73, 0x69, 0x81, 0xca, 0x22, 0x03, 0xe1, 0x45, 0xd1,
	0xd1, 0x06, 0x19, 0x86, 0xe2, 0x9b, 0x14, 0x35, 0x21, 0xca, 0xc9, 0xf0, 0xb8, 0x2d, 0xa0, 0x69,
	0x1b, 0xc8, 0xe4, 0x9e, 0x0f, 0x2e, 0x9f, 0x0c, 0x92, 0xbb, 0xc2, 0xdc, 0xde, 0x84, 0xa4, 0x14,
	0x1e, 0x57, 0x60, 0x5d, 0x74, 0xcd, 0x77, 0x09, 0x2c, 0xae, 0x28, 0x83, 0x65, 0x8d, 0x0c, 0xd8,
	0xe1, 0xf6, 0xa8, 0x3f, 0x3d, 0x11, 0x6d, 0x15, 0xd6, 0xc1, 0x22, 0x5d, 0x13, 0xb5, 0x5f, 0xba,
	0x43, 0x8d, 0x5b, 0x72, 0xc0, 0xf7, 0x5b, 0xb2, 0xac, 0x41, 0x6d, 0x61, 0xfa, 0xca, 0x78, 0xbf,
	0x31, 0x8f, 0x5e, 0x18, 0xef, 0x05, 0x75, 0x74, 0x2e, 0x3a, 0xed, 0x2c, 0xfe, 0x5a, 0x64, 0x78,
	0xf1, 0x7f, 0x83, 0xd0, 0x2e, 0xd9, 0x9a, 0xbd, 0xb1, 0xf7, 0x2a, 0x66, 0x1f, 0x55, 0xcc, 0x3e,
	0xab, 0x98, 0xf1, 0x51, 0x81, 0xdd, 0x4c, 0x67, 0x7b, 0x0d, 0xbf, 0x79, 0xad, 0x99, 0xb3, 0xc5,
	0x22, 0xff, 0x49, 0x17, 0x28, 0xd1, 0x81, 0x92, 0xe0, 0x2b, 0x00, 0x59, 0x89, 0x4e, 0x19, 0x59,
	0x68, 0x07, 0xa4, 0x6f, 0x95, 0xdc, 0x50, 0x8f, 0x7d, 0xd6, 0x38, 0x07, 0xfd, 0xfb, 0xcf, 0xdc,
	0xf5, 0xfc, 0xe3, 0xe9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x1c, 0xa2, 0x18, 0x63, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogsServiceClient is the client API for LogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error)
}

type logsServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogsServiceClient(cc *grpc.ClientConn) LogsServiceClient {
	return &logsServiceClient{cc}
}

func (c *logsServiceClient) Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error) {
	out := new(ExportLogsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.logs.v1.LogsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServiceServer is the server API for LogsService service.
type LogsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error)
}

// UnimplementedLogsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogsServiceServer struct {
}

func (*UnimplementedLogsServiceServer) Export(ctx context.Context, req *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterLogsServiceServer(s *grpc.Server, srv LogsServiceServer) {
	s.RegisterService(&_LogsService_serviceDesc, srv)
}

func _LogsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportLogsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.logs.v1.LogsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Export(ctx, req.(*ExportLogsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.logs.v1.LogsService",
	HandlerType: (*LogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _LogsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/logs/v1/logs_service.proto",
}

func (m *ExportLogsServiceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportLogsServiceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportLogsServiceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResourceLogs) > 0 {
		for iNdEx := len(m.ResourceLogs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceLogs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLogsService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExportLogsServiceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportLogsServiceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportLogsServiceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogsService(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogsService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExportLogsServiceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ResourceLogs) > 0 {
		for _, e := range m.ResourceLogs {
			l = e.Size()
			n += 1 + l + sovLogsService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportLogsServiceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLogsService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogsService(x uint64) (n int) {
	return sovLogsService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExportLogsServiceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportLogsServiceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportLogsServiceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogsService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogsService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceLogs = append(m.ResourceLogs, &v1.ResourceLogs{})
			if err := m.ResourceLogs[len(m.ResourceLogs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogsService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogsService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportLogsServiceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportLogsServiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportLogsServiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogsService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogsService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogsService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogsService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogsService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogsService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLogsService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogsService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogsService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogsService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogsService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogsService = fmt.Errorf("proto: unexpected end of group")
)
