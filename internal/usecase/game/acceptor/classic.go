package acceptor

import (
	"easy-quizy/internal/contracts"
	"easy-quizy/internal/model"
	"errors"
)

type ClassicAcceptor struct{}

func NewClassicAcceptor() *ClassicAcceptor {
	return &ClassicAcceptor{}
}

func (a *ClassicAcceptor) Accept(question *model.Question, answers []int64) (*contracts.AcceptAnswersOut, error) {
	if len(answers) > 1 {
		return nil, errors.New("simple choice can't have multiple answers")
	}

	correctAnswers := question.GetCorrectAnswers()
	correctAnswersMap := make(map[int64]struct{}, len(correctAnswers))
	for _, item := range correctAnswers {
		correctAnswersMap[item.ID] = struct{}{}
	}

	_, ok := correctAnswersMap[answers[0]]
	return &contracts.AcceptAnswersOut{
		IsCorrect:   ok,
		Explanation: question.Explanation,
	}, nil
}
