package game

import (
	"context"
	"database/sql"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"
	"easy-quizy/pkg/structs/collections/slices"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	sqlxGame struct {
		ID        uuid.UUID `db:"id"`
		Type      string    `db:"type"`
		Payload   []byte    `db:"payload"`
		CreatedAt time.Time `db:"created_at"`
	}

	sqlxGameSession struct {
		GameID     uuid.UUID `db:"game_id"`
		PlayerID   uuid.UUID `db:"player_id"`
		QuestionID int64     `db:"question_id"`
		AnswerID   int64     `db:"answer_id"`
		IsCorrect  bool      `db:"is_correct"`
	}
)

func (r *DefaultRepository) GetGamesByIDs(ctx context.Context, ids []uuid.UUID) ([]model.Game, error) {
	const query = `
		select 
			id,
			type, 
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

func (r *DefaultRepository) GetDailyGame(ctx context.Context) (model.Game, error) {
	const query = `
		select 
			g.id, 
			g.type, 
			g.payload, 
			g.created_at
		from easy_quizy_game g
		inner join game_daily gd on g.id = gd.game_id
		where gd.ended_at is null
		order by gd.created_at asc
		limit 1
	`

	var result sqlxGame
	if err := r.db(ctx).GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Game{}, contracts.ErrGameNotFound
		}

		return model.Game{}, err
	}

	return convertToGame(result)
}

func (r *DefaultRepository) InsertGameSessionAnswer(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID, data model.GameSessionAnswer) error {
	const query = `
		insert into easy_quizy_game_session
		(game_id, player_id, question_id, answer_id, is_correct)
		values ($1, $2, $3, $4, $5)
	`

	_, err := r.db(ctx).ExecContext(
		ctx,
		query,
		gameID,
		playerID,
		data.QuestionID,
		data.AnswerID,
		data.IsCorrect,
	)

	return err
}

func (r *DefaultRepository) DeleteGameSessionAnswers(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) error {
	const query = `
		delete from easy_quizy_game_session 
		where game_id = $1 and player_id = $2
	`

	_, err := r.db(ctx).ExecContext(
		ctx,
		query,
		gameID,
		playerID,
	)

	return err
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
