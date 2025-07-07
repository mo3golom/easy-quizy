package game

import (
	"context"
	"easy-quizy/internal/model"

	"github.com/google/uuid"
)

func (u *Usecase) Reset(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) error {
	specificGame, err := u.Get(ctx, gameID)
	if err != nil {
		return err
	}

	// Ежедневный квиз нельзя перезапустить
	if specificGame.Type == model.GameTypeDaily {
		return nil
	}

	return u.games.DeleteGameSessionAnswers(ctx, gameID, playerID)
}
