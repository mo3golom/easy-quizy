<script lang="ts">
    import { onMount } from "svelte";

    interface Props {
        startSeconds?: number;
        onComplete?: () => void;
    }

    let { startSeconds = 60, onComplete }: Props = $props();
    let remainingSeconds = $state(startSeconds);
    let interval: NodeJS.Timeout | null = null;

    let minutes = $derived(Math.floor(remainingSeconds / 60));
    let seconds = $derived(remainingSeconds % 60);
    let formattedTime = $derived(
        `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`,
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

<div
    class="flex items-center gap-2 badge badge-soft badge-primary text-primary py-4 px-3 rounded-lg"
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
            d="M12 6v6h4.5m4.5 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
        />
    </svg>
    <span class="text-xl font-bold font-mono">
        {formattedTime}
    </span>
</div>
