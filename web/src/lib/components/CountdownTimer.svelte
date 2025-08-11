<script lang="ts">
    import { onMount } from "svelte";

    interface Props {
        durationSeconds?: number;
        onComplete?: () => void;
        view?: "default" | "compact";
    }

    let {
        durationSeconds = 60,
        onComplete,
        view = "default",
    }: Props = $props();
    let remainingSeconds = $state(durationSeconds);
    let interval: NodeJS.Timeout | null = null;

    // Проверка на максимум 99 секунд для compact режима
    $effect(() => {
        if (view === "compact" && durationSeconds > 99) {
            console.warn("Compact timer supports maximum 99 seconds");
            remainingSeconds = 99;
        }
    });

    let formattedTime = $derived(
        `${remainingSeconds.toString().padStart(2, "0")}`,
    );

    function startTimer() {
        if (interval) return;

        interval = setInterval(() => {
            if (remainingSeconds > 0) {
                remainingSeconds--;
            } else {
                stopTimer();
                onComplete?.();
            }
        }, 1000);
    }

    function stopTimer() {
        if (interval) {
            clearInterval(interval);
            interval = null;
        }
    }

    onMount(() => {
        startTimer();

        return () => stopTimer();
    });
</script>

{#if view === "compact"}
    <!-- Compact Timer - только секунды, без фона и иконки -->
    <span class="text-lg font-mono font-bold text-primary">
        {formattedTime}
    </span>
{:else}
    <!-- Default Timer - полный вид с иконкой и фоном -->
    <div
        class="flex items-center gap-3 badge badge-soft badge-primary text-primary py-4 px-4 rounded-lg"
    >
        <!-- Clock Icon -->
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
                d="M12 6v6h4.5m4.5 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
            />
        </svg>

        <!-- DaisyUI Countdown Display -->
        <span class="text-xl font-mono font-bold">
            {formattedTime}
        </span>
    </div>
{/if}
