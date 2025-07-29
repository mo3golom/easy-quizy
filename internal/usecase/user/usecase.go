package user

import (
	"context"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"
	"errors"

	"github.com/google/uuid"
)

type Usecase struct {
	repository repository
}

func NewUsecase(
	repository repository,
) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (u *Usecase) RetrieveUser(ctx context.Context, userIDext string, source string) (model.User, error) {
	user, err := u.repository.GetUserBySource(ctx, userIDext, source)
	if err != nil && !errors.Is(err, contracts.ErrUserNotFound) {
		return model.User{}, err
	}

	if errors.Is(err, contracts.ErrUserNotFound) {
		user = model.User{
			ID: uuid.New(),
		}

		err := u.repository.InsertSource(ctx, model.UserSource{
			User:   user,
			IDext:  userIDext,
			Source: source,
		})
		if err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}
