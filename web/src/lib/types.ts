export type Source = 'browser' | 'telegram' | string;
export interface User {
	id: number;
	language_code?: string;
	username?: string;
	source: Source;
	chatId?: number | null;
	chatType?: string | null;
}

export interface QuizOption {
	id: number;
	text: string;
}

export interface QuizQuestion {
	id: number;
	text: string;
	image?: string;
	options: QuizOption[];
}

export interface QuizProgress {
	answered: number;
	correct: number;
	total: number;
};

export interface SavedResult {
	quizName: string;
	totalScore: number;
	resultText: string;
	completedAt: string;
}

// Новые типы для API
export interface QuizState {
	gameId: string;
	gameName: string;
	currentQuestion: QuizQuestion | null;
	progress: QuizProgress;
	result: {
		totalScore: number;
		resultText: string;
	} | null;
	isComplete: boolean;
}