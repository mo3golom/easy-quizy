export interface QuizOption {
	text: string;
	score?: number;
	isCorrect: boolean;
}

export interface QuizQuestion {
	question: string;
	image?: string;
	options: QuizOption[];
	explanation?: string;
}

export interface QuizData {
	id: string;
	name: string;
	description: string;
	questions: QuizQuestion[];
	result: Record<string, string>;
}

export interface QuizState {
	currentQuestionIndex: number;
	answers: number[];
	totalScore: number;
	isComplete: boolean;
	selectedOption: number | null;
	isLoading: boolean;
	showResult: boolean;
}

export interface SavedResult {
	quizName: string;
	totalScore: number;
	resultText: string;
	completedAt: string;
}

// Новые типы для API
export interface ApiQuizState {
	gameId: string;
	gameName: string;
	currentQuestion: {
		id: number;
		text: string;
		image?: string;
		options: {
			id: number;
			text: string;
		}[];
	} | null;
	progress: {
		answered: number;
		total: number;
	};
	result: {
		totalScore: number;
		resultText: string;
	} | null;
	isComplete: boolean;
	selectedOption: number | null;
	isLoading: boolean;
	showResult: boolean;
	answerResult: {
		isCorrect: boolean;
		explanation?: string;
	} | null;
}