package user

import (
	"context"
	"easy-quizy/internal/model"

	"github.com/google/uuid"
)

type (
	repository interface {
		InsertSource(ctx context.Context, user model.UserSource) error
		InsertUserChat(ctx context.Context, user model.UserChat) error
		GetUserBySource(ctx context.Context, userIDext string, source string) (model.User, error)
		GetUserChat(ctx context.Context, userID uuid.UUID, chatID int64) (model.UserChat, error)
	}
)
