package model

import "github.com/google/uuid"

const (
	GameTypeClassic = "classic"
	GameTypeDaily   = "daily"
)

type (
	GameType string

	Game struct {
		ID           uuid.UUID
		Type         GameType
		Title        string
		Description  *string
		Questions    []Question
		ScoreResults []ScoreResult
	}

	GameInfo struct {
		ID    uuid.UUID
		Title string
	}

	ScoreResult struct {
		From int64
		To   int64
		Text string
	}

	GameSession struct {
		GameID   uuid.UUID
		PlayerID uuid.UUID
		Answers  []GameSessionAnswer
	}

	GameSessionAnswer struct {
		QuestionID int64
		AnswerID   int64
		IsCorrect  bool
	}

	Question struct {
		ID            int64
		Text          string
		ImageID       *string
		Explanation   *string
		AnswerOptions []AnswerOption
	}

	AnswerOption struct {
		ID        int64
		Answer    string
		IsCorrect bool
		Score     *int64
	}

	State struct {
		Question *Question
		Result   *Result
		Progress Progress
		GameInfo GameInfo
	}

	Result struct {
		TotalScore int64
		ResultText string
	}

	Progress struct {
		Answered int64
		Total    int64
	}
)

func (q Question) GetCorrectAnswers() []AnswerOption {
	result := make([]AnswerOption, 0, len(q.AnswerOptions))
	for _, answer := range q.AnswerOptions {
		answer := answer
		if !answer.IsCorrect {
			continue
		}

		result = append(result, answer)
	}

	return result
}
