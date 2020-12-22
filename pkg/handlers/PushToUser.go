package handlers

import (
	"context"

	pb "getitqec.com/server/mailnotification/pkg/api/v1"
	"getitqec.com/server/mailnotification/pkg/model"
)

type PushToUserHandler struct {
	Model model.NotificationModelI
}

func (s *PushToUserHandler) PushToUser(ctx context.Context, req *pb.PushRequest) error {
	// userID, err := commons.GetUserID(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	return s.Model.PushToUser(ctx, req.Notification, req.Domain, req.Target, req.Title, req.Body, req.ImgUrl, req.Data)
}
