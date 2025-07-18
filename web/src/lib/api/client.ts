import { browser } from '$app/environment';
import { PUBLIC_API_BASE_URL } from '$env/static/public';

// Генерация и сохранение Player ID
function getOrCreatePlayerId(): string {
	if (!browser) return '';
	
	let playerId = localStorage.getItem('player-id');
	if (!playerId) {
		playerId = crypto.randomUUID();
		localStorage.setItem('player-id', playerId);
	}
	return playerId;
}

// Базовая функция для API запросов
async function apiRequest(endpoint: string, options: RequestInit = {}): Promise<Response> {
	const playerId = getOrCreatePlayerId();
	
	const response = await fetch(`${PUBLIC_API_BASE_URL}${endpoint}`, {
		...options,
		headers: {
			'X-Player-ID': playerId,
			'Content-Type': 'application/json',
			...options.headers,
		},
	});

	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}

	return response;
}

// Типы для API ответов
export interface ApiQuestion {
	ID: number;
	text: string;
	image_id?: string;
	answer_options: {
		id: number;
		answer: string;
	}[];
}

export interface ApiProgress {
	answered: number;
	correct: number;
	total: number;
}

export interface GameInfo {
	id: string;
	title: string;
}

export interface ApiGameStateWithQuestion {
	question: ApiQuestion;
	progress: ApiProgress;
	gameInfo: GameInfo;
}

export interface ApiGameStateWithResult {
	result: {
		total_score: number;
		result_text: string;
	};
	progress: ApiProgress;
	gameInfo: GameInfo;
}

export type ApiGameState = ApiGameStateWithQuestion | ApiGameStateWithResult;

export interface ApiAnswerResponse {
	isCorrect: boolean;
	explanation?: string;
}

export interface DailyGameResponse {
	gameId: string;
}

// Проверка типа ответа
export function hasQuestion(state: ApiGameState): state is ApiGameStateWithQuestion {
	return 'question' in state;
}

export function hasResult(state: ApiGameState): state is ApiGameStateWithResult {
	return 'result' in state;
}

// API методы
export async function getGameState(gameId: string): Promise<ApiGameState> {
	const response = await apiRequest(`/api/game/${gameId}`);
	return response.json();
}

export async function resetGame(gameId: string): Promise<void> {
	await apiRequest(`/api/game/${gameId}/reset`, {
		method: 'GET',
	});
}
export async function submitAnswer(gameId: string, questionId: number, answerId: number): Promise<ApiAnswerResponse> {
	const response = await apiRequest(`/api/game/${gameId}/accept-answer`, {
		method: 'POST',
		body: JSON.stringify({
			questionId: questionId,
			answerId: answerId,
		}),
	});
	return response.json();
}

export async function getDailyGame(): Promise<DailyGameResponse> {
	const response = await apiRequest('/api/game/daily');
	return response.json();
}