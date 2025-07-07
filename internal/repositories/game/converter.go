package game

import (
	"easy-quizy/internal/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Вспомогательные структуры для парсинга JSON
type rawGame struct {
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Type        string            `json:"type"`
	Questions   []rawQuestion     `json:"questions"`
	Result      map[string]string `json:"result"`
}

type rawQuestion struct {
	Question    string            `json:"question"`
	Image       *string           `json:"image"`
	Explanation *string           `json:"explanation"`
	Options     []rawAnswerOption `json:"options"`
}

type rawAnswerOption struct {
	Text      string `json:"text"`
	Score     *int64 `json:"score"`
	IsCorrect bool   `json:"isCorrect"`
}

func parseScoreResults(resultMap map[string]string) ([]model.ScoreResult, error) {
	var results []model.ScoreResult
	for k, v := range resultMap {
		parts := strings.Split(k, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid score range: %s", k)
		}
		from, err1 := strconv.ParseInt(parts[0], 10, 64)
		to, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid score range: %s", k)
		}
		results = append(results, model.ScoreResult{
			From: from,
			To:   to,
			Text: v,
		})
	}
	return results, nil
}

func rawToGame(rg rawGame) (model.Game, error) {
	scoreResults, err := parseScoreResults(rg.Result)
	if err != nil {
		return model.Game{}, err
	}
	var questions []model.Question
	for idxq, rq := range rg.Questions {
		var options []model.AnswerOption
		for idxao, ro := range rq.Options {
			options = append(options, model.AnswerOption{
				ID:        int64(idxao),
				Answer:    ro.Text,
				IsCorrect: ro.IsCorrect,
				Score:     ro.Score,
			})
		}
		questions = append(questions, model.Question{
			ID:            int64(idxq),
			Text:          rq.Question,
			ImageID:       rq.Image,
			Explanation:   rq.Explanation,
			AnswerOptions: options,
		})
	}
	return model.Game{
		Title:        rg.Name,
		Description:  rg.Description,
		Questions:    questions,
		ScoreResults: scoreResults,
	}, nil
}

func convertToGame(in sqlxGame) (model.Game, error) {
	var rg rawGame
	if err := json.Unmarshal(in.Payload, &rg); err != nil {
		return model.Game{}, err
	}
	game, err := rawToGame(rg)
	if err != nil {
		return model.Game{}, err
	}

	game.ID = in.ID
	game.Type = model.GameType(in.Type)
	return game, nil
}

func convertToSession(in []sqlxGameSession) model.GameSession {
	if len(in) == 0 {
		return model.GameSession{}
	}

	result := model.GameSession{
		GameID:   in[0].GameID,
		PlayerID: in[0].PlayerID,
		Answers:  make([]model.GameSessionAnswer, 0, len(in)),
	}

	for _, item := range in {
		result.Answers = append(result.Answers, model.GameSessionAnswer{
			QuestionID: item.QuestionID,
			AnswerID:   item.AnswerID,
			IsCorrect:  item.IsCorrect,
		})
	}

	return result
}
