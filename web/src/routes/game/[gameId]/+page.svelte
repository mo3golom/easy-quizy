<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { apiQuizState } from '$lib/stores/quiz';
	import { getGameState, submitAnswer, resetGame, hasQuestion, hasResult } from '$lib/api/client';
	import ApiQuizHeader from '$lib/components/ApiQuizHeader.svelte';
	import ApiQuizQuestion from '$lib/components/ApiQuizQuestion.svelte';
	import ApiQuizResult from '$lib/components/ApiQuizResult.svelte';

	const displayResultDuration = 3000;
	
	let mounted = false;
	let autoProgressTimeout: NodeJS.Timeout | null = null;
	let countdownInterval: NodeJS.Timeout | null = null;
	let countdown = 0;
	let gameId: string;
	let error = '';
	let isInitialLoading = true;

	$: gameId = $page.params.gameId == undefined ? '' : $page.params.gameId.trim();

	onMount(() => {
		loadGameState();
		mounted = true;

		// Cleanup timeouts on unmount
		return () => {
			if (autoProgressTimeout) {
				clearTimeout(autoProgressTimeout);
			}
			if (countdownInterval) {
				clearInterval(countdownInterval);
			}
		};
	});

	async function loadGameState() {
		try {
			isInitialLoading = true;
			error = '';
			
			const gameState = await getGameState(gameId);
			
			if (hasQuestion(gameState)) {
				apiQuizState.update(state => ({
					...state,
					gameId,
					gameName: gameState.gameInfo.title,
					currentQuestion: {
						id: gameState.question.ID,
						text: gameState.question.text,
						...(gameState.question.image_id !== undefined ? { image: gameState.question.image_id } : {}),
						options: gameState.question.answer_options.map(opt => ({
							id: opt.id,
							text: opt.answer
						}))
					},
					progress: gameState.progress,
					result: null,
					isComplete: false,
					selectedOption: null,
					isLoading: false,
					showResult: false,
					answerResult: null
				}));
			} else if (hasResult(gameState)) {
				apiQuizState.update(state => ({
					...state,
					gameId,
					gameName: gameState.gameInfo.title,
					currentQuestion: null,
					progress: gameState.progress,
					result: {
						totalScore: gameState.result.total_score,
						resultText: gameState.result.result_text
					},
					isComplete: true,
					selectedOption: null,
					isLoading: false,
					showResult: false,
					answerResult: null
				}));
			}
		} catch (err) {
			console.error('Failed to load game state:', err);
			error = 'Не удалось загрузить игру. Проверьте ID игры.';
		} finally {
			isInitialLoading = false;
		}
	}

	function startCountdown(milliseconds: number) {
		countdown = Math.floor(milliseconds / 1000);
		countdownInterval = setInterval(() => {
			countdown--;
			if (countdown <= 0 && countdownInterval) {
				clearInterval(countdownInterval);
				countdownInterval = null;
			}
		}, 1000);
	}

	function clearCountdown() {
		if (countdownInterval) {
			clearInterval(countdownInterval);
			countdownInterval = null;
		}
		countdown = 0;
	}

	async function handleOptionSelect(event: CustomEvent<number>) {
		const optionId = event.detail;
		
		apiQuizState.update(state => ({
			...state,
			selectedOption: optionId,
			isLoading: true
		}));

		try {
			// Отправляем ответ на сервер
			const currentQuestion = $apiQuizState.currentQuestion;
			if (!currentQuestion) return;

			const answerResponse = await submitAnswer(gameId, currentQuestion.id, optionId);
			
			// Обновляем состояние с результатом ответа
			setTimeout(() => {
				apiQuizState.update(state => ({
					...state,
					isLoading: false,
					showResult: true,
					answerResult: answerResponse
				}));

				// Запускаем таймер для автоматического перехода
				startCountdown(displayResultDuration);
				autoProgressTimeout = setTimeout(() => {
					handleNext();
				}, displayResultDuration);
			}, 1000); // Симуляция задержки для UX
		} catch (err) {
			console.error('Failed to submit answer:', err);
			apiQuizState.update(state => ({
				...state,
				isLoading: false
			}));
			error = 'Не удалось отправить ответ. Попробуйте еще раз.';
		}
	}

	async function handleNext() {
		clearCountdown();
		if (autoProgressTimeout) {
			clearTimeout(autoProgressTimeout);
			autoProgressTimeout = null;
		}

		try {
			// Загружаем новое состояние игры
			const gameState = await getGameState(gameId);
			
			if (hasQuestion(gameState)) {
				// Есть следующий вопрос
				apiQuizState.update(state => ({
					...state,
					gameName: gameState.gameInfo.title,
					currentQuestion: {
						id: gameState.question.ID,
						text: gameState.question.text,
						...(gameState.question.image_id !== undefined ? { image: gameState.question.image_id } : {}),
						options: gameState.question.answer_options.map(opt => ({
							id: opt.id,
							text: opt.answer
						}))
					},
					progress: gameState.progress,
					selectedOption: null,
					showResult: false,
					answerResult: null
				}));
			} else if (hasResult(gameState)) {
				// Игра завершена
				handleComplete(gameState);
			}
		} catch (err) {
			console.error('Failed to load next question:', err);
			error = 'Не удалось загрузить следующий вопрос.';
		}
	}

	function handleComplete(finalGameState?: any) {
		clearCountdown();
		if (autoProgressTimeout) {
			clearTimeout(autoProgressTimeout);
			autoProgressTimeout = null;
		}

		if (finalGameState && hasResult(finalGameState)) {
			apiQuizState.update(state => ({
				...state,
				currentQuestion: null,
				progress: finalGameState.progress,
				result: {
					totalScore: finalGameState.result.total_score,
					resultText: finalGameState.result.result_text
				},
				isComplete: true,
				selectedOption: null,
				showResult: false,
				answerResult: null
			}));
		}
	}

	async function handleRestart() {
		clearCountdown();
		if (autoProgressTimeout) {
			clearTimeout(autoProgressTimeout);
			autoProgressTimeout = null;
		}

		try {
			// Сбрасываем игру через API
			await resetGame(gameId);
			
			// Перезагружаем страницу
			window.location.reload();
		} catch (err) {
			console.error('Failed to restart game:', err);
			// В случае ошибки просто перезагружаем страницу
			window.location.reload();
		}
	}

	$: currentQuestionNumber = $apiQuizState.progress.answered + 1;
