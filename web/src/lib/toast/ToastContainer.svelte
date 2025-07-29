<script lang="ts">
  import { run } from 'svelte/legacy';

  import { toastStore } from "./toastStore";
  import Toast from "./Toast.svelte";
  // Subscribe to toast store
  let toasts = $derived($toastStore.toasts);

  function handleKeydown(event: KeyboardEvent) {
    // ESC key to dismiss all toasts
    if (event.key === "Escape" && toasts.length > 0) {
      event.preventDefault();
      toastStore.clearAll();
    }
  }

  // Focus management for keyboard navigation
  function handleFocusManagement() {
    // When toasts are present, ensure they can receive focus for keyboard navigation
    const toastElements = document.querySelectorAll("[data-toast-id]");
    if (toastElements.length > 0) {
      // Set tabindex on the first toast for keyboard accessibility
      const firstToast = toastElements[0] as HTMLElement;
      if (firstToast && !firstToast.hasAttribute("tabindex")) {
        firstToast.setAttribute("tabindex", "0");
      }
    }
  }

  run(() => {
    if (toasts.length > 0) {
      handleFocusManagement();
    }
  });
</script>

<svelte:window onkeydown={handleKeydown} />

<div class="toast toast-end">
  {#each toasts as toast (toast.id)}
    <Toast {toast} />
  {/each}
</div>
