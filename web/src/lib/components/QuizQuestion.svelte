<script lang="ts">
	import type { QuizQuestion, QuizProgress } from "../types";
	import {
		triggerFeedback,
		triggerFeedbackFunction,
	} from "$lib/actions/feedback";
	import CorrectAnswerEffect from "$lib/components/CorrectAnswerEffect.svelte";
	import WrongAnswerEffect from "$lib/components/WrongAnswerEffect.svelte";
	import CountdownTimer from "$lib/components/CountdownTimer.svelte";
	import { submitAnswer } from "$lib/api/client";
	import { toast } from "$lib/toast";

	const RIGHT_ANSWER = "Правильно!";
	const WRONG_ANSWER = "Неправильно!";
	const NEXT_QUESTION_AFTER = "Следующий вопрос";
	const SHOW_RESULT_DURATION_SECONDS = 2;
	const WAITING_ANSWER_DURATION_SECONDS = 30;
	const DEFAULT_ANSWER_OPTION_ID = -1;

	interface Props {
		gameId: string;
		question: QuizQuestion;
		progress: QuizProgress;
		onnext?: () => void;
	}
	interface AnswerResult {
		isCorrect: boolean;
		explanation?: string;
	}

	let { gameId, question, progress, onnext }: Props = $props();

	let selectedOption: number | null = $state(null);
	let answerResult: AnswerResult | null = $state(null);
	let isLoading: boolean = $state(false);
	
	const isLastQuestion = $derived(progress.answered === progress.total - 1);
	const currentQuestionNumber = $derived(
		progress.total - (progress.total - progress.answered) + 1,
	);
	const progressPercent = $derived(
		progress.total > 0 ? (currentQuestionNumber / progress.total) * 100 : 0,
	);

	// Reset state and scroll to top when question changes
	$effect(() => {
		if (question) {
			// Reset state for new question
			selectedOption = null;
			answerResult = null;
			isLoading = false;
			window.scrollTo({ top: 0, behavior: "smooth" });
		}
	});

	async function handleOptionSelect(optionId: number) {
		if (isLoading || answerResult !== null) return;

		selectedOption = optionId;
		isLoading = true;

		try {
			const answerResponse = await submitAnswer(
				gameId,
				question.id,
				optionId,
			);

			if (answerResponse.isCorrect) {
				triggerFeedbackFunction("success");
			} else {
				triggerFeedbackFunction("error");
			}

			isLoading = false;
			answerResult = answerResponse;
		} catch (err) {
			console.error("Failed to submit answer:", err);
			isLoading = false;
			toast.error("Не удалось отправить ответ. Попробуйте еще раз.");
		}
	}

	function handleNext() {
		onnext?.();
	}
</script>

<div class="p-2 relative pt-8">
	<div
		class="absolute -top-2 left-1/2 -translate-x-1/2 z-10 bg-primary text-primary-content flex items-center gap-2 p-2 pl-4 pr-4 rounded-full translate-y-5 -rotate-2 text-main-font text-xl"
	>
		<progress
			class="progress progress-neutral w-56"
			value={progressPercent}
			max="100"
		></progress>
		{currentQuestionNumber}&nbsp;из&nbsp;{progress.total}
	</div>
	<div class="card bg-primary rounded-3xl">
		<div class="card-body p-4">
			<div class="text-primary-content relative">
				<div class="mb-4 mt-4">
					{#if question.image}
						<div class="mb-3 place-self-center">
							<img
								src={question.image}
								alt="Quiz question illustration"
								class="w-full mx-auto rounded-xl"
								loading="lazy"
							/>
						</div>
					{/if}
					<h2 class="text-2xl font-semibold">
						{question.text}
					</h2>
				</div>
			</div>
			<div class="space-y-3 items-stretch gap-4">
				{#each question.options as option, index}
					<button
						type="button"
						class={`relative btn-answer-option text-left break-words hyphens-auto`}
						class:hover={false}
						class:option-correct={answerResult !== null &&
							selectedOption === option.id &&
							answerResult.isCorrect}
						class:option-incorrect={answerResult !== null &&
							selectedOption === option.id &&
							!answerResult.isCorrect}
						onclick={() => {
							handleOptionSelect(option.id);
						}}
						use:triggerFeedback={"light"}
						disabled={isLoading || answerResult !== null}
					>
						{index + 1}.&nbsp;{option.text}
						{#if isLoading && (selectedOption === option.id || selectedOption === DEFAULT_ANSWER_OPTION_ID)}
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
		</div>
	</div>
</div>
<div class="sticky bottom-0 z-40">
	<div
		class="card bg-primary-content rounded-3xl rounded-b-none w-full"
	>
		<div class="card-body p-4">
			{#if answerResult !== null}
				{#if answerResult.isCorrect}
					<div
						class="mb-4 bg-success text-success-content p-4 rounded-lg"
					>
						<span class="text-xl">{RIGHT_ANSWER} 🎉</span>
					</div>
				{:else}
					<div
						class="mb-4 bg-error text-error-content p-4 rounded-lg"
					>
						<span class="text-xl">{WRONG_ANSWER} 😞</span>
						{#if answerResult?.explanation}
							<p class="text-sm mt-1 opacity-90">
								{answerResult?.explanation}
							</p>
						{/if}
					</div>
				{/if}
				<div class="flex items-center justify-between gap-4">
					<div class="flex items-center gap-2 text-base-content/70">
						<span class="text-sm">{NEXT_QUESTION_AFTER}</span>
						<CountdownTimer
							durationSeconds={SHOW_RESULT_DURATION_SECONDS}
							onComplete={handleNext}
							view={"compact"}
						/>
					</div>

					<button
						class="btn btn-primary transition-transform"
						onclick={handleNext}
						use:triggerFeedback={"light"}
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
			{:else}
				<div class="flex items-center justify-end gap-4">
					<CountdownTimer
						durationSeconds={WAITING_ANSWER_DURATION_SECONDS}
						onComplete={() => {
							handleOptionSelect(DEFAULT_ANSWER_OPTION_ID);
						}}
					/>
				</div>
			{/if}
		</div>
	</div>
</div>
{#if answerResult !== null}
	{#if answerResult?.isCorrect}
		<CorrectAnswerEffect duration={SHOW_RESULT_DURATION_SECONDS} />
	{:else}
		<WrongAnswerEffect duration={SHOW_RESULT_DURATION_SECONDS} />
	{/if}
{/if}
