import type { QuizData, SavedResult } from '../types';

export function getResultText(score: number, resultMapping: Record<string, string>): string {
	for (const [range, text] of Object.entries(resultMapping)) {
		const [min, max] = range.split('-').map(Number);

		if (min === undefined || max === undefined || isNaN(min) || isNaN(max)) {
			continue; // Пропускаем некорректные диапазоны
		}
		
		if (score >= min && score <= max) {
			return text;
		}
	}
	return 'Unknown result';
}

export function saveQuizResult(gameId: string, quizName: string, totalScore: number, resultText: string): void {
	const result: SavedResult = {
		quizName,
		totalScore,
		resultText,
		completedAt: new Date().toISOString()
	};
	localStorage.setItem(`quiz-result-${gameId}`, JSON.stringify(result));
}

export function loadSavedResult(gameId: string): SavedResult | null {
	try {
		const saved = localStorage.getItem(`quiz-result-${gameId}`);
		return saved ? JSON.parse(saved) : null;
	} catch {
		return null;
	}
}

export function clearSavedResult(gameId: string): void {
	localStorage.removeItem(`quiz-result-${gameId}`);
}

export function getDefaultScore(): number {
	return 7;
}