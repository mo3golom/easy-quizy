import { writable } from 'svelte/store';
import type { QuizState } from '../types';

// Новый store для API состояния
export const quizState = writable<QuizState>({
	gameId: '',
	gameName: '',
	currentQuestion: null,
	progress: { answered: 0, correct: 0, total: 0 },
	result: null,
	isComplete: false,
});