package model

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"
	"getitqec.com/server/mailnotification/pkg/commons"
	"getitqec.com/server/mailnotification/pkg/dto"
)

func (m *NotificationModel) SetUserToken(ctx context.Context, domain string, userId string, tokenType string, platform string, token string) error {
	return m.TokenDAO.Set(ctx, &dto.MessagingToken{Id: userId, Domain: domain, Token: token, Type: tokenType, Platform: platform})
}

// Push to particular user
func (m *NotificationModel) PushToUser(ctx context.Context, notification bool, domain string, userId string, title string, body string, imgUrl string, data map[string]string) error {
	// // load PATH_TO_SERVICE_FILE.json base on domain
	// opt := option.WithCredentialsFile("PATH_TO_SERVICE_FILE.json")
	// app, err := firebase.NewApp(context.Background(), nil, opt)

	// if err != nil {
	// 	log.Fatalf("error initializing app: %v\n", err)
	// }

	// Obtain a messaging.Client from the App.
	client, err := m.App.Messaging(ctx)
	if err != nil {
		return err
	}

	// get user token
	tokens, err := m.TokenDAO.Get(ctx, &dto.MessagingToken{Id: userId, Domain: domain})
	if err != nil {
		return err
	}
	// See documentation on defining a message payload.
	if notification {
		data["clickAction"] = "FLUTTER_NOTIFICATION_CLICK" // For show notification
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: imgUrl,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
			},
		},
		Data: data,
	}
	for _, token := range tokens {
		// use fcm
		if token.Type == commons.FirebaseMessaging {
			// Send a message to the device corresponding to the provided
			// registration token.
			message.Token = token.Token
			response, err := client.Send(ctx, message)
			if err != nil {
				return err
			}
			// Response is a message ID string.
			fmt.Println("Successfully sent message:", response)
		}
	}
	return nil
}

// Push to a group, such as rider
func (m *NotificationModel) SendNotificationToTopic(ctx context.Context, notification bool, domain string, topic string, title string, body string, imgUrl string, data map[string]string) error {
	// // load PATH_TO_SERVICE_FILE.json base on domain
	// opt := option.WithCredentialsFile("PATH_TO_SERVICE_FILE.json")
	// app, err := firebase.NewApp(context.Background(), nil, opt)

	// if err != nil {
	// 	log.Fatalf("error initializing app: %v\n", err)
	// }

	// Obtain a messaging.Client from the App.
	client, err := m.App.Messaging(ctx)
	if err != nil {
		return err
	}

	// See documentation on defining a message payload.
	data["clickAction"] = "FLUTTER_NOTIFICATION_CLICK" // For show notification
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data:  data,
		Topic: topic,
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	return nil
}
