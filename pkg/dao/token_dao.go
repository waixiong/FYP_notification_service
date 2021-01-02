package dao

import (
	"context"
	"fmt"

	"getitqec.com/server/mailnotification/pkg/commons"
	"getitqec.com/server/mailnotification/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessagingTokenDAO struct {
	// db      commons.DB
	mongodb commons.MongoDB
}

func InitMessagingTokenDAO(m commons.MongoDB) IMessagingTokenDAO {
	return &MessagingTokenDAO{mongodb: m}
}

func (v *MessagingTokenDAO) Set(ctx context.Context, token *dto.MessagingToken) error {
	option := options.Find()
	option = option.SetLimit(1)
	c, err := v.mongodb.Client().Database(commons.NotificationTable).Collection(commons.MessagingTokenColection).CountDocuments(ctx, bson.D{
		{Key: "id", Value: token.Id},
		{Key: "type", Value: token.Type},
		{Key: "platform", Value: token.Platform},
		{Key: "domain", Value: token.Domain},
	})
	if err != nil {
		fmt.Println("Error here DAO set 1")
		return err
	}
	if c == 0 {
		return v.mongodb.Create(ctx, commons.NotificationTable, commons.MessagingTokenColection, token)
	}
	return v.mongodb.Update(ctx, commons.NotificationTable, commons.MessagingTokenColection, bson.D{
		{Key: "id", Value: token.Id},
		{Key: "type", Value: token.Type},
		{Key: "platform", Value: token.Platform},
		{Key: "domain", Value: token.Domain},
	}, bson.D{{"$set", bson.M{"token": token.Token}}})
}

func (v *MessagingTokenDAO) Get(ctx context.Context, token *dto.MessagingToken) ([]*dto.MessagingToken, error) {
	document := bson.M{
		"id": token.Id,
	}
	if token.Type != "" {
		document["type"] = token.Type
	}
	if token.Platform != "" {
		document["platform"] = token.Platform
	}
	raws, err := v.mongodb.BatchRead(ctx, commons.NotificationTable, commons.MessagingTokenColection, document)
	if err != nil {
		return nil, err
	}

	tokens := []*dto.MessagingToken{}
	for _, raw := range raws {
		token := &dto.MessagingToken{}
		err := bson.Unmarshal(*raw, token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (v *MessagingTokenDAO) Delete(ctx context.Context, token *dto.MessagingToken) (*dto.MessagingToken, error) {
	// data, err := v.db.Delete("User", "UserID", id)
	// if err != nil {
	// 	return nil, err
	// }
	// if data == nil {
	// 	return nil, commons.ErrExpiredToken
	// }
	// return dto.AttributeMapToUser(data)
	result := v.mongodb.Delete(ctx, commons.NotificationTable, commons.MessagingTokenColection, bson.D{
		{Key: "Id", Value: token.Id},
		{Key: "type", Value: token.Type},
		{Key: "platform", Value: token.Platform},
		{Key: "domain", Value: token.Domain},
	})
	if result.Err() != nil {
		return nil, result.Err()
	}
	t := &dto.MessagingToken{}
	err := result.Decode(t)
	return t, err
}
