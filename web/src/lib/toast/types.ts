export type ToastType = 'success' | 'error' | 'warning' | 'info';

export interface Toast {
  id: string;
  type: ToastType;
  message: string;
  duration?: number; // milliseconds, default varies by type
  dismissible?: boolean; // default true
  timestamp: number;
}

export interface ToastConfig {
  type: ToastType;
  message: string;
  duration?: number;
  dismissible?: boolean;
  persistent?: boolean; // for critical errors
}

export interface ToastTypeConfig {
  alertClass: string;
  icon: string;
  defaultDuration: number;
}

export const TOAST_CONFIGS: Record<ToastType, ToastTypeConfig> = {
  success: {
    alertClass: 'alert-success',
    icon: 'check-circle',
    defaultDuration: 3000
  },
  error: {
    alertClass: 'alert-error', 
    icon: 'x-circle',
    defaultDuration: 7000
  },
  warning: {
    alertClass: 'alert-warning',
    icon: 'exclamation-triangle', 
    defaultDuration: 5000
  },
  info: {
    alertClass: 'alert-info',
    icon: 'information-circle',
    defaultDuration: 4000
  }
};