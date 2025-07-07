package game

import (
	"context"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"

	"github.com/google/uuid"
)

func (u *Usecase) Get(ctx context.Context, id uuid.UUID) (model.Game, error) {
	specificGames, err := u.games.GetGamesByIDs(ctx, []uuid.UUID{id})
	if err != nil {
		return model.Game{}, err
	}
	if len(specificGames) == 0 {
		return model.Game{}, contracts.ErrGameNotFound
	}

	return specificGames[0], nil
}

func (u *Usecase) GetDaily(ctx context.Context) (model.Game, error) {
	return u.games.GetDailyGame(ctx)
}
