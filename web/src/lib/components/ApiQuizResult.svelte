<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ApiQuizState } from "../types";
	import { toast } from "$lib/toast";

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
           toast.success("Скопировано!");
		});
	}
</script>

<div class="container mx-auto max-w-md p-2">
	<div class="relative mt-14">
	<div class="absolute -top-10 left-1/2 -translate-x-1/2 z-10 bg-primary text-primary-content flex items-center gap-2 p-2 pl-4 pr-4 rounded-full translate-y-5 -rotate-2 text-main-font text-xl">
		Твой&nbsp;результат&nbsp;в&nbsp;игре 
	</div>
	<div class="card bg-base-100 bg-primary rounded-3xl mt-14">
		<!-- Two main sections structure with improved desktop layout -->
			<!-- Left section: Result display -->
			<div class="flex flex flex-col items-center text-primary-content text-center h-full p-4">
				{#if state.result}
					<h2 class="text-5xl font-bold mb-6 mt-4 text-main-font">"{state.gameName}"</h2>
					<div class="w-40 h-40 md:w-48 md:h-48 mask mask-squircle bg-gradient-to-b from-black/50 to-primary flex flex-col items-center justify-center mb-6">
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
					<div class="text-xl md:text-2xl">
						{state.result.resultText}
					</div>
				{/if}
			</div>

			<!-- Right section: Action buttons -->
			<div class="flex-1 flex flex-col justify-center gap-4 p-4">
				<!-- Share button -->
				<button 
					class="btn btn-lg btn-secondary transition-transform" 
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
					class="btn btn-neutral btn-lg transition-transform"
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
			</div>
	</div>
	</div>
</div>