</script>

<svelte:head>
	<title>Quiz App - Игра {gameId.slice(0, 8)}...</title>
	<meta name="description" content="Интерактивное квиз-приложение" />
</svelte:head>

{#if isInitialLoading}
	<div class="min-h-screen flex items-center justify-center">
		<div class="text-center">
			<span class="loading loading-spinner loading-lg text-primary mb-4"></span>
			<p class="text-lg text-base-content/70">Загрузка игры...</p>
		</div>
	</div>
{:else if error}
	<div class="min-h-screen flex items-center justify-center">
		<div class="text-center max-w-md">
			<div class="text-6xl mb-4">❌</div>
			<h1 class="text-2xl font-bold text-error mb-4">Ошибка</h1>
			<p class="text-base-content/70 mb-6">{error}</p>
			<button class="btn btn-primary" on:click={() => goto('/')}>
				<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
				</svg>
				Вернуться на главную
			</button>
		</div>
	</div>
{:else if $apiQuizState.isComplete}
	<ApiQuizResult state={$apiQuizState} on:restart={handleRestart} />
{:else if $apiQuizState.currentQuestion}
	<div class="min-h-screen">
		<!-- Header -->
		<div class="px-4 pt-4 pb-2">
			<div class="container mx-auto max-w-4xl">
				<ApiQuizHeader 
					gameName={$apiQuizState.gameName}
					currentQuestionNumber={currentQuestionNumber}
					totalQuestions={$apiQuizState.progress.total}
					answeredQuestions={$apiQuizState.progress.answered}
				/>
			</div>
		</div>
		
		<!-- Question content -->
		<div class="px-4 pb-4">
			<div class="container mx-auto max-w-4xl">
				<ApiQuizQuestion
					state={$apiQuizState}
					countdown={countdown}
					on:selectOption={handleOptionSelect}
					on:next={() => handleNext()}
					on:complete={() => handleComplete()}
				/>
			</div>
		</div>
	</div>
{:else}
	<div class="min-h-screen flex items-center justify-center">
		<div class="text-center">
			<h1 class="text-2xl font-bold text-error mb-4">Неизвестное состояние игры</h1>
			<p class="text-base-content/70">Попробуйте перезагрузить страницу</p>
		</div>
	</div>
{/if}