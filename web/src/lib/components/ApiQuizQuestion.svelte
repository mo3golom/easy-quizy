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
	const NEXT_QUESTION_AFTER = "Следующий вопрос";

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

	function getOptionCorrectClass(index: number, optionId: number): string {
		let baseClass = ``;

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

<div class="container mx-auto max-w-md p-2">
	<div class="relative mt-8">
		<div
			class="absolute -top-10 left-1/2 -translate-x-1/2 z-10 bg-primary text-primary-content flex items-center gap-2 p-2 pl-4 pr-4 rounded-full translate-y-5 -rotate-2 text-main-font text-xl"
		>
			<progress
				class="progress progress-neutral w-56"
				value={progress}
				max="100"
			></progress>
			{currentQuestionNumber}&nbsp;из&nbsp;{state.progress.total}
		</div>
		<div class="card bg-primary rounded-3xl">
			<div class="card-body p-4">
				{#if currentQuestion}
					<div class="text-primary-content relative">
						<div class="mb-4 mt-4">
							{#if currentQuestion?.image}
								<div class="mb-3 place-self-center">
									<img
										src={currentQuestion.image}
										alt="Quiz question illustration"
										class="w-full mx-auto rounded-xl"
										loading="lazy"
									/>
								</div>
							{/if}
							<h2 class="text-2xl font-semibold">
								{currentQuestion?.text}
							</h2>
						</div>
					</div>
					<div class="space-y-3 items-stretch gap-4">
						{#each currentQuestion.options as option, index}
							<button
								class={`relative btn-answer-option ${getOptionCorrectClass(index, option.id)}`}
								on:click={() => handleOptionSelect(option.id)}
								disabled={state.isLoading || state.showResult}
							>
								<span class="text-left break-words hyphens-auto"
									>{index + 1}.&nbsp;{option.text}</span
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

		{#if state.showResult || countdown > 0}
			<div class="sticky mt-4 bottom-0 z-40">
				<div
					class="card bg-primary-content w-full mx-auto rounded-3xl rounded-b-none"
				>
					<div class="card-body p-4">
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
						<div class="flex items-center justify-between gap-4">
							<div
								class="flex items-center gap-2 text-base-content/70"
							>
								<span class="text-sm"
									>{NEXT_QUESTION_AFTER}</span
								>
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
