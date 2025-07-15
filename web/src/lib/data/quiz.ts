import type { QuizData } from '../types';

export const defaultQuizData: QuizData = {
	id: '801b737c-0197-4999-b6f1-41d281ed808a',
	name: "Classic",
	description: "Classic quiz",
	questions: [
		{
			question: "What is the capital of France?",
			image: "https://i.pinimg.com/736x/7b/de/e5/7bdee5a30249a67d9c65187a8417937a.jpg",
			options: [
				{
					text: "Paris",
					score: 2,
					isCorrect: true
				},
				{
					text: "London",
					score: 1,
					isCorrect: false
				},
				{
					text: "Berlin",
					score: 0,
					isCorrect: false
				}
			],
			explanation: "Paris is the capital of France"
		},
		{
			question: "What is the capital of Germany?",
			options: [
				{
					text: "Paris",
					score: 0,
					isCorrect: false
				},
				{
					text: "London",
					score: 1,
					isCorrect: false
				},
				{
					text: "Berlin",
					score: 2,
					isCorrect: true
				}
			],
			explanation: "Berlin is the capital of Germany"
		}
	],
	result: {
		"0-1": "you know nothing",
		"2-3": "you know a little", 
		"4-5": "you know a lot"
	}
};

// Simulated database of quizzes
const quizDatabase: Record<string, QuizData> = {
	'801b737c-0197-4999-b6f1-41d281ed808a': defaultQuizData
};

export function getQuizById(id: string): QuizData | null {
	return quizDatabase[id] || null;
}

export function getAllQuizzes(): QuizData[] {
	return Object.values(quizDatabase);
}