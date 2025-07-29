<script lang="ts">
	import { goto } from "$app/navigation";
	import { getDailyGame } from "$lib/api/client";
	import { toast } from "$lib/toast";
	import { triggerFeedback } from "$lib/actions/feedback";

	let gameId = $state("");
	let isLoading = $state(false);
	let isDailyLoading = $state(false);

	async function startGame() {
		if (!gameId.trim()) {
			toast.error("Пожалуйста, введите ID игры");
			return;
		}

		isLoading = true;

		try {
			// Просто переходим к игре, проверка ID будет на странице игры
			await goto(`/game/${gameId.trim()}`);
		} catch (err) {
			toast.error("Произошла ошибка при переходе к игре");
			console.error("Navigation error:", err);
		} finally {
			isLoading = false;
		}
	}

	function handleKeyPress(event: KeyboardEvent) {
		if (event.key === "Enter") {
			startGame();
		}
	}

	// Daily quiz functions
	async function startDailyGame() {
		isDailyLoading = true;

		try {
			const response = await getDailyGame();
			await goto(`/game/${response.gameId}`);
		} catch (err) {
			toast.error("Произошла ошибка при загрузке ежедневного квиза");
			console.error("Daily game error:", err);
		} finally {
			isDailyLoading = false;
		}
	}
</script>

<svelte:head>
	<title>Quiz App - Введите ID игры</title>
	<meta name="description" content="Введите ID игры для начала квиза" />
</svelte:head>

<div class="w-full min-h-screen bg-primary">
	<div class="container mx-auto max-w-md p-2">
		<div class="text-center p-6">
			<h1
				class="text-5xl font-bold text-primary-content text-main-font -rotate-5"
			>
				<p class="-translate-x-8">ИЗИ</p>
				<p class="translate-x-8">КВИЗИ!</p>
			</h1>
		</div>

		<div class="p-4 bg-primary-content rounded-3xl">
			<!-- Daily Quiz Section -->
			<section class="mb-6" aria-labelledby="daily-quiz-heading">
				<div class="flex flex-row">
					<!-- Left section: Daily quiz description -->
					<div class="basis-2/3 justify-center content-center">
						<div class="text-primary">
							<h2
								id="daily-quiz-heading"
								class="text-3xl md:text-4xl font-bold text-main-font"
							>
								<p
									class="bg-primary text-primary-content w-min p-2 pl-4 pr-4 rounded-full mx-auto -rotate-5 translate-y-2"
								>
									Ежедневный
								</p>
								<p
									class="bg-primary text-primary-content w-min p-2 pl-4 pr-4 rounded-full mx-auto"
								>
									квиз
								</p>
							</h2>
						</div>
					</div>

					<!-- Right section: Play button -->
					<div class="flex flex-col items-center lg:flex-1">
						<button
							class="btn btn-secondary btn-lg text-base sm:text-lg font-semibold transition-transform w-30 h-30 md:w-35 md:h-35 mask mask-squircle flex flex-col items-center justify-center"
							class:loading={isDailyLoading}
							disabled={isDailyLoading}
							onclick={startDailyGame}
							onkeydown={(e) => {
								if (e.key === "Enter" || e.key === " ") {
									e.preventDefault();
									if (!isDailyLoading) {
										startDailyGame();
									}
								}
							}}
							use:triggerFeedback={"light"}
							aria-label={isDailyLoading
								? "Загружается ежедневный квиз"
								: "Начать ежедневный квиз"}
							aria-describedby="daily-quiz-heading"
						>
							{#if isDailyLoading}
								<span
									class="loading loading-spinner loading-md"
									aria-hidden="true"
								></span>
								Загрузка...
							{:else}
								<svg
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
									stroke-width="1.5"
									stroke="currentColor"
									class="size-12 align-center"
									aria-hidden="true"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986V5.653Z"
									/>
								</svg>
							{/if}
						</button>
					</div>
				</div>
			</section>
			<div class="divider text-main-font text-xl text-base-content">
				ИЛИ
			</div>
			<section aria-labelledby="game-id-heading">
				<div class="form-control w-full">
					<label for="game-id" class="label">
						<span
							id="game-id-heading"
							class="sm:text-lg lg:text-xl text-main-font text-base-content"
						>
							Введите ID игры
						</span>
					</label>
					<div class="relative mt-1">
						<input
							id="game-id"
							type="text"
							placeholder="c886247f-9dc6-4e3f-b2f0-50f912438079"
							class="input input-lg input-bordered w-full text-sm"
							bind:value={gameId}
							onkeypress={handleKeyPress}
							disabled={isLoading}
							aria-describedby="game-id-help"
							aria-invalid="false"
						/>
					</div>
					<div
						id="game-id-help"
						class="mt-1 text-base-content/70 flex items-center gap-1 text-xs sm:text-sm"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							class="size-3 sm:size-4 flex-shrink-0"
							aria-hidden="true"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z"
							></path>
						</svg>
						<span>Получите ID игры от организатора квиза</span>
					</div>
				</div>

				<button
					class="btn btn-secondary text-base sm:text-lg font-semibold h-14 sm:h-16 w-full transition-transform mt-4"
					class:loading={isLoading}
					disabled={!gameId.trim() || isLoading}
					onclick={startGame}
					onkeydown={(e) => {
						if (e.key === "Enter" || e.key === " ") {
							e.preventDefault();
							if (gameId.trim() && !isLoading) {
								startGame();
							}
						}
					}}
					use:triggerFeedback={"light"}
					aria-label={isLoading
						? "Загружается игра"
						: "Начать игру с введенным ID"}
					aria-describedby="game-id-heading"
				>
					{#if isLoading}
						<span
							class="loading loading-spinner loading-md"
							aria-hidden="true"
						></span>
						<span class="sr-only">Загружается</span>
						Загрузка...
					{:else}
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							class="size-5 sm:size-6"
							aria-hidden="true"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986V5.653Z"
							></path>
						</svg>
						Начать игру
					{/if}
				</button>
			</section>
		</div>
	</div>
</div>
