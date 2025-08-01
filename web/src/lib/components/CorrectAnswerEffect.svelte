<script lang="ts">
    import confetti from "canvas-confetti";
    import { onDestroy } from "svelte";

    interface Props {
        visible: boolean;
        duration?: number;
    }

    let { visible = false, duration = 3000 }: Props = $props();
    let canvasEl: HTMLCanvasElement | undefined = $state(undefined);
    let internalVisible = $state(false);
    let timeout: any;

    onDestroy(() => clearTimeout(timeout));

    $effect(() => {
        if (visible) {
            internalVisible = true;
            triggerEffects();
        }
    });

    function triggerEffects() {
        const myConfetti = confetti.create(canvasEl, {
            resize: true,
            useWorker: true,
        });
        myConfetti({
            particleCount: 120,
            spread: 180,
            startVelocity: 50,
            angle: -90,
            origin: { x: 0.5, y: -0.5 },
        });

        clearTimeout(timeout);
        timeout = setTimeout(() => {
            internalVisible = false;
        }, duration);
    }
</script>

{#if internalVisible}
    <!-- Конфетти -->
    <canvas bind:this={canvasEl} class="confetti-canvas"></canvas>
{/if}

<style>
    .confetti-canvas {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        pointer-events: none;
        z-index: 999;
    }
</style>
