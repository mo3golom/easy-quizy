package game

import (
	"context"
	"quizzly-v2/internal/model"

	"github.com/google/uuid"
)

type (
	repository interface {
		GetGamesByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Game, error)
		GetGameSession(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) (model.GameSession, error)
	}
)
