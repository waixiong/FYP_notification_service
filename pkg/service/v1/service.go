package service

import (
	"context"
	"net/http"

	pb "getitqec.com/server/mailnotification/pkg/api/v1"
	"getitqec.com/server/mailnotification/pkg/handlers"
	"getitqec.com/server/mailnotification/pkg/model"
	"github.com/golang/protobuf/ptypes/empty"
	//pb "./proto"
)

var httpClient = &http.Client{}

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
// func logger(format string, a ...interface{}) {
// 	fmt.Printf("LOG:\t"+format+"\n", a...)
// }

// var (
// 	//port = flag.Int("port", 50051, "the port to serve on")

// 	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
// 	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
// )

// Server class
type Server struct {
	mailModel         model.MailModelI
	notificationModel model.NotificationModelI
	pb.UnimplementedMailNotificationServiceServer
}

func (s *Server) SendNoReplyMail(ctx context.Context, req *pb.NoReplyMessage) (*empty.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method SendNoReplyMail not implemented")
	handler := &handlers.SendNoReplyMailHandler{Model: s.mailModel}
	return &empty.Empty{}, handler.SendNoReplyMail(ctx, req)
}

func (s *Server) SendToken(ctx context.Context, req *pb.MessagingToken) (*empty.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method SendToken not implemented")
	handler := &handlers.SendTokenHandler{Model: s.notificationModel}
	return &empty.Empty{}, handler.SendToken(ctx, req)
}

func (s *Server) PushToUser(ctx context.Context, req *pb.PushRequest) (*empty.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method PushNotificationToUser not implemented")
	handler := &handlers.PushToUserHandler{Model: s.notificationModel}
	return &empty.Empty{}, handler.PushToUser(ctx, req)
}

func (s *Server) PushNotificationToTopic(ctx context.Context, req *pb.PushRequest) (*empty.Empty, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method PushNotificationToTopic not implemented")
	handler := &handlers.PushToTopicHandler{Model: s.notificationModel}
	return &empty.Empty{}, handler.PushToTopic(ctx, req)
}

// NewServer return new auth server service
func NewServer(mmodel model.MailModelI, nmodel model.NotificationModelI) *Server {
	server := &Server{}
	server.mailModel = mmodel
	server.notificationModel = nmodel
	return server
}
