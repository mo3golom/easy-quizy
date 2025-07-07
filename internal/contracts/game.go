package contracts

import (
	"context"
	"easy-quizy/internal/model"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrQuestionQueueIsEmpty     = errors.New("question queue is empty")
	ErrNotActiveSessionNotFound = errors.New("player's active session not found")
	ErrSessionNotFound          = errors.New("player's session not found")
	ErrSessionNotFinished       = errors.New("player's session not finished")
	ErrGameNotFound             = errors.New("game not found")
	ErrEmptyQuestions           = errors.New("empty questions")
	ErrEmptyAnswerOptions       = errors.New("empty answer options")
)

type (
	AcceptAnswersIn struct {
		GameID     uuid.UUID
		PlayerID   uuid.UUID
		QuestionID int64
		Answer     int64
	}

	AcceptAnswersOut struct {
		IsCorrect   bool
		Explanation *string
	}

	GameUsecase interface {
		Get(ctx context.Context, id uuid.UUID) (model.Game, error)
		GetDaily(ctx context.Context) (model.Game, error)
		AcceptAnswer(ctx context.Context, in *AcceptAnswersIn) (*AcceptAnswersOut, error)
		GetCurrentState(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) (model.State, error)
		Reset(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) error
	}
)
