package contracts

import (
	"context"
	"easy-quizy/internal/model"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type (
	UserUsecase interface {
		RetrieveUser(ctx context.Context, userIDext string, source string) (model.User, error)
	}
)
