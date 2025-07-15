<script lang="ts">
	export let gameName: string;
	export let currentQuestionNumber: number;
	export let totalQuestions: number;
	export let answeredQuestions: number;
	
	$: progress = totalQuestions > 0 ? (answeredQuestions / totalQuestions) * 100 : 0;
</script>

<div class="quiz-header">
	<div class="flex justify-between items-start">
		<h1 class="text-2xl lg:text-3xl font-bold text-primary text-main-font">{gameName}</h1>
		<div class="text-right">
			<div class="text-sm text-base-content/70 mb-1">
				Вопрос {currentQuestionNumber} из {totalQuestions}
			</div>
			<div class="w-24 bg-base-300 rounded-full h-1.5">
				<div 
					class="bg-primary h-1.5 rounded-full transition-all duration-300 ease-out"
					style="width: {progress}%"
				></div>
			</div>
		</div>
	</div>
</div>

<style>
	.quiz-header {
		animation: slideUp 0.5s ease-out;
	}
</style>