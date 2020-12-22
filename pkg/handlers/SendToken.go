package handlers

import (
	"context"

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
	token, err := commons.VerifyGoogleAccessToken(ctx)
	if err != nil {
		return err
	}
	if token.UserId != req.Id && req.Id != "" {
		return commons.NotAuthorized
	}
	req.Id = token.UserId
	return s.Model.SetUserToken(ctx, req.Domain, req.Id, req.Type, req.Platform, req.Token)
}
