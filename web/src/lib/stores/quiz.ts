import { writable } from 'svelte/store';
import type { QuizData, QuizState, SavedResult, ApiQuizState } from '../types';

export const quizData = writable<QuizData | null>(null);
export const quizState = writable<QuizState>({
	currentQuestionIndex: 0,
	answers: [],
	totalScore: 0,
	isComplete: false,
	selectedOption: null,
	isLoading: false,
	showResult: false
});

export const savedResult = writable<SavedResult | null>(null);

// Новый store для API состояния
export const apiQuizState = writable<ApiQuizState>({
	gameId: '',
	gameName: '',
	currentQuestion: null,
	progress: { answered: 0, correct: 0, total: 0 },
	result: null,
	isComplete: false,
	selectedOption: null,
	isLoading: false,
	showResult: false,
	answerResult: null
});