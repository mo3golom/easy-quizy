<script lang="ts">
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";
	import { apiQuizState } from "$lib/stores/quiz";
	import {
		getGameState,
		submitAnswer,
		resetGame,
		hasQuestion,
		hasResult,
	} from "$lib/api/client";
	import QuizQuestion from "$lib/components/QuizQuestion.svelte";
	import QuizResult from "$lib/components/QuizResult.svelte";
	import BackButton from "$lib/components/BackButton.svelte";
	import Loading from "$lib/components/Loading.svelte";
	import { triggerFeedbackFunction } from "$lib/actions/feedback";

	const displayResultDuration = 2000;

	let autoProgressTimeout: NodeJS.Timeout | null = null;
	let countdownInterval: NodeJS.Timeout | null = null;
	let countdown = $state(0);
	let gameId: string = $derived($page.params.gameId == undefined ? "" : $page.params.gameId.trim());
	let error = $state("");
	let isInitialLoading = $state(true);

	
	let gameName = $derived($apiQuizState.gameName);
	let pageUrl = $derived($page.url.href);

	let ogTitle = $derived(gameName ? `Квиз: ${gameName}` : "Quiz App");
	let ogDescription = $derived(gameName
		? `Сыграй в квиз "${gameName}" и проверь свои знания!`
		: "Интерактивное квиз-приложение");

	onMount(() => {
		loadGameState();

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
			error = "";

			const gameState = await getGameState(gameId);

			if (hasQuestion(gameState)) {
				apiQuizState.update((state) => ({
					...state,
					gameId,
					gameName: gameState.gameInfo.title,
					currentQuestion: {
						id: gameState.question.ID,
						text: gameState.question.text,
						...(gameState.question.image_id !== undefined
							? { image: gameState.question.image_id }
							: {}),
						options: gameState.question.answer_options.map(
							(opt) => ({
								id: opt.id,
								text: opt.answer,
							}),
						),
					},
					progress: {
						answered: gameState.progress.answered,
						correct: gameState.progress.correct,
						total: gameState.progress.total,
					},
					result: null,
					isComplete: false,
					selectedOption: null,
					isLoading: false,
					showResult: false,
					answerResult: null,
				}));
			} else if (hasResult(gameState)) {
				apiQuizState.update((state) => ({
					...state,
					gameId,
					gameName: gameState.gameInfo.title,
					currentQuestion: null,
					progress: {
						answered: gameState.progress.answered,
						correct: gameState.progress.correct,
						total: gameState.progress.total,
					},
					result: {
						totalScore: gameState.result.total_score,
						resultText: gameState.result.result_text,
					},
					isComplete: true,
					selectedOption: null,
					isLoading: false,
					showResult: false,
					answerResult: null,
				}));
			}
		} catch (err) {
			console.error("Failed to load game state:", err);
			error = "Не удалось загрузить игру. Проверьте ID игры.";
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

	async function handleOptionSelect(optionId: number) {

		apiQuizState.update((state) => ({
			...state,
			selectedOption: optionId,
			isLoading: true,
		}));

		try {
			// Отправляем ответ на сервер
			const currentQuestion = $apiQuizState.currentQuestion;
			if (!currentQuestion) return;

			const answerResponse = await submitAnswer(
				gameId,
				currentQuestion.id,
				optionId,
			);

			if (answerResponse.isCorrect) {
				triggerFeedbackFunction("success")
			} else {
				triggerFeedbackFunction("error")
			}
			
			// Обновляем состояние с результатом ответа
			apiQuizState.update((state) => ({
				...state,
				isLoading: false,
				showResult: true,
				answerResult: answerResponse,
			}));

			// Запускаем таймер для автоматического перехода
			startCountdown(displayResultDuration);
			autoProgressTimeout = setTimeout(() => {
				handleNext();
			}, displayResultDuration);
		} catch (err) {
			console.error("Failed to submit answer:", err);
			apiQuizState.update((state) => ({
				...state,
				isLoading: false,
			}));
			error = "Не удалось отправить ответ. Попробуйте еще раз.";
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
				apiQuizState.update((state) => ({
					...state,
					gameName: gameState.gameInfo.title,
					currentQuestion: {
						id: gameState.question.ID,
						text: gameState.question.text,
						...(gameState.question.image_id !== undefined
							? { image: gameState.question.image_id }
							: {}),
						options: gameState.question.answer_options.map(
							(opt) => ({
								id: opt.id,
								text: opt.answer,
							}),
						),
					},
					progress: {
						answered: gameState.progress.answered,
						correct: gameState.progress.correct,
						total: gameState.progress.total,
					},
					selectedOption: null,
					showResult: false,
					answerResult: null,
				}));
			} else if (hasResult(gameState)) {
				// Игра завершена
				handleComplete(gameState);
			}
		} catch (err) {
			console.error("Failed to load next question:", err);
			error = "Не удалось загрузить следующий вопрос.";
		}
	}

	function handleComplete(finalGameState?: any) {
		clearCountdown();
		if (autoProgressTimeout) {
			clearTimeout(autoProgressTimeout);
			autoProgressTimeout = null;
		}

		if (finalGameState && hasResult(finalGameState)) {
			apiQuizState.update((state) => ({
				...state,
				currentQuestion: null,
				progress: {
					answered: finalGameState.progress.answered,
					correct: finalGameState.progress.correct,
					total: finalGameState.progress.total,
				},
				result: {
					totalScore: finalGameState.result.total_score,
					resultText: finalGameState.result.result_text,
				},
				isComplete: true,
				selectedOption: null,
				showResult: false,
				answerResult: null,
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
			console.error("Failed to restart game:", err);
			// В случае ошибки просто перезагружаем страницу
			window.location.reload();
		}
	}
</script>

<svelte:head>
	<title>{ogTitle}</title>
	<meta name="description" content={ogDescription} />

	<!-- Open Graph / Facebook -->
	<meta property="og:type" content="website" />
	<meta property="og:url" content={pageUrl} />
	<meta property="og:title" content={ogTitle} />
	<meta property="og:description" content={ogDescription} />
	<!-- <meta property="og:image" content="URL_TO_IMAGE"> -->

	<!-- Twitter -->
	<meta property="twitter:card" content="summary_large_image" />
	<meta property="twitter:url" content={pageUrl} />
	<meta property="twitter:title" content={ogTitle} />
	<meta property="twitter:description" content={ogDescription} />
	<!-- <meta property="twitter:image" content="URL_TO_IMAGE"> -->
</svelte:head>

<div class="container mx-auto max-w-md p-2">
	<BackButton targetPage="/" />

	{#if isInitialLoading}
		<Loading />
	{:else if error}
		<div class="min-h-screen flex items-center justify-center p-6">
			<div class="text-center max-w-md">
				<div class="text-6xl mb-4">❌</div>
				<h1 class="text-2xl font-bold text-error mb-4">Ошибка</h1>
				<p class="text-base-content/70 mb-6">{error}</p>

				<button
					class="btn btn-block btn-primary mb-4 transition-transform"
					onclick={() => window.location.reload()}
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke-width="1.5"
						stroke="currentColor"
						class="size-6"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99"
						/>
					</svg>
					Перезагрузить
				</button>
				<button
					class="btn btn-block btn-outline btn-primary transition-transform"
					onclick={() => goto("/")}
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke-width="1.5"
						stroke="currentColor"
						class="size-6"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							d="m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25"
						/>
					</svg>
					На главную
				</button>
			</div>
		</div>
	{:else if $apiQuizState.isComplete}
		<QuizResult state={$apiQuizState} onrestart={handleRestart} />
	{:else if $apiQuizState.currentQuestion}
			<QuizQuestion
				state={$apiQuizState}
				{countdown}
				onselectOption={handleOptionSelect}
				onnext={handleNext}
				oncomplete={handleComplete}
			/>
	{:else}
		<div class="min-h-screen flex items-center justify-center p-6">
			<div class="text-center">
				<h1 class="text-2xl font-bold text-error mb-4">
					Неизвестное состояние игры
				</h1>
				<p class="text-base-content/70">
					Попробуйте перезагрузить страницу
				</p>
			</div>
		</div>
	{/if}
</div>
