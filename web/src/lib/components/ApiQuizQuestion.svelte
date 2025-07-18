<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { ApiQuizState } from "../types";

	const ANSWER_OPTION_COLORS = [
		"btn-warning",
		"btn-secondary",
		"btn-accent",
		"btn-info",
		"btn-success",
	];
	const RIGHT_ANSWER = "Правильно!";
	const WRONG_ANSWER = "Неправильно!";
	const NEXT_QUESTION_AFTER = "Следующий вопрос через";

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
		let baseClass = `${ANSWER_OPTION_COLORS[index % ANSWER_OPTION_COLORS.length]}`;

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

	$: currentQuestionNumber =
		state.progress.total -
		(state.progress.total - state.progress.answered) +
		1;
	$: progress =
		state.progress.total > 0
			? (currentQuestionNumber / state.progress.total) * 100
			: 0;
</script>

<div class="px-4 pt-4 pb-2">
	<div class="container mx-auto max-w-4xl">
		<div class="flex-grow text-center">
			<h1
				class="text-4xl sm:text-5xl font-bold text-primary text-main-font"
			>
				{state.gameName}
			</h1>
		</div>
	</div>
</div>
<div class="px-4 pb-4">
	<div class="container mx-auto max-w-4xl">
		<div class="relative mt-14">
			<div class="absolute -top-10 left-1/2 -translate-x-1/2 z-10">
				<div
					class="radial-progress bg-primary text-primary-content border border-4 border-base-200"
					style="--value:{progress};"
					aria-valuenow={progress}
					role="progressbar"
				>
					{currentQuestionNumber} из {state.progress.total}
				</div>
			</div>
			<div class="card bg-base-100 card-border border-primary shadow-xl">
				<div class="card-body p-0">
					{#if currentQuestion}
						<div
							class="bg-primary text-primary-content outline outline-2 outline-primary rounded-xl p-6 pt-14 shadow-lg relative"
						>
							<button
								class="absolute top-2 right-2 btn btn-sm btn-ghost text-primary-content hover:bg-primary-content/20 rounded-full"
								on:click={() => (window.location.href = "/")}
								title="На главную"
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-5 w-5"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
									/>
								</svg>
							</button>
							<div class="mb-4 mt-4">
								<div
									class="flex flex-col md:flex-row md:gap-4 md:items-start"
								>
									{#if currentQuestion?.image}
										<div
											class="md:flex-shrink-0 mb-3 md:mb-0 place-self-center md:place-self-start max-w-xs"
										>
											<img
												src={currentQuestion.image}
												alt="Quiz question illustration"
												class="w-full md:w-64 max-w-sm mx-auto md:mx-0 shadow-lg mask mask-squircle"
												loading="lazy"
											/>
										</div>
									{/if}
									<div class="md:flex-1">
										<h2
											class="text-2xl font-semibold leading-relaxed"
										>
											{currentQuestion?.text}
										</h2>
									</div>
								</div>
							</div>
						</div>
						<div
							class="space-y-3 md:space-y-0 md:grid items-stretch gap-4 md:gap-6 p-4 md:p-6"
							class:md:grid-cols-2={currentQuestion.options
								.length === 2}
							class:md:grid-cols-3={currentQuestion.options
								.length === 3}
							class:md:grid-cols-4={currentQuestion.options
								.length === 4}
						>
							{#each currentQuestion.options as option, index}
								<button
									class={`relative btn btn-block justify-start h-auto p-3 normal-case h-full flex flex-col text-lg justify-center duration-300 ease-in-out transform hover:scale-105 transition-all ${getOptionClass(index, option.id)}`}
									on:click={() =>
										handleOptionSelect(option.id)}
									disabled={state.isLoading ||
										state.showResult}
								>
									<span
										class="text-left break-words hyphens-auto"
										>{option.text}</span
									>
									{#if state.isLoading && state.selectedOption === option.id}
										<div
											class="absolute top-0 left-0 w-full h-full flex items-center justify-center backdrop-blur-sm"
										>
											<span
												class="ml-1 loading loading-spinner loading-sm"
											></span>
										</div>
									{/if}
								</button>
							{/each}
						</div>
					{/if}
				</div>
			</div>

			{#if state.showResult && countdown > 0}
				<div class="sticky mt-4 bottom-4 z-40">
					<div
						class="card bg-base-100 card-border border-primary shadow-xl w-full mx-auto"
					>
						<div class="card-body p-6">
							{#if state.answerResult?.isCorrect}
								<div
									class="mb-4 bg-success text-success-content p-4 rounded-lg"
								>
									<span class="text-xl">
										{RIGHT_ANSWER}
									</span>
								</div>
							{:else}
								<div
									class="mb-4 bg-error text-error-content p-4 rounded-lg"
								>
									<span class="text-xl">
										{WRONG_ANSWER}
									</span>
									{#if showExplanation}
										<div
											role="alert"
											class="alert alert-dash border-neutral/50 mt-4"
										>
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
												<span
													class="text-sm leading-relaxed"
													>{state.answerResult
														?.explanation}</span
												>
											</div>
										</div>
									{/if}
								</div>
							{/if}
							<div
								class="flex items-center justify-between gap-4"
							>
								<div
									class="flex items-center gap-2 text-base-content/70"
								>
									<span class="text-sm"
										>{NEXT_QUESTION_AFTER}</span
									>
									<span
										class="countdown text-primary font-bold"
									>
										<span
											style={`--value:${countdown};`}
											aria-live="polite"
											aria-label={countdown.toString()}
											>{countdown}</span
										>
									</span>
								</div>
								<button
									class="btn btn-primary transition-transform"
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
											></path>
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
											></path>
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
	</div>
</div>
