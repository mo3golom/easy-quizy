<script lang="ts">
  import { toastStore } from './toastStore';
  import Toast from './Toast.svelte';
  import type { Toast as ToastType } from './types';

  // Subscribe to toast store
  $: toasts = $toastStore.toasts;
  $: maxVisible = $toastStore.maxVisible;
  
  // Get visible toasts (top 3) and hidden count
  $: visibleToasts = toasts.slice(0, maxVisible);
  $: hiddenCount = Math.max(0, toasts.length - maxVisible);

  function handleKeydown(event: KeyboardEvent) {
    // ESC key to dismiss all toasts
    if (event.key === 'Escape' && toasts.length > 0) {
      toastStore.clearAll();
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<!-- Toast Container with fixed positioning -->
{#if toasts.length > 0}
  <div 
    class="toast toast-end toast-bottom z-50 fixed bottom-4 right-4 max-w-sm w-full pointer-events-none"
    role="region"
    aria-label="Уведомления"
    aria-live="polite"
  >
    <!-- Stack container for multiple toasts -->
    <div class="stack pointer-events-auto">
      <!-- Hidden toasts indicator -->
      {#if hiddenCount > 0}
        <div 
          class="alert alert-neutral shadow-sm mb-2 text-xs opacity-75 cursor-default"
          role="status"
          aria-label="Дополнительных уведомлений: {hiddenCount}"
        >
          <div class="flex items-center justify-center w-full">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4 mr-1">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z" />
            </svg>
            <span>+{hiddenCount} уведомлений</span>
          </div>
        </div>
      {/if}

      <!-- Visible toasts -->
      {#each visibleToasts as toast (toast.id)}
        <div class="stack-item">
          <Toast {toast} />
        </div>
      {/each}
    </div>
  </div>
{/if}

<style>
  /* Ensure proper stacking and positioning */
  .toast {
    /* Override daisyUI defaults for better mobile support */
    position: fixed !important;
    pointer-events: none;
  }

  .stack {
    display: flex;
    flex-direction: column-reverse; /* Newest toasts appear at top */
    gap: 0.5rem;
    pointer-events: auto;
  }

  .stack-item {
    position: relative;
    z-index: 1;
  }

  /* Responsive adjustments */
  @media (max-width: 640px) {
    .toast {
      left: 1rem !important;
      right: 1rem !important;
      bottom: 1rem !important;
      max-width: none !important;
      width: auto !important;
    }
  }

  @media (max-width: 480px) {
    .toast {
      left: 0.5rem !important;
      right: 0.5rem !important;
      bottom: 0.5rem !important;
    }
  }

  /* Ensure toasts appear above everything */
  .toast {
    z-index: 9999 !important;
  }

  /* Stack visual effect - slight offset for stacked appearance */
  .stack-item:nth-child(n+2) {
    transform: translateY(-2px) scale(0.98);
    opacity: 0.9;
  }

  .stack-item:nth-child(n+3) {
    transform: translateY(-4px) scale(0.96);
    opacity: 0.8;
  }

  /* Smooth transitions for stack items */
  .stack-item {
    transition: transform 0.2s ease-in-out, opacity 0.2s ease-in-out;
  }

  /* Accessibility improvements */
  @media (prefers-reduced-motion: reduce) {
    .stack-item {
      transition: none;
    }
  }
</style>