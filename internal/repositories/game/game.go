package game

import (
	"context"
	"quizzly-v2/internal/model"
	"quizzly-v2/pkg/structs/collections/slices"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	sqlxGame struct {
		ID        uuid.UUID `db:"id"`
		Payload   []byte    `db:"payload"`
		CreatedAt time.Time `db:"created_at"`
	}

	sqlxGameSession struct {
		GameID     uuid.UUID `db:"game_id"`
		PlayerID   uuid.UUID `db:"player_id"`
		QuestionID int64     `db:"question_id"`
		AnswerID   int64     `db:"answer_id"`
	}
)

func (r *DefaultRepository) GetGamesByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Game, error) {
	const query = `
		select 
			id, 
			payload, 
			created_at
		from easy_quizy_game
		where id = any($1)
	`

	var result []sqlxGame
	if err := r.db(ctx).SelectContext(
		ctx,
		&result,
		query,
		pq.Array(ids),
	); err != nil {
		return nil, err
	}

	return slices.Map(result, func(i sqlxGame) (model.Game, error) {
		return convertToGame(i)
	})
}

func (r *DefaultRepository) GetGameSession(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) (model.GameSession, error) {
	const query = `
		select 
			game_id,
            player_id,
            question_id,
    		answer_id
		from easy_quizy_game_session
		where game_id = $1 and player_id = $2
	`

	var result []sqlxGameSession
	if err := r.db(ctx).SelectContext(
		ctx,
		&result,
		query,
		gameID,
		playerID,
	); err != nil {
		return model.GameSession{}, err
	}

	return convertToSession(result), nil
}
