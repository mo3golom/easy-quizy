package user

import (
	"context"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"
	"errors"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/google/uuid"
)

type Usecase struct {
	trm        trm.Manager
	repository repository
}

func NewUsecase(
	repository repository,
	trm trm.Manager,
) *Usecase {
	return &Usecase{
		trm:        trm,
		repository: repository,
	}
}

func (u *Usecase) RetrieveUser(ctx context.Context, data contracts.UserData) (model.User, error) {
	var user model.User
	err := u.trm.Do(ctx, func(ctx context.Context) error {
		var err error
		user, err = u.repository.GetUserBySource(ctx, data.UserIDext, data.Source)
		if err != nil && !errors.Is(err, contracts.ErrUserNotFound) {
			return err
		}

		if errors.Is(err, contracts.ErrUserNotFound) {
			user = model.User{
				ID: uuid.New(),
			}

			return u.repository.InsertSource(ctx, model.UserSource{
				User:   user,
				IDext:  data.UserIDext,
				Source: data.Source,
			})
		}

		return nil
	})
	if err != nil {
		return model.User{}, err
	}

	if data.ChatID != nil && data.ChatType != nil {
		err := u.trm.Do(ctx, func(ctx context.Context) error {
			_, err := u.repository.GetUserChat(ctx, user.ID, *data.ChatID)
			if err != nil && !errors.Is(err, contracts.ErrUserChatNotFound) {
				return err
			}

			if errors.Is(err, contracts.ErrUserChatNotFound) {
				err = u.repository.InsertUserChat(ctx, model.UserChat{
					User:     user,
					ChatID:   *data.ChatID,
					ChatType: *data.ChatType,
				})

				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return model.User{}, err
		}

	}

	return user, nil
}
