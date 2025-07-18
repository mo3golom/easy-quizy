package game

import (
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"
	"easy-quizy/internal/usecase/game/acceptor"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type (
	Acceptor interface {
		Accept(question *model.Question, answers []int64) (*contracts.AcceptAnswersOut, error)
	}

	Usecase struct {
		games repository
		trm   trm.Manager

		acceptors map[model.GameType]Acceptor
	}
)

func NewUsecase(
	games repository,
	trm trm.Manager,
) contracts.GameUsecase {
	return &Usecase{
		games: games,
		trm:   trm,
		acceptors: map[model.GameType]Acceptor{
			model.GameTypeClassic: acceptor.NewClassicAcceptor(),
		},
	}
}
