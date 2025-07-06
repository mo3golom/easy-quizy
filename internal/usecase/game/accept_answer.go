package game

import (
	"context"
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"
	"errors"
)

func (u *Usecase) AcceptAnswer(ctx context.Context, in *contracts.AcceptAnswersIn) (*contracts.AcceptAnswersOut, error) {
	var result *contracts.AcceptAnswersOut
	return result, u.trm.Do(ctx, func(ctx context.Context) error {
		// Получаем игру
		specificGame, err := u.Get(ctx, in.GameID)
		if err != nil {
			return err
		}

		if len(specificGame.Questions) == 0 {
			return errors.New("question not found")
		}

		// Находим вопрос по ID
		var question *model.Question
		for i := range specificGame.Questions {
			if specificGame.Questions[i].ID == in.QuestionID {
				question = &specificGame.Questions[i]
				break
			}
		}
		if question == nil {
			return errors.New("question not found")
		}

		// Получаем сессию игрока
		session, err := u.games.GetGameSession(ctx, in.GameID, in.PlayerID)
		if err != nil {
			return err
		}

		// Проверяем, отвечал ли игрок на этот вопрос
		for _, ans := range session.Answers {
			if ans.QuestionID != in.QuestionID {
				continue
			}

			result = &contracts.AcceptAnswersOut{
				IsCorrect:   ans.IsCorrect,
				Explanation: question.Explanation,
			}
			return nil
		}

		acceptor, ok := u.acceptors[specificGame.Type]
		if !ok {
			return errors.New("game type is not supported")
		}

		result, err = acceptor.Accept(question, []int64{in.Answer})
		if err != nil {
			return err
		}

		result.Explanation = question.Explanation

		return u.games.InsertGameSessionAnswer(
			ctx,
			in.GameID,
			in.PlayerID,
			model.GameSessionAnswer{
				QuestionID: in.QuestionID,
				AnswerID:   in.Answer,
				IsCorrect:  result.IsCorrect,
			},
		)
	})
}
