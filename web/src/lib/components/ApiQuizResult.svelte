<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ApiQuizState } from "../types";

	export let state: ApiQuizState;

	const dispatch = createEventDispatcher<{
		restart: void;
	}>();

	let isRestarting = false;
	let shareButtonText = "Поделиться";

	async function handleRestart() {
		isRestarting = true;
		dispatch("restart");
	}

	function handleShare() {
		if (!state.result) return;

		const shareText = `Мой результат ${state.result.totalScore} в игре "${state.gameName}". Сыграй и ты!`;
		const gameUrl = window.location.href;

		navigator.clipboard.writeText(`${shareText}\n${gameUrl}`).then(() => {
			shareButtonText = "Скопировано!";
			setTimeout(() => {
				shareButtonText = "Поделиться";
			}, 2000);
		});
	}
</script>

<div class="min-h-screen container mx-auto p-4 max-w-2xl">
	<div class="text-center mb-10">
			<h1
				class="text-4xl sm:text-5xl font-bold text-primary mb-4 text-main-font"
			>
				{#if state.gameName}
			Игра "{state.gameName}" завершена!
		{:else}
			Игра завершена!
		{/if}
			</h1>
		</div>
	<div class="card bg-base-100 card-border border-primary shadow-xl">
		<!-- Two main sections structure with improved desktop layout -->
		<div class="flex flex-col grow md:flex-row">
			<!-- Left section: Result display -->
			<div class="flex flex flex-col items-center text-primary-content text-center md:items-center md:justify-center h-full bg-primary outline outline-2 outline-primary rounded-xl p-6 shadow-lg">
				{#if state.result}
					<h2 class="text-2xl md:text-3xl font-bold mb-6 text-main-font">Твой результат</h2>
					<div class="w-40 h-40 md:w-48 md:h-48 mask mask-squircle bg-gradient-to-b from-neutral/50 to-primary flex flex-col items-center justify-center mb-6">
						<!-- Score display -->
						<div class="text-center">
							<div class="text-4xl md:text-5xl font-bold text-main-font">{state.result.totalScore}</div>
							<div>
								<span class="text-xl md:text-2xl font-bold text-main-font">из</span>
								<span class="text-2xl md:text-3xl font-bold text-main-font">{state.progress.total}</span>
							</div>
						</div>
					</div>
					
					<!-- Result text below the circle - placeholder for next task -->
					<div class="text-xl md:text-2xl text-main-font">
						{state.result.resultText}
					</div>
				{/if}
			</div>

			<!-- Right section: Action buttons -->
			<div class="flex-1 flex flex-col justify-center gap-4 md:gap-6 p-4 md:p-6">
				<!-- Share button -->
				<button 
					class="btn btn-lg btn-accent transition-transform" 
					on:click={handleShare}
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
							d="M7.217 10.907a2.25 2.25 0 1 0 0 2.186m0-2.186c.18.324.283.696.283 1.093s-.103.77-.283 1.093m0-2.186 9.566-5.314m-9.566 7.5 9.566 5.314m0 0a2.25 2.25 0 1 0 3.935 2.186 2.25 2.25 0 0 0-3.935-2.186Zm0-12.814a2.25 2.25 0 1 0 3.933-2.185 2.25 2.25 0 0 0-3.933 2.185Z"
						/>
					</svg>
					{shareButtonText}
				</button>

				<!-- Restart button -->
				<button
					class="btn btn-primary btn-lg transition-transform"
					class:loading={isRestarting}
					disabled={isRestarting}
					on:click={handleRestart}
				>
					{#if isRestarting}
						<span class="loading loading-spinner loading-md"></span>
						Перезапуск...
					{:else}
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-6 w-6"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
							/>
						</svg>
						Еще раз
					{/if}
				</button>

				<!-- Home button -->
				<a 
					href="/" 
					class="btn btn-lg btn-secondary transition-transform"
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
				</a>
			</div>
		</div>
	</div>
</div>
