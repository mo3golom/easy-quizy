package game

import (
	"context"
	"easy-quizy/internal/model"

	"github.com/google/uuid"
)

type (
	repository interface {
		GetGamesByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Game, error)
		GetGameSession(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) (model.GameSession, error)
		InsertGameSessionAnswer(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID, data model.GameSessionAnswer) error
	}
)
