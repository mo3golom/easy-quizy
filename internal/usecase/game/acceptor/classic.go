package acceptor

import (
	"errors"
	"quizzly-v2/internal/contracts"
	"quizzly-v2/internal/model"
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
	for id := range correctAnswers {
		correctAnswersMap[int64(id)] = struct{}{}
	}

	_, ok := correctAnswersMap[answers[0]]
	return &contracts.AcceptAnswersOut{
		IsCorrect:   ok,
		Explanation: question.Explanation,
	}, nil
}
