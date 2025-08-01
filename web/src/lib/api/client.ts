import { env } from '$env/dynamic/public';
import { toast } from '$lib/toast';
import { userData } from "$lib/stores/user";
import type { User } from "../types"
import { get } from 'svelte/store';

// Функция для получения текущего playerId
function getUser(): User {
	return get(userData);
}

// Базовая функция для API запросов
async function apiRequest(endpoint: string, options: RequestInit = {}): Promise<Response> {
	const baseUrl = env.PUBLIC_API_BASE_URL || '';
	const currentPlayer = getUser();

	try {
		const response = await fetch(`${baseUrl}${endpoint}`, {
			...options,
			headers: {
				'X-Player-ID': currentPlayer.id.toString(),
				'X-Source': currentPlayer.source,
				'Content-Type': 'application/json',
				...options.headers,
			},
		});

		if (!response.ok) {
			// Handle different API status error types
			let errorMessage: string;

			switch (response.status) {
				case 400:
					errorMessage = 'Неверный запрос. Проверьте введенные данные.';
					break;
				case 401:
					errorMessage = 'Ошибка авторизации. Попробуйте обновить страницу.';
					break;
				case 403:
					errorMessage = 'Доступ запрещен.';
					break;
				case 404:
					errorMessage = 'Запрашиваемый ресурс не найден.';
					break;
				case 429:
					errorMessage = 'Слишком много запросов. Попробуйте позже.';
					break;
				case 500:
					errorMessage = 'Ошибка сервера. Попробуйте позже.';
					break;
				case 502:
				case 503:
				case 504:
					errorMessage = 'Сервер временно недоступен. Попробуйте позже.';
					break;
				default:
					errorMessage = `Ошибка API: ${response.status} ${response.statusText}`;
			}

			// Trigger error toast for API status errors
			toast.error(errorMessage);

			// Maintain existing error throwing behavior for component error handling
			throw new Error(`API request failed: ${response.status} ${response.statusText}`);
		}

		return response;
	} catch (error) {
		// Re-throw the error to maintain existing error handling behavior
		throw error;
	}
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