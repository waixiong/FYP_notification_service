package model

import (
	"context"
)

type MailModelI interface {
	SendNoReplyMail(ctx context.Context, repName string, repAddress string, subj string, body string) error
}

type NotificationModelI interface {
	// set token
	SetUserToken(ctx context.Context, domain string, userId string, tokenType string, platform string, token string) error
	// Push to particular user
	PushToUser(ctx context.Context, notification bool, domain string, userId string, title string, body string, imgUrl string, data map[string]string) error
	// Push to a group, such as rider
	SendNotificationToTopic(ctx context.Context, notification bool, domain string, topic string, title string, body string, imgUrl string, data map[string]string) error
}
