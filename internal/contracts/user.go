package contracts

import (
	"context"
	"easy-quizy/internal/model"
	"errors"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserChatNotFound = errors.New("user chat not found")
)

type (
	UserData struct {
		UserIDext string
		Source    string
		ChatID    *int64
		ChatType  *string
	}

	UserUsecase interface {
		RetrieveUser(ctx context.Context, data UserData) (model.User, error)
	}
)
