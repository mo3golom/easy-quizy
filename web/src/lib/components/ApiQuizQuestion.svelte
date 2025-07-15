<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ApiQuizState } from "../types";

	const optionColors = [
		"btn-warning",
		"btn-secondary",
		"btn-accent",
		"btn-info",
		"btn-success",
	];

	export let state: ApiQuizState;
	export let countdown: number = 0;

	const dispatch = createEventDispatcher<{
		selectOption: number;
		next: void;
		complete: void;
	}>();

	function handleOptionSelect(optionId: number) {
		if (state.isLoading || state.showResult) return;
		dispatch("selectOption", optionId);
	}

	function handleNext() {
		if (state.isComplete) {
			dispatch("complete");
		} else {
			dispatch("next");
		}
	}

	function getOptionClass(index: number, optionId: number): string {
		let baseClass =
			"btn btn-outline btn-block justify-start h-auto p-3 normal-case h-full flex flex-col text-lg justify-center";

		baseClass += ` ${optionColors[index % optionColors.length]}`;

		if (state.showResult && state.selectedOption === optionId) {
			if (state.answerResult?.isCorrect) {
				baseClass += " option-correct";
			} else {
				baseClass += " option-incorrect";
			}
		}

		return baseClass;
	}

	$: currentQuestion = state.currentQuestion;
	$: isLastQuestion = state.progress.answered === state.progress.total - 1;
	$: showExplanation =
		state.showResult &&
		state.answerResult &&
		!state.answerResult.isCorrect &&
		state.answerResult.explanation;
</script>

<div class="relative">
	<!-- Main Quiz Card -->
	<div class="card bg-base-100 card-border border-base-300 shadow-xl">
		<div class="card-body p-6">
			<div>
				{#if currentQuestion?.image}
					<div class="mb-4">
						<div
							class="flex flex-col md:flex-row md:gap-4 md:items-start"
						>
							<div class="md:flex-shrink-0 mb-3 md:mb-0">
								<img
									src={currentQuestion.image}
									alt="Quiz question illustration"
									class="w-full md:w-64 max-w-sm mx-auto md:mx-0 shadow-lg mask mask-squircle"
									loading="lazy"
								/>
							</div>
							<div class="md:flex-1">
								<h2
									class="text-xl md:text-2xl font-semibold leading-relaxed"
								>
									{currentQuestion.text}
								</h2>
							</div>
						</div>
					</div>
				{:else if currentQuestion}
					<h2
						class="text-xl md:text-2xl font-semibold mb-4 leading-relaxed"
					>
						{currentQuestion.text}
					</h2>
				{/if}
			</div>

			<!-- Options -->
			{#if currentQuestion}
				<div
					class="space-y-3 md:space-y-0 md:grid gap-4 items-stretch"
					class:md:grid-cols-2={currentQuestion.options.length === 2}
					class:md:grid-cols-3={currentQuestion.options.length === 3}
					class:md:grid-cols-4={currentQuestion.options.length === 4}
				>
					{#each currentQuestion.options as option, index}
						<button
							class={getOptionClass(index, option.id) +
								" "}
							on:click={() => handleOptionSelect(option.id)}
							disabled={state.isLoading || state.showResult}
						>
							<span
								class="flex items-center justify-between w-full"
							>
								<span class="text-left break-words hyphens-auto"
									>{option.text}</span
								>
								{#if state.isLoading && state.selectedOption === option.id}
									<span
										class="loading loading-spinner loading-sm"
									></span>
								{:else if state.showResult && state.selectedOption === option.id}
									{#if state.answerResult?.isCorrect}
										<svg
											xmlns="http://www.w3.org/2000/svg"
											class="h-5 w-5 flex-shrink-0"
											fill="none"
											viewBox="0 0 24 24"
											stroke="currentColor"
										>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												stroke-width="2"
												d="M5 13l4 4L19 7"
											/>
										</svg>
									{:else}
										<svg
											xmlns="http://www.w3.org/2000/svg"
											class="h-5 w-5 flex-shrink-0"
											fill="none"
											viewBox="0 0 24 24"
											stroke="currentColor"
										>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												stroke-width="2"
												d="M6 18L18 6M6 6l12 12"
											/>
										</svg>
									{/if}
								{/if}
							</span>
						</button>
					{/each}
				</div>
			{/if}
		</div>
	</div>

	<!-- Floating Timer and Next Button -->
	{#if state.showResult && countdown > 0}
		<div class="sticky mt-4 bottom-4 z-40">
			<div
				class="card bg-base-100 card-border border-base-300 shadow-xl w-full mx-auto"
			>
				<div class="card-body p-6">
					{#if showExplanation}
						<div role="alert" class="alert alert-soft mb-4">
							<div class="flex items-start gap-3">
								<svg
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
									class="stroke-current shrink-0 w-6 h-6 mt-0.5"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
									></path>
								</svg>
								<span class="text-sm leading-relaxed"
									>{state.answerResult?.explanation}</span
								>
							</div>
						</div>
					{/if}
					<div class="flex items-center justify-between gap-4">
						<div
							class="flex items-center gap-2 text-base-content/70"
						>
							<span class="text-sm">Следующий вопрос через</span>
							<span class="countdown text-primary font-bold">
								<span
									style={`--value:${countdown};`}
									aria-live="polite"
									aria-label={countdown.toString()}
									>{countdown}</span
								>
							</span>
						</div>
						<button
							class="btn btn-primary btn-sm"
							on:click={handleNext}
						>
							{#if isLastQuestion}
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-4 w-4 mr-1"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"
									/>
								</svg>
								Завершить
							{:else}
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-4 w-4 mr-1"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M13 7l5 5m0 0l-5 5m5-5H6"
									/>
								</svg>
								Далее
							{/if}
						</button>
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
