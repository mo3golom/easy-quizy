import { writable } from 'svelte/store';
import type { Toast, ToastType } from './types';

interface ToastState {
  toasts: Toast[];
  maxVisible: number;
  defaultDuration: number;
}

const initialState: ToastState = {
  toasts: [],
  maxVisible: 3,
  defaultDuration: 5000
};

function createToastStore() {
  const { subscribe, update } = writable<ToastState>(initialState);

  // Auto-dismiss timers map
  const timers = new Map<string, NodeJS.Timeout>();

  function generateId(): string {
    return crypto.randomUUID();
  }

  function addToast(toast: Omit<Toast, 'id' | 'timestamp'>) {
    const id = generateId();
    const timestamp = Date.now();
    const duration = toast.duration ?? initialState.defaultDuration;
    
    const newToast: Toast = {
      ...toast,
      id,
      timestamp,
      duration,
      dismissible: toast.dismissible ?? true
    };

    update(state => ({
      ...state,
      toasts: [newToast, ...state.toasts]
    }));

    // Set up auto-dismiss timer if duration > 0
    if (duration > 0) {
      const timer = setTimeout(() => {
        removeToast(id);
      }, duration);
      timers.set(id, timer);
    }

    return id;
  }

  function removeToast(id: string) {
    // Clear timer if exists
    const timer = timers.get(id);
    if (timer) {
      clearTimeout(timer);
      timers.delete(id);
    }

    update(state => ({
      ...state,
      toasts: state.toasts.filter(toast => toast.id !== id)
    }));
  }

  function clearAll() {
    // Clear all timers
    timers.forEach(timer => clearTimeout(timer));
    timers.clear();

    update(state => ({
      ...state,
      toasts: []
    }));
  }

  // Convenience methods for different toast types
  function showSuccess(message: string, duration?: number) {
    return addToast({
      type: 'success',
      message,
      duration: duration ?? 3000
    });
  }

  function showError(message: string, duration?: number) {
    return addToast({
      type: 'error',
      message,
      duration: duration ?? 7000
    });
  }

  function showWarning(message: string, duration?: number) {
    return addToast({
      type: 'warning',
      message,
      duration: duration ?? 5000
    });
  }

  function showInfo(message: string, duration?: number) {
    return addToast({
      type: 'info',
      message,
      duration: duration ?? 4000
    });
  }

  return {
    subscribe,
    addToast,
    removeToast,
    clearAll,
    showSuccess,
    showError,
    showWarning,
    showInfo
  };
}

export const toastStore = createToastStore();