<script lang="ts">
	import { goto } from "$app/navigation";

	let gameId = "";
	let isLoading = false;
	let error = "";
	let isFocused = false;

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
		isFocused = true;
	}

	function handleBlur() {
		isFocused = false;
	}
</script>

<svelte:head>
	<title>Quiz App - Введите ID игры</title>
	<meta name="description" content="Введите ID игры для начала квиза" />
</svelte:head>

<div class="min-h-screen">
	<div class="container mx-auto p-4 max-w-2xl pb-24 sm:pb-0">
		<div class="text-center mb-10">
			<h1
				class="text-4xl sm:text-5xl lg:text-6xl font-bold text-primary mb-4 text-main-font"
			>
				Quiz App
			</h1>
		</div>

		<!-- Main Form Card -->
		<div class="card bg-base-100 card-border border-base-300 shadow-xl">
			<div class="card-body p-6">
				<div class="form-control w-full">
					<label for="game-id">
						<span class="text-lg sm:text-xl">ID игры</span>
					</label>

					<!-- Enhanced Form Input Group -->
					<div class="relative mt-1">
						<input
							id="game-id"
							type="text"
							placeholder="Введите ID игры..."
							class="input input-lg input-bordered w-full"
							class:input-error={error}
							class:border-primary={isFocused && !error}
							bind:value={gameId}
							on:keypress={handleKeyPress}
							on:input={clearError}
							on:focus={handleFocus}
							on:blur={handleBlur}
							disabled={isLoading}
						/>
					</div>

					<!-- Error Message -->
					{#if error}
						<div class="mt-3 animate-fade-in">
							<div class="alert alert-error shadow-lg">
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="stroke-current shrink-0 h-6 w-6"
									fill="none"
									viewBox="0 0 24 24"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
									/>
								</svg>
								<span class="text-sm font-medium">{error}</span>
							</div>
						</div>
					{/if}
				</div>
			</div>
		</div>

		<!-- Footer Info / Start Game Button - Fixed on mobile, normal on desktop -->
		<div class="fixed bottom-0 left-0 right-0 p-4 sm:static sm:text-center sm:mt-8 sm:p-0">
			<button
				class="btn btn-primary text-lg font-semibold h-16 w-full sm:w-auto transition-all duration-600 ease-in-out transform hover:scale-105 shadow-lg sm:shadow-none"
				class:loading={isLoading}
				disabled={!gameId.trim() || isLoading}
				on:click={startGame}
			>
				{#if !gameId.trim()}
					<div class="p-1.5 bg-base-content/20 rounded-full">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-4 w-4 text-base-content"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
							/>
						</svg>
					</div>
					Получите ID игры от организатора квиза
				{:else if isLoading}
					<span class="loading loading-spinner loading-md"></span>
					Загрузка...
				{:else}
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
							d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986V5.653Z"
						/>
					</svg>
					Начать игру
				{/if}
			</button>
		</div>
	</div>
</div>
