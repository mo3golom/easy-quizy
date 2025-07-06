package game

import (
	"context"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"

	"github.com/google/uuid"
)

func (u *Usecase) Get(ctx context.Context, id uuid.UUID) (*model.Game, error) {
	specificGames, err := u.games.GetGamesByIDs(ctx, []uuid.UUID{id})
	if err != nil {
		return nil, err
	}
	if len(specificGames) == 0 {
		return nil, contracts.ErrGameNotFound
	}

	return &specificGames[0], nil
}
