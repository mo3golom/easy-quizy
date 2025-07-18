<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { toastStore } from './toastStore';
  import { TOAST_CONFIGS, type Toast } from './types';

  export let toast: Toast;

  let progressElement: HTMLDivElement;
  let toastElement: HTMLDivElement;
  let animationFrame: number;
  let startTime: number;
  let isVisible = false;

  const config = TOAST_CONFIGS[toast.type];
  const duration = toast.duration || config.defaultDuration;

  // Animation and auto-dismiss logic
  onMount(() => {
    // Trigger enter animation
    requestAnimationFrame(() => {
      isVisible = true;
    });

    // Start progress bar animation if toast has duration
    if (duration > 0) {
      startTime = Date.now();
      updateProgress();
    }
  });

  onDestroy(() => {
    if (animationFrame) {
      cancelAnimationFrame(animationFrame);
    }
  });

  function updateProgress() {
    if (!progressElement || duration <= 0) return;

    const elapsed = Date.now() - startTime;
    const progress = Math.min(elapsed / duration, 1);
    
    progressElement.style.width = `${(1 - progress) * 100}%`;

    if (progress < 1) {
      animationFrame = requestAnimationFrame(updateProgress);
    }
  }

  function handleClose() {
    if (toast.dismissible !== false) {
      isVisible = false;
      // Wait for exit animation before removing from store
      setTimeout(() => {
        toastStore.removeToast(toast.id);
      }, 300);
    }
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Escape') {
      handleClose();
    }
  }

  // Icon SVGs for each toast type
  function getIconSvg(iconName: string) {
    switch (iconName) {
      case 'check-circle':
        return `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>`;
      case 'x-circle':
        return `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="m9.75 9.75 4.5 4.5m0-4.5-4.5 4.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>`;
      case 'exclamation-triangle':
        return `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" />
        </svg>`;
      case 'information-circle':
        return `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>`;
      default:
        return '';
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<div
  bind:this={toastElement}
  class="toast-item {config.alertClass} alert shadow-lg mb-2 transition-all duration-300 ease-in-out transform relative overflow-hidden"
  class:translate-x-full={!isVisible}
  class:opacity-0={!isVisible}
  class:translate-x-0={isVisible}
  class:opacity-100={isVisible}
  role="alert"
  aria-live="polite"
  aria-atomic="true"
  aria-label="{toast.type} notification: {toast.message}"
>
  <!-- Progress bar for auto-dismiss countdown -->
  {#if duration > 0}
    <div class="absolute bottom-0 left-0 h-1 bg-current opacity-30 transition-all duration-100 ease-linear" bind:this={progressElement}></div>
  {/if}

  <!-- Icon -->
  <div class="flex-shrink-0" aria-hidden="true">
    {@html getIconSvg(config.icon)}
  </div>

  <!-- Message -->
  <div class="flex-1 min-w-0">
    <span class="text-sm font-medium break-words">{toast.message}</span>
  </div>

  <!-- Close button -->
  {#if toast.dismissible !== false}
    <button
      type="button"
      class="btn btn-ghost btn-sm btn-circle ml-2 flex-shrink-0"
      on:click={handleClose}
      aria-label="Закрыть уведомление"
      title="Закрыть уведомление"
    >
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>
    </button>
  {/if}
</div>

<style>
  .toast-item {
    max-width: 400px;
    min-width: 300px;
  }

  @media (max-width: 640px) {
    .toast-item {
      max-width: calc(100vw - 2rem);
      min-width: calc(100vw - 2rem);
    }
  }

  /* Ensure smooth animations */
  .toast-item {
    will-change: transform, opacity;
  }
</style>