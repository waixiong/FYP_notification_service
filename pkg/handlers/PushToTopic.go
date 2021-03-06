package handlers

import (
	"context"

	pb "getitqec.com/server/mailnotification/pkg/api/v1"
	"getitqec.com/server/mailnotification/pkg/model"
)

type PushToTopicHandler struct {
	Model model.NotificationModelI
}

func (s *PushToTopicHandler) PushToTopic(ctx context.Context, req *pb.PushRequest) error {
	// userID, err := commons.GetUserID(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	return s.Model.SendNotificationToTopic(ctx, req.Notification, req.Domain, req.Target, req.Title, req.Body, req.ImgUrl, req.Data)
}
