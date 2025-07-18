<script lang="ts">
	import { goto } from "$app/navigation";
	import { getDailyGame } from "$lib/api/client";
	let gameId = "";
	let isLoading = false;
	let error = "";

	// Daily quiz state variables
	let isDailyLoading = false;
	let dailyError = "";

	async function startGame() {
		if (!gameId.trim()) {
			error = "Пожалуйста, введите ID игры";
			return;
		}

		error = "";
		isLoading = true;

		try {
			// Просто переходим к игре, проверка ID будет на странице игры
			await goto(`/game/${gameId.trim()}`);
		} catch (err) {
			error = "Произошла ошибка при переходе к игре";
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

	function clearError() {
		if (error) {
			error = "";
		}
	}

	function handleFocus() {
		// Focus handling can be added here if needed
	}

	function handleBlur() {
		// Blur handling can be added here if needed
	}

	// Daily quiz functions
	async function startDailyGame() {
		dailyError = "";
		isDailyLoading = true;

		try {
			const response = await getDailyGame();
			await goto(`/game/${response.gameId}`);
		} catch (err) {
			dailyError = "Произошла ошибка при загрузке ежедневного квиза";
			console.error("Daily game error:", err);
		} finally {
			isDailyLoading = false;
		}
	}

	function clearDailyError() {
		if (dailyError) {
			dailyError = "";
		}
	}
</script>

<svelte:head>
	<title>Quiz App - Введите ID игры</title>
	<meta name="description" content="Введите ID игры для начала квиза" />
</svelte:head>

<div class="min-h-screen container mx-auto p-4 max-w-2xl">
	<div class="text-center mb-10">
		<h1 class="text-4xl sm:text-5xl font-bold text-primary text-main-font">
			Easy Quizy!
		</h1>
	</div>

	<!-- Daily Quiz Section -->
	<section
		class="card bg-base-100 card-border border-primary shadow-xl mb-6"
		aria-labelledby="daily-quiz-heading"
	>
		<div class="flex flex-row">
			<!-- Left section: Daily quiz description -->
			<div
				class="basis-2/3 justify-center bg-primary outline outline-2 outline-primary rounded-xl p-4 sm:p-6 shadow-lg"
			>
				<div class="text-primary-content">
					<h2
						id="daily-quiz-heading"
						class="text-xl sm:text-2xl lg:text-3xl font-bold mb-2 sm:mb-4 text-main-font"
					>
						Ежедневный квиз
					</h2>
					<p class="text-base sm:text-lg lg:text-xl mb-2 sm:mb-3">
						Новый квиз каждый день!
					</p>
				</div>
			</div>

			<!-- Right section: Play button -->
			<div class="flex flex-col items-center p-4 sm:p-6 lg:flex-1">
				<button
					class="btn btn-accent btn-lg text-base sm:text-lg font-semibold transition-transform w-35 h-35 mask mask-squircle flex flex-col items-center justify-center"
					class:loading={isDailyLoading}
					disabled={isDailyLoading}
					on:click={() => {
						clearDailyError();
						startDailyGame();
					}}
					on:keydown={(e) => {
						if (e.key === "Enter" || e.key === " ") {
							e.preventDefault();
							if (!isDailyLoading) {
								clearDailyError();
								startDailyGame();
							}
						}
					}}
					aria-label={isDailyLoading
						? "Загружается ежедневный квиз"
						: "Начать ежедневный квиз"}
					aria-describedby={dailyError
						? "daily-error-message"
						: "daily-quiz-heading"}
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

				<!-- Daily quiz error message -->
				{#if dailyError}
					<div class="mt-3 animate-fade-in" role="alert">
						<div class="alert alert-error shadow-lg">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="stroke-current shrink-0 h-5 w-5 sm:h-6 sm:w-6"
								fill="none"
								viewBox="0 0 24 24"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
								></path>
							</svg>
							<span
								id="daily-error-message"
								class="text-sm font-medium"
								aria-live="polite"
							>
								{dailyError}
							</span>
						</div>
					</div>
				{/if}
			</div>
		</div>
	</section>
	<div class="divider"></div>
	<section
		class="card bg-base-100 card-border border-primary shadow-xl"
		aria-labelledby="game-id-heading"
	>
		<div class="card-body p-0">
			<div
				class="bg-primary outline outline-2 outline-primary rounded-xl p-4 sm:p-6 shadow-lg"
			>
				<div class="form-control w-full">
					<label for="game-id" class="label">
						<span
							id="game-id-heading"
							class="text-base sm:text-lg lg:text-xl text-primary-content text-main-font"
						>
							Введите ID игры
						</span>
					</label>
					<div class="relative mt-1">
						<input
							id="game-id"
							type="text"
							placeholder="c886247f-9dc6-4e3f-b2f0-50f912438079"
							class="input input-lg input-bordered w-full text-sm sm:text-base"
							class:input-error={error}
							bind:value={gameId}
							on:keypress={handleKeyPress}
							on:input={clearError}
							on:focus={handleFocus}
							on:blur={handleBlur}
							disabled={isLoading}
							aria-describedby={error
								? "game-id-error-message"
								: "game-id-help"}
							aria-invalid={error ? "true" : "false"}
						/>
					</div>
					<div
						id="game-id-help"
						class="mt-2 text-primary-content/70 flex items-center gap-1 text-xs sm:text-sm"
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
					<!-- Error Message -->
					{#if error}
						<div class="mt-3 animate-fade-in" role="alert">
							<div class="alert alert-error shadow-lg">
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="stroke-current shrink-0 h-5 w-5 sm:h-6 sm:w-6"
									fill="none"
									viewBox="0 0 24 24"
									aria-hidden="true"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
									></path>
								</svg>
								<span
									id="game-id-error-message"
									class="text-sm font-medium"
									aria-live="polite"
								>
									{error}
								</span>
							</div>
						</div>
					{/if}
				</div>
			</div>
			<div class="p-4 sm:p-6">
				<button
					class="btn btn-accent text-base sm:text-lg font-semibold h-14 sm:h-16 w-full transition-transform hover:scale-105 focus:scale-105"
					class:loading={isLoading}
					disabled={!gameId.trim() || isLoading}
					on:click={startGame}
					on:keydown={(e) => {
						if (e.key === "Enter" || e.key === " ") {
							e.preventDefault();
							if (gameId.trim() && !isLoading) {
								startGame();
							}
						}
					}}
					aria-label={isLoading
						? "Загружается игра"
						: "Начать игру с введенным ID"}
					aria-describedby={error
						? "game-id-error-message"
						: "game-id-heading"}
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
			</div>
		</div>
	</section>
</div>
