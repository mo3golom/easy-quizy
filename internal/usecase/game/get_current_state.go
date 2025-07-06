package game

import (
	"context"
	"easy-quizy/internal/model"
	"errors"

	"github.com/google/uuid"
)

func (u *Usecase) GetCurrentState(ctx context.Context, gameID uuid.UUID, playerID uuid.UUID) (*model.State, error) {
	var result *model.State
	return result, u.trm.Do(ctx, func(ctx context.Context) error {
		// Получаем игру
		specificGame, err := u.Get(ctx, gameID)
		if err != nil {
			return err
		}

		// Получаем сессию игрока
		specificSession, err := u.games.GetGameSession(ctx, gameID, playerID)
		if err != nil {
			return err
		}

		// Собираем map[QuestionID]AnswerID для быстрых проверок
		answeredMap := make(map[int64]int64)
		for _, ans := range specificSession.Answers {
			answeredMap[ans.QuestionID] = ans.AnswerID
		}

		result = &model.State{
			GameID: gameID,
			Progress: model.Progress{
				Total:    int64(len(specificGame.Questions)),
				Answered: int64(len(specificSession.Answers)),
			},
		}

		// Ищем следующий вопрос
		nextQuestionFound := false
		for _, item := range specificGame.Questions {
			if _, ok := answeredMap[item.ID]; ok {
				continue
			}

			// Неотвеченный вопрос найден
			result.Question = &specificGame.Questions[item.ID]
			nextQuestionFound = true
			break
		}

		if nextQuestionFound {
			return nil
		}

		// Все вопросы отвечены, считаем результат
		totalScore := int64(0)
		for _, ans := range specificSession.Answers {
			if int(ans.QuestionID) < 0 || int(ans.QuestionID) >= len(specificGame.Questions) {
				return errors.New("invalid question id in session answers")
			}
			question := specificGame.Questions[ans.QuestionID]
			found := false
			for _, opt := range question.AnswerOptions {
				if opt.ID == ans.AnswerID {
					score := int64(1)
					if opt.Score != nil {
						score = *opt.Score
					}
					totalScore += score
					found = true
					break
				}
			}
			if !found {
				return errors.New("answer option not found for question")
			}
		}

		// Находим подходящий результат
		var resultText string
		for _, res := range specificGame.ScoreResults {
			if totalScore >= res.From && totalScore <= res.To {
				resultText = res.Text
				break
			}
		}
		if resultText == "" {
			return errors.New("no result found for total score")
		}
		result.Result = &model.Result{
			TotalScore: &totalScore,
			ResultText: resultText,
		}

		return nil
	})
}
