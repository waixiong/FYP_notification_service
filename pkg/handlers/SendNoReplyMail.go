package handlers

import (
	"context"

	pb "getitqec.com/server/mailnotification/pkg/api/v1"
	"getitqec.com/server/mailnotification/pkg/model"
)

type SendNoReplyMailHandler struct {
	Model model.MailModelI
}

func (s *SendNoReplyMailHandler) SendNoReplyMail(ctx context.Context, req *pb.NoReplyMessage) error {
	// userID, err := commons.GetUserID(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	return s.Model.SendNoReplyMail(ctx, req.RepName, req.RepEmail, req.Subj, req.Body)
}
