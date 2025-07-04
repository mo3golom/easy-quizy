package game

import (
	"context"
	"errors"
	"quizzly-v2/internal/contracts"
)

func (u *Usecase) AcceptAnswer(ctx context.Context, in *contracts.AcceptAnswersIn) (*contracts.AcceptAnswersOut, error) {
	var result *contracts.AcceptAnswersOut
	return result, u.trm.Do(ctx, func(ctx context.Context) error {
		specificGame, err := u.Get(ctx, in.GameID)
		if err != nil {
			return err
		}

		if len(specificGame.Questions) == 0 {
			return errors.New("question not found")
		}

		acceptor, ok := u.acceptors[specificGame.Type]
		if !ok {
			return errors.New("game type is not supported")
		}

		result, err = acceptor.Accept(&specificGame.Questions[in.QuestionID], []int64{in.Answer})
		return err
	})
}
