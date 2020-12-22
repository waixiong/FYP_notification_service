package commons

import (
	"context"
	"fmt"

	//pb "./proto"

	"getitqec.com/server/mailnotification/pkg/logger"
	"google.golang.org/grpc/metadata"
	// "google.golang.org/grpc/metadata"
)

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
// func logger(format string, a ...interface{}) {
// 	fmt.Printf("LOG:\t"+format+"\n", a...)
// }

var (
//port = flag.Int("port", 50051, "the port to serve on")

// ErrMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
// ErrInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
// ErrExpiredToken    = status.Errorf(codes.Unauthenticated, "expired token")
// UnimplementedError = status.Errorf(codes.Unimplemented, "Not done yet")
// UserNotFound       = status.Errorf(codes.NotFound, "User not found")
// UserAlreadyExist   = status.Errorf(codes.AlreadyExists, "User already exist")
)

func GetUserID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	// fmt.Printf("\tAuth Func Get metadata...\n")
	if !ok {
		logger.Log.Error(fmt.Sprintf("Get User ID from metadata fail"))
		return "", ErrMissingMetadata
	}
	users := md.Get("User")
	if len(users) != 1 {
		logger.Log.Error(fmt.Sprintf("Get User ID from metadata fail"))
		return "", ErrMissingMetadata
	}
	return users[0], nil
}

const (
	NotificationTable       = "Notification"
	MessagingTokenColection = "MessagingToken"
	FirebaseMessaging       = "FirebaseMessaging"
)
