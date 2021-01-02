package handlers

import (
	"context"
	"fmt"

	pb "getitqec.com/server/mailnotification/pkg/api/v1"
	"getitqec.com/server/mailnotification/pkg/commons"
	"getitqec.com/server/mailnotification/pkg/model"
)

type SendTokenHandler struct {
	Model model.NotificationModelI
}

func (s *SendTokenHandler) SendToken(ctx context.Context, req *pb.MessagingToken) error {
	// userID, err := commons.GetUserID(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("Send FCM Token Handler")
	token, err := commons.VerifyGoogleAccessToken(ctx)
	if err != nil {
		fmt.Println("Error found in handler")
		return err
	}
	fmt.Println(token)
	if token.UserId != req.Id && req.Id != "" {
		return commons.NotAuthorized
	}
	req.Id = token.UserId
	return s.Model.SetUserToken(ctx, req.Domain, req.Id, req.Type, req.Platform, req.Token)
}
