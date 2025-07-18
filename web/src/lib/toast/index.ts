import { toastStore } from './toastStore';
import type { Toast, ToastType, ToastConfig } from './types';

// Re-export types for convenience
export type { Toast, ToastType, ToastConfig };

// Export the store for advanced usage
export { toastStore };

// Simple API functions for easy component integration
export const toast = {
  /**
   * Show a success toast notification
   * @param message - The message to display
   * @param duration - Optional duration in milliseconds (default: 3000)
   * @returns The toast ID
   */
  success: (message: string, duration?: number): string => {
    return toastStore.showSuccess(message, duration);
  },

  /**
   * Show an error toast notification
   * @param message - The message to display
   * @param duration - Optional duration in milliseconds (default: 7000)
   * @returns The toast ID
   */
  error: (message: string, duration?: number): string => {
    return toastStore.showError(message, duration);
  },

  /**
   * Show a warning toast notification
   * @param message - The message to display
   * @param duration - Optional duration in milliseconds (default: 5000)
   * @returns The toast ID
   */
  warning: (message: string, duration?: number): string => {
    return toastStore.showWarning(message, duration);
  },

  /**
   * Show an info toast notification
   * @param message - The message to display
   * @param duration - Optional duration in milliseconds (default: 4000)
   * @returns The toast ID
   */
  info: (message: string, duration?: number): string => {
    return toastStore.showInfo(message, duration);
  },

  /**
   * Add a custom toast with full configuration
   * @param config - Toast configuration object
   * @returns The toast ID
   */
  add: (config: Omit<Toast, 'id' | 'timestamp'>): string => {
    return toastStore.addToast(config);
  },

  /**
   * Remove a specific toast by ID
   * @param id - The toast ID to remove
   */
  remove: (id: string): void => {
    toastStore.removeToast(id);
  },

  /**
   * Clear all active toasts
   */
  clear: (): void => {
    toastStore.clearAll();
  }
};

// Convenience functions for direct import
export const showSuccess = toast.success;
export const showError = toast.error;
export const showWarning = toast.warning;
export const showInfo = toast.info;
export const addToast = toast.add;
export const removeToast = toast.remove;
export const clearToasts = toast.clear;

// Default export for easy usage
export default toast;