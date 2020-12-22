// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package serviceproto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MailNotificationServiceClient is the client API for MailNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailNotificationServiceClient interface {
	// reuqest response
	SendNoReplyMail(ctx context.Context, in *NoReplyMessage, opts ...grpc.CallOption) (*empty.Empty, error)
	SendToken(ctx context.Context, in *MessagingToken, opts ...grpc.CallOption) (*empty.Empty, error)
	PushToUser(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PushNotificationToTopic(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type mailNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailNotificationServiceClient(cc grpc.ClientConnInterface) MailNotificationServiceClient {
	return &mailNotificationServiceClient{cc}
}

func (c *mailNotificationServiceClient) SendNoReplyMail(ctx context.Context, in *NoReplyMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mailnotificationproto.MailNotificationService/SendNoReplyMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailNotificationServiceClient) SendToken(ctx context.Context, in *MessagingToken, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mailnotificationproto.MailNotificationService/SendToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailNotificationServiceClient) PushToUser(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mailnotificationproto.MailNotificationService/PushToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailNotificationServiceClient) PushNotificationToTopic(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mailnotificationproto.MailNotificationService/PushNotificationToTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailNotificationServiceServer is the server API for MailNotificationService service.
// All implementations must embed UnimplementedMailNotificationServiceServer
// for forward compatibility
type MailNotificationServiceServer interface {
	// reuqest response
	SendNoReplyMail(context.Context, *NoReplyMessage) (*empty.Empty, error)
	SendToken(context.Context, *MessagingToken) (*empty.Empty, error)
	PushToUser(context.Context, *PushRequest) (*empty.Empty, error)
	PushNotificationToTopic(context.Context, *PushRequest) (*empty.Empty, error)
	mustEmbedUnimplementedMailNotificationServiceServer()
}

// UnimplementedMailNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailNotificationServiceServer struct {
}

func (UnimplementedMailNotificationServiceServer) SendNoReplyMail(context.Context, *NoReplyMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNoReplyMail not implemented")
}
func (UnimplementedMailNotificationServiceServer) SendToken(context.Context, *MessagingToken) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToken not implemented")
}
func (UnimplementedMailNotificationServiceServer) PushToUser(context.Context, *PushRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToUser not implemented")
}
func (UnimplementedMailNotificationServiceServer) PushNotificationToTopic(context.Context, *PushRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushNotificationToTopic not implemented")
}
func (UnimplementedMailNotificationServiceServer) mustEmbedUnimplementedMailNotificationServiceServer() {
}

// UnsafeMailNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailNotificationServiceServer will
// result in compilation errors.
type UnsafeMailNotificationServiceServer interface {
	mustEmbedUnimplementedMailNotificationServiceServer()
}

func RegisterMailNotificationServiceServer(s grpc.ServiceRegistrar, srv MailNotificationServiceServer) {
	s.RegisterService(&_MailNotificationService_serviceDesc, srv)
}

func _MailNotificationService_SendNoReplyMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoReplyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailNotificationServiceServer).SendNoReplyMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailnotificationproto.MailNotificationService/SendNoReplyMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailNotificationServiceServer).SendNoReplyMail(ctx, req.(*NoReplyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailNotificationService_SendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagingToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailNotificationServiceServer).SendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailnotificationproto.MailNotificationService/SendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailNotificationServiceServer).SendToken(ctx, req.(*MessagingToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailNotificationService_PushToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailNotificationServiceServer).PushToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailnotificationproto.MailNotificationService/PushToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailNotificationServiceServer).PushToUser(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailNotificationService_PushNotificationToTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailNotificationServiceServer).PushNotificationToTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailnotificationproto.MailNotificationService/PushNotificationToTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailNotificationServiceServer).PushNotificationToTopic(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailNotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mailnotificationproto.MailNotificationService",
	HandlerType: (*MailNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNoReplyMail",
			Handler:    _MailNotificationService_SendNoReplyMail_Handler,
		},
		{
			MethodName: "SendToken",
			Handler:    _MailNotificationService_SendToken_Handler,
		},
		{
			MethodName: "PushToUser",
			Handler:    _MailNotificationService_PushToUser_Handler,
		},
		{
			MethodName: "PushNotificationToTopic",
			Handler:    _MailNotificationService_PushNotificationToTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}