<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { toastStore } from './toastStore';
  import { TOAST_CONFIGS, type Toast } from './types';

  interface Props {
    toast: Toast;
  }

  let { toast }: Props = $props();

  let progressElement: HTMLDivElement | undefined = $state();
  let toastElement: HTMLDivElement | undefined = $state();
  let animationFrame: number;
  let startTime: number;
  let isVisible = $state(false);

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
    // Only handle keydown if this toast element has focus or is the target
    if (event.target === toastElement || toastElement?.contains(event.target as Node)) {
      if (event.key === 'Escape') {
        event.preventDefault();
        handleClose();
      } else if (event.key === 'Enter' || event.key === ' ') {
        // Allow Enter or Space to close dismissible toasts when focused
        if (toast.dismissible !== false) {
          event.preventDefault();
          handleClose();
        }
      }
    }
  }

  // Enhanced focus management
  function handleFocus() {
    // When toast receives focus, ensure it's visible and announce it
    if (toastElement) {
      toastElement.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
    }
  }
</script>



<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<div
  bind:this={toastElement}
  class="{config.alertClass} alert  mb-2 transition-all duration-300 ease-in-out transform relative overflow-hidden"
  class:opacity-0={!isVisible}
  class:opacity-100={isVisible}
  role="alert"
  data-toast-id={toast.id}
  data-toast-type={toast.type}
  tabindex={toast.dismissible !== false ? 0 : -1}
  onfocus={handleFocus}
  onkeydown={handleKeydown}
>
  <!-- Progress bar for auto-dismiss countdown -->
  {#if duration > 0}
    <div class="absolute bottom-0 left-0 h-1 bg-current opacity-30 transition-all duration-100 ease-linear" bind:this={progressElement}></div>
  {/if}

  <!-- Message -->
  <div class="flex-1 min-w-0">
    <span 
      id="toast-message-{toast.id}"
      class="text-sm font-medium break-words"
    >
      {toast.message}
    </span>
  </div>

  <!-- Close button -->
  {#if toast.dismissible !== false}
    <button
      type="button"
      class="btn btn-ghost btn-sm btn-circle ml-2 flex-shrink-0"
      onclick={handleClose}
      aria-label="Закрыть уведомление"
      title="Закрыть уведомление"
    >
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>
    </button>
  {/if}
</div>