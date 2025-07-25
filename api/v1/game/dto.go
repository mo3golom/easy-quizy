package game

import (
	"easy-quizy/internal/model"

	"github.com/google/uuid"
)

type AnswerOption struct {
	ID     int64  `json:"id"`
	Answer string `json:"answer"`
}

type Question struct {
	ID            int64
	Text          string         `json:"text"`
	ImageID       *string        `json:"image_id,omitempty"`
	AnswerOptions []AnswerOption `json:"answer_options"`
}

type Result struct {
	TotalScore int64  `json:"total_score"`
	ResultText string `json:"result_text"`
}

type Progress struct {
	Answered int64 `json:"answered"`
	Total    int64 `json:"total"`
}

type GameInfo struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type StateResponse struct {
	Question *Question `json:"question,omitempty"`
	Result   *Result   `json:"result,omitempty"`
	Progress Progress  `json:"progress"`
	GameInfo GameInfo  `json:"gameInfo"`
}

type AcceptAnswerRequest struct {
	QuestionID int64 `json:"questionId"`
	AnswerID   int64 `json:"answerId"`
}

type AcceptAnswerResponse struct {
	IsCorrect   bool    `json:"isCorrect"`
	Explanation *string `json:"explanation,omitempty"`
}

type GetDailyGameResponse struct {
	GameID string `json:"gameId"`
}

func toStateResponse(state model.State) StateResponse {
	resp := StateResponse{
		Progress: Progress{
			Answered: state.Progress.Answered,
			Total:    state.Progress.Total,
		},
		GameInfo: GameInfo{
			ID:    state.GameInfo.ID,
			Title: state.GameInfo.Title,
		},
	}

	if state.Question != nil {
		resp.Question = &Question{
			ID:            state.Question.ID,
			Text:          state.Question.Text,
			ImageID:       state.Question.ImageID,
			AnswerOptions: make([]AnswerOption, len(state.Question.AnswerOptions)),
		}
		for i, opt := range state.Question.AnswerOptions {
			resp.Question.AnswerOptions[i] = AnswerOption{
				ID:     opt.ID,
				Answer: opt.Answer,
			}
		}
	}

	if state.Result != nil {
		resp.Result = &Result{
			TotalScore: state.Result.TotalScore,
			ResultText: state.Result.ResultText,
		}
	}

	return resp
}
