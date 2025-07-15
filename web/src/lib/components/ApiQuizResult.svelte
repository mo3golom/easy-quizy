<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ApiQuizState } from "../types";

	export let state: ApiQuizState;

	const dispatch = createEventDispatcher<{
		restart: void;
	}>();

	let isRestarting = false;

	async function handleRestart() {
		isRestarting = true;
		dispatch("restart");
	}
</script>

<div class="min-h-screen p-4 max-w-2xl mx-auto">
	<div class="text-primary text-3xl sm:text-4xl font-bold mb-4 text-main-font">
		{#if state.gameName}
			Игра "{state.gameName}" завершена!
		{:else}
			Игра завершена!
		{/if}
	</div>
	<div class="card bg-base-100 card-border border-base-300 shadow-xl">
		<div class="card-body p-6 text-center">
			{#if state.result}
				<div
					class="stats stats-vertical w-full mb-2 bg-accent text-accent-content"
				>
					<div class="stat">
						<div class="stat-value text-main-font">
							{state.result.resultText}
						</div>
					</div>
					<div class="stat">
						<div
							class="stat-title text-accent-content"
						>
							Общий счет
						</div>
						<div class="stat-value">
							{state.result.totalScore}
						</div>
					</div>
				</div>
			{/if}

			<button
				class="btn btn-primary btn-lg"
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
						class="h-6 w-6 mr-2"
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
					Пройти еще раз
				{/if}
			</button>
		</div>
	</div>
</div>
