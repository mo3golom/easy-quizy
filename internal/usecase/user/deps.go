package user

import (
	"context"
	"easy-quizy/internal/model"
)

type (
	repository interface {
		InsertSource(ctx context.Context, user model.UserSource) error
		GetUserBySource(ctx context.Context, userIDext string, source string) (model.User, error)
	}
)
