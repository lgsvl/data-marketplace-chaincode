// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/conversion_action_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [ConversionActionService.GetConversionAction].
type GetConversionActionRequest struct {
	// The resource name of the conversion action to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConversionActionRequest) Reset()         { *m = GetConversionActionRequest{} }
func (m *GetConversionActionRequest) String() string { return proto.CompactTextString(m) }
func (*GetConversionActionRequest) ProtoMessage()    {}
func (*GetConversionActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_service_a3114b9b067f0aff, []int{0}
}
func (m *GetConversionActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConversionActionRequest.Unmarshal(m, b)
}
func (m *GetConversionActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConversionActionRequest.Marshal(b, m, deterministic)
}
func (dst *GetConversionActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConversionActionRequest.Merge(dst, src)
}
func (m *GetConversionActionRequest) XXX_Size() int {
	return xxx_messageInfo_GetConversionActionRequest.Size(m)
}
func (m *GetConversionActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConversionActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConversionActionRequest proto.InternalMessageInfo

func (m *GetConversionActionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [ConversionActionService.MutateConversionActions].
type MutateConversionActionsRequest struct {
	// The ID of the customer whose conversion actions are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual conversion actions.
	Operations []*ConversionActionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateConversionActionsRequest) Reset()         { *m = MutateConversionActionsRequest{} }
func (m *MutateConversionActionsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateConversionActionsRequest) ProtoMessage()    {}
func (*MutateConversionActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_service_a3114b9b067f0aff, []int{1}
}
func (m *MutateConversionActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateConversionActionsRequest.Unmarshal(m, b)
}
func (m *MutateConversionActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateConversionActionsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateConversionActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateConversionActionsRequest.Merge(dst, src)
}
func (m *MutateConversionActionsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateConversionActionsRequest.Size(m)
}
func (m *MutateConversionActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateConversionActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateConversionActionsRequest proto.InternalMessageInfo

func (m *MutateConversionActionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateConversionActionsRequest) GetOperations() []*ConversionActionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateConversionActionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateConversionActionsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a conversion action.
type ConversionActionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*ConversionActionOperation_Create
	//	*ConversionActionOperation_Update
	//	*ConversionActionOperation_Remove
	Operation            isConversionActionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ConversionActionOperation) Reset()         { *m = ConversionActionOperation{} }
func (m *ConversionActionOperation) String() string { return proto.CompactTextString(m) }
func (*ConversionActionOperation) ProtoMessage()    {}
func (*ConversionActionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_service_a3114b9b067f0aff, []int{2}
}
func (m *ConversionActionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionOperation.Unmarshal(m, b)
}
func (m *ConversionActionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionOperation.Marshal(b, m, deterministic)
}
func (dst *ConversionActionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionOperation.Merge(dst, src)
}
func (m *ConversionActionOperation) XXX_Size() int {
	return xxx_messageInfo_ConversionActionOperation.Size(m)
}
func (m *ConversionActionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionOperation proto.InternalMessageInfo

func (m *ConversionActionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isConversionActionOperation_Operation interface {
	isConversionActionOperation_Operation()
}

type ConversionActionOperation_Create struct {
	Create *resources.ConversionAction `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type ConversionActionOperation_Update struct {
	Update *resources.ConversionAction `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type ConversionActionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*ConversionActionOperation_Create) isConversionActionOperation_Operation() {}

func (*ConversionActionOperation_Update) isConversionActionOperation_Operation() {}

func (*ConversionActionOperation_Remove) isConversionActionOperation_Operation() {}

func (m *ConversionActionOperation) GetOperation() isConversionActionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *ConversionActionOperation) GetCreate() *resources.ConversionAction {
	if x, ok := m.GetOperation().(*ConversionActionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ConversionActionOperation) GetUpdate() *resources.ConversionAction {
	if x, ok := m.GetOperation().(*ConversionActionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ConversionActionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*ConversionActionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConversionActionOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConversionActionOperation_OneofMarshaler, _ConversionActionOperation_OneofUnmarshaler, _ConversionActionOperation_OneofSizer, []interface{}{
		(*ConversionActionOperation_Create)(nil),
		(*ConversionActionOperation_Update)(nil),
		(*ConversionActionOperation_Remove)(nil),
	}
}

func _ConversionActionOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConversionActionOperation)
	// operation
	switch x := m.Operation.(type) {
	case *ConversionActionOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *ConversionActionOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *ConversionActionOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("ConversionActionOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _ConversionActionOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConversionActionOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.ConversionAction)
		err := b.DecodeMessage(msg)
		m.Operation = &ConversionActionOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.ConversionAction)
		err := b.DecodeMessage(msg)
		m.Operation = &ConversionActionOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &ConversionActionOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _ConversionActionOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConversionActionOperation)
	// operation
	switch x := m.Operation.(type) {
	case *ConversionActionOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConversionActionOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConversionActionOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for conversion action mutate.
type MutateConversionActionsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateConversionActionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateConversionActionsResponse) Reset()         { *m = MutateConversionActionsResponse{} }
func (m *MutateConversionActionsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateConversionActionsResponse) ProtoMessage()    {}
func (*MutateConversionActionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_service_a3114b9b067f0aff, []int{3}
}
func (m *MutateConversionActionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateConversionActionsResponse.Unmarshal(m, b)
}
func (m *MutateConversionActionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateConversionActionsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateConversionActionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateConversionActionsResponse.Merge(dst, src)
}
func (m *MutateConversionActionsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateConversionActionsResponse.Size(m)
}
func (m *MutateConversionActionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateConversionActionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateConversionActionsResponse proto.InternalMessageInfo

func (m *MutateConversionActionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateConversionActionsResponse) GetResults() []*MutateConversionActionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the conversion action mutate.
type MutateConversionActionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateConversionActionResult) Reset()         { *m = MutateConversionActionResult{} }
func (m *MutateConversionActionResult) String() string { return proto.CompactTextString(m) }
func (*MutateConversionActionResult) ProtoMessage()    {}
func (*MutateConversionActionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_service_a3114b9b067f0aff, []int{4}
}
func (m *MutateConversionActionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateConversionActionResult.Unmarshal(m, b)
}
func (m *MutateConversionActionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateConversionActionResult.Marshal(b, m, deterministic)
}
func (dst *MutateConversionActionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateConversionActionResult.Merge(dst, src)
}
func (m *MutateConversionActionResult) XXX_Size() int {
	return xxx_messageInfo_MutateConversionActionResult.Size(m)
}
func (m *MutateConversionActionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateConversionActionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateConversionActionResult proto.InternalMessageInfo

func (m *MutateConversionActionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetConversionActionRequest)(nil), "google.ads.googleads.v0.services.GetConversionActionRequest")
	proto.RegisterType((*MutateConversionActionsRequest)(nil), "google.ads.googleads.v0.services.MutateConversionActionsRequest")
	proto.RegisterType((*ConversionActionOperation)(nil), "google.ads.googleads.v0.services.ConversionActionOperation")
	proto.RegisterType((*MutateConversionActionsResponse)(nil), "google.ads.googleads.v0.services.MutateConversionActionsResponse")
	proto.RegisterType((*MutateConversionActionResult)(nil), "google.ads.googleads.v0.services.MutateConversionActionResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConversionActionServiceClient is the client API for ConversionActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversionActionServiceClient interface {
	// Returns the requested conversion action.
	GetConversionAction(ctx context.Context, in *GetConversionActionRequest, opts ...grpc.CallOption) (*resources.ConversionAction, error)
	// Creates, updates or removes conversion actions. Operation statuses are
	// returned.
	MutateConversionActions(ctx context.Context, in *MutateConversionActionsRequest, opts ...grpc.CallOption) (*MutateConversionActionsResponse, error)
}

type conversionActionServiceClient struct {
	cc *grpc.ClientConn
}

func NewConversionActionServiceClient(cc *grpc.ClientConn) ConversionActionServiceClient {
	return &conversionActionServiceClient{cc}
}

func (c *conversionActionServiceClient) GetConversionAction(ctx context.Context, in *GetConversionActionRequest, opts ...grpc.CallOption) (*resources.ConversionAction, error) {
	out := new(resources.ConversionAction)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.ConversionActionService/GetConversionAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionActionServiceClient) MutateConversionActions(ctx context.Context, in *MutateConversionActionsRequest, opts ...grpc.CallOption) (*MutateConversionActionsResponse, error) {
	out := new(MutateConversionActionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.ConversionActionService/MutateConversionActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionActionServiceServer is the server API for ConversionActionService service.
type ConversionActionServiceServer interface {
	// Returns the requested conversion action.
	GetConversionAction(context.Context, *GetConversionActionRequest) (*resources.ConversionAction, error)
	// Creates, updates or removes conversion actions. Operation statuses are
	// returned.
	MutateConversionActions(context.Context, *MutateConversionActionsRequest) (*MutateConversionActionsResponse, error)
}

func RegisterConversionActionServiceServer(s *grpc.Server, srv ConversionActionServiceServer) {
	s.RegisterService(&_ConversionActionService_serviceDesc, srv)
}

func _ConversionActionService_GetConversionAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversionActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionActionServiceServer).GetConversionAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.ConversionActionService/GetConversionAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionActionServiceServer).GetConversionAction(ctx, req.(*GetConversionActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionActionService_MutateConversionActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionActionServiceServer).MutateConversionActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.ConversionActionService/MutateConversionActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionActionServiceServer).MutateConversionActions(ctx, req.(*MutateConversionActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversionActionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.ConversionActionService",
	HandlerType: (*ConversionActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConversionAction",
			Handler:    _ConversionActionService_GetConversionAction_Handler,
		},
		{
			MethodName: "MutateConversionActions",
			Handler:    _ConversionActionService_MutateConversionActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/conversion_action_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/conversion_action_service.proto", fileDescriptor_conversion_action_service_a3114b9b067f0aff)
}

var fileDescriptor_conversion_action_service_a3114b9b067f0aff = []byte{
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x4f, 0xd4, 0x4c,
	0x18, 0xc0, 0xdf, 0x76, 0xdf, 0xa0, 0x4c, 0x51, 0x93, 0x21, 0x86, 0x75, 0x43, 0x60, 0x53, 0x49,
	0x24, 0x7b, 0x68, 0x37, 0x4b, 0x34, 0xb1, 0x2b, 0x86, 0x42, 0x04, 0x3c, 0x20, 0xa4, 0x24, 0xc4,
	0xe8, 0x26, 0xcd, 0xd0, 0x0e, 0x9b, 0x86, 0xb6, 0x53, 0x67, 0xa6, 0x6b, 0x08, 0xe1, 0xa2, 0x1f,
	0xc1, 0xb3, 0x17, 0x8f, 0x7e, 0x0d, 0xc3, 0xc5, 0xab, 0x9f, 0xc0, 0xc4, 0x93, 0xf1, 0x43, 0x98,
	0xe9, 0x74, 0x56, 0x58, 0xa8, 0x6b, 0xe0, 0xb4, 0xcf, 0x3e, 0xcf, 0x33, 0xbf, 0xe7, 0xef, 0x4c,
	0xc1, 0x4a, 0x9f, 0x90, 0x7e, 0x8c, 0x6d, 0x14, 0x32, 0x5b, 0x8a, 0x42, 0x1a, 0xb4, 0x6d, 0x86,
	0xe9, 0x20, 0x0a, 0x30, 0xb3, 0x03, 0x92, 0x0e, 0x30, 0x65, 0x11, 0x49, 0x7d, 0x14, 0x70, 0xf1,
	0x53, 0x9a, 0xac, 0x8c, 0x12, 0x4e, 0x60, 0x53, 0x1e, 0xb3, 0x50, 0xc8, 0xac, 0x21, 0xc1, 0x1a,
	0xb4, 0x2d, 0x45, 0x68, 0x3c, 0xae, 0x8a, 0x41, 0x31, 0x23, 0x39, 0xbd, 0x34, 0x88, 0x84, 0x37,
	0x66, 0xd5, 0xd1, 0x2c, 0xb2, 0x51, 0x9a, 0x12, 0x8e, 0x84, 0x91, 0x95, 0xd6, 0x32, 0xb4, 0x5d,
	0xfc, 0xdb, 0xcf, 0x0f, 0xec, 0x83, 0x08, 0xc7, 0xa1, 0x9f, 0x20, 0x76, 0x58, 0x7a, 0xcc, 0x8d,
	0x7a, 0xbc, 0xa5, 0x28, 0xcb, 0x30, 0x55, 0x84, 0x99, 0xd2, 0x4e, 0xb3, 0xc0, 0x66, 0x1c, 0xf1,
	0xbc, 0x34, 0x98, 0x2e, 0x68, 0x6c, 0x60, 0xbe, 0x36, 0x4c, 0xcb, 0x2d, 0xb2, 0xf2, 0xf0, 0x9b,
	0x1c, 0x33, 0x0e, 0xef, 0x83, 0x5b, 0x2a, 0x77, 0x3f, 0x45, 0x09, 0xae, 0x6b, 0x4d, 0x6d, 0x71,
	0xd2, 0x9b, 0x52, 0xca, 0x17, 0x28, 0xc1, 0xe6, 0x2f, 0x0d, 0xcc, 0x6d, 0xe5, 0x1c, 0x71, 0x3c,
	0x8a, 0x61, 0x8a, 0x33, 0x0f, 0x8c, 0x20, 0x67, 0x9c, 0x24, 0x98, 0xfa, 0x51, 0x58, 0x52, 0x80,
	0x52, 0x3d, 0x0f, 0xe1, 0x6b, 0x00, 0x48, 0x86, 0xa9, 0xac, 0xba, 0xae, 0x37, 0x6b, 0x8b, 0x46,
	0xa7, 0x6b, 0x8d, 0xeb, 0xb8, 0x35, 0x1a, 0x70, 0x5b, 0x31, 0xbc, 0x33, 0x38, 0xf8, 0x00, 0xdc,
	0xc9, 0x10, 0xe5, 0x11, 0x8a, 0xfd, 0x03, 0x14, 0xc5, 0x39, 0xc5, 0xf5, 0x5a, 0x53, 0x5b, 0xbc,
	0xe9, 0xdd, 0x2e, 0xd5, 0xeb, 0x52, 0x2b, 0xca, 0x1d, 0xa0, 0x38, 0x0a, 0x11, 0xc7, 0x3e, 0x49,
	0xe3, 0xa3, 0xfa, 0xff, 0x85, 0xdb, 0x94, 0x52, 0x6e, 0xa7, 0xf1, 0x91, 0xf9, 0x51, 0x07, 0xf7,
	0x2a, 0xe3, 0xc2, 0x2e, 0x30, 0xf2, 0xac, 0x00, 0x88, 0xe9, 0x14, 0x00, 0xa3, 0xd3, 0x50, 0x95,
	0xa8, 0xf1, 0x58, 0xeb, 0x62, 0x80, 0x5b, 0x88, 0x1d, 0x7a, 0x40, 0xba, 0x0b, 0x19, 0x6e, 0x81,
	0x89, 0x80, 0x62, 0xc4, 0x65, 0x9f, 0x8d, 0xce, 0x52, 0x65, 0x07, 0x86, 0x1b, 0x75, 0xa1, 0x05,
	0x9b, 0xff, 0x79, 0x25, 0x44, 0xe0, 0x24, 0xbc, 0xae, 0x5f, 0x0b, 0x27, 0x21, 0xb0, 0x0e, 0x26,
	0x28, 0x4e, 0xc8, 0x40, 0x76, 0x6f, 0x52, 0x58, 0xe4, 0xff, 0x55, 0x03, 0x4c, 0x0e, 0xdb, 0x6d,
	0x7e, 0xd1, 0xc0, 0x7c, 0xe5, 0x3a, 0xb0, 0x8c, 0xa4, 0x0c, 0xc3, 0x75, 0x70, 0x77, 0x64, 0x22,
	0x3e, 0xa6, 0x94, 0xd0, 0x82, 0x6c, 0x74, 0xa0, 0x4a, 0x94, 0x66, 0x81, 0xb5, 0x5b, 0xac, 0xab,
	0x37, 0x7d, 0x7e, 0x56, 0xcf, 0x84, 0x3b, 0x7c, 0x09, 0x6e, 0x50, 0xcc, 0xf2, 0x98, 0xab, 0x9d,
	0x79, 0x3a, 0x7e, 0x67, 0x2e, 0xcf, 0xcd, 0x2b, 0x30, 0x9e, 0xc2, 0x99, 0x6b, 0x60, 0xf6, 0x6f,
	0x8e, 0xff, 0x74, 0x33, 0x3a, 0xa7, 0x35, 0x30, 0x33, 0x7a, 0x7e, 0x57, 0xe6, 0x01, 0x4f, 0x35,
	0x30, 0x7d, 0xc9, 0xcd, 0x83, 0x4f, 0xc6, 0x57, 0x50, 0x7d, 0x61, 0x1b, 0x57, 0x19, 0xb1, 0xd9,
	0x7d, 0xf7, 0xed, 0xc7, 0x07, 0xfd, 0x21, 0x5c, 0x12, 0x6f, 0xd5, 0xf1, 0xb9, 0xb2, 0x96, 0xd5,
	0x1d, 0x65, 0x76, 0xeb, 0xcc, 0xe3, 0x55, 0xce, 0xd3, 0x6e, 0x9d, 0xc0, 0xef, 0x1a, 0x98, 0xa9,
	0x18, 0x37, 0x5c, 0xb9, 0xea, 0x34, 0xd4, 0xc3, 0xd1, 0x70, 0xaf, 0x41, 0x90, 0xbb, 0x66, 0xba,
	0x45, 0x75, 0x5d, 0xf3, 0x91, 0xa8, 0xee, 0x4f, 0x39, 0xc7, 0x67, 0x1e, 0xa4, 0xe5, 0xd6, 0xc9,
	0xc5, 0xe2, 0x9c, 0xa4, 0x00, 0x3b, 0x5a, 0x6b, 0xf5, 0xbd, 0x0e, 0x16, 0x02, 0x92, 0x8c, 0xcd,
	0x65, 0x75, 0xb6, 0x62, 0xda, 0x3b, 0xe2, 0xde, 0xef, 0x68, 0xaf, 0x36, 0x4b, 0x42, 0x9f, 0xc4,
	0x28, 0xed, 0x5b, 0x84, 0xf6, 0xed, 0x3e, 0x4e, 0x8b, 0x57, 0x41, 0x7d, 0x31, 0xb2, 0x88, 0x55,
	0x7f, 0xa4, 0xba, 0x4a, 0xf8, 0xa4, 0xd7, 0x36, 0x5c, 0xf7, 0xb3, 0xde, 0xdc, 0x90, 0x40, 0x37,
	0x64, 0x96, 0x14, 0x85, 0xb4, 0xd7, 0xb6, 0xca, 0xc0, 0xec, 0xab, 0x72, 0xe9, 0xb9, 0x21, 0xeb,
	0x0d, 0x5d, 0x7a, 0x7b, 0xed, 0x9e, 0x72, 0xf9, 0xa9, 0x2f, 0x48, 0xbd, 0xe3, 0xb8, 0x21, 0x73,
	0x9c, 0xa1, 0x93, 0xe3, 0xec, 0xb5, 0x1d, 0x47, 0xb9, 0xed, 0x4f, 0x14, 0x79, 0x2e, 0xfd, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x5f, 0x8d, 0xab, 0xce, 0x4b, 0x07, 0x00, 0x00,
}
