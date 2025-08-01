<script lang="ts">
  import { onDestroy } from "svelte";

  interface Props {
    visible: boolean;
    duration?: number;
  }

  let { visible = false, duration = 3000 }: Props = $props();
  let internalVisible = $state(false);
  let timeout: any;

  $effect(() => {
    if (visible) {
      internalVisible = true;
      clearTimeout(timeout);
      timeout = setTimeout(() => {
        internalVisible = false;
      }, duration);
    }
  });

  onDestroy(() => clearTimeout(timeout));
</script>

{#if internalVisible}
  <div class="wrong-effect">
    <div class="vertical-lines">
      {#each Array(12) as _, i}
        <div
          class="line"
          style="left: {i * 8}vw; --final-height: {Math.random() * 80 + 10}%"
        ></div>
      {/each}
    </div>
  </div>
{/if}

<style>
  .wrong-effect {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 50vh;
    pointer-events: none;
    z-index: 999;
    background: linear-gradient(to bottom, rgba(80, 0, 160, 0.4), transparent);
  }

  .vertical-lines {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  .line {
    position: absolute;
    top: -20%;
    width: 8px;
    background-color: rgba(80, 0, 160, 0.8);
    animation: drop 1s ease-out forwards;
  }

  @keyframes drop {
    0% {
      height: 0;
      opacity: 0;
    }
    100% {
      height: var(--final-height);
      opacity: 1;
    }
  }
</style>
