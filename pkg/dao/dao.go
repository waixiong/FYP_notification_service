package dao

import (
	"context"

	"getitqec.com/server/mailnotification/pkg/dto"
)

type IMessagingTokenDAO interface {
	// Set(*dto.User) error
	Set(ctx context.Context, token *dto.MessagingToken) error
	Get(ctx context.Context, token *dto.MessagingToken) ([]*dto.MessagingToken, error)
	Delete(ctx context.Context, token *dto.MessagingToken) (*dto.MessagingToken, error)
}
