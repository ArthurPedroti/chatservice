package gateway

import (
	"context"

	"github.com/ArthurPedroti/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreatedAt(ctx context.Context, chat *entity.Chat) error
	FindChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error ///
}
