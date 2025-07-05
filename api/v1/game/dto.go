package game

import (
	"quizzly-v2/internal/model"
)

type AnswerOption struct {
	ID     int64  `json:"id"`
	Answer string `json:"answer"`
}

type Question struct {
	Text          string         `json:"text"`
	ImageID       *string        `json:"image_id,omitempty"`
	AnswerOptions []AnswerOption `json:"answer_options"`
}

type Result struct {
	TotalScore *int64 `json:"total_score,omitempty"`
	ResultText string `json:"result_text"`
}

type Progress struct {
	Answered int64 `json:"answered"`
	Total    int64 `json:"total"`
}

type StateResponse struct {
	Question *Question `json:"question,omitempty"`
	Result   *Result   `json:"result,omitempty"`
	Progress Progress  `json:"progress"`
}

func toStateResponse(state *model.State) *StateResponse {
	resp := &StateResponse{
		Progress: Progress{
			Answered: state.Progress.Answered,
			Total:    state.Progress.Total,
		},
	}

	if state.Question != nil {
		resp.Question = &Question{
			Text:          state.Question.Text,
			ImageID:       state.Question.ImageID,
			AnswerOptions: make([]AnswerOption, len(state.Question.AnswerOptions)),
		}
		for i, opt := range state.Question.AnswerOptions {
			resp.Question.AnswerOptions[i] = AnswerOption{
				ID:     int64(i),
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
