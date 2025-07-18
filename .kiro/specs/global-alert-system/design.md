# Design Document

## Overview

The Global Alert System provides a centralized toast notification system for the Easy Quizy web application. The system leverages daisyUI components to create consistent, accessible, and user-friendly notifications that appear across all pages of the application. The design focuses on non-intrusive user experience while ensuring important messages are clearly communicated.

## Architecture

### Component Structure

```
GlobalAlertSystem/
├── ToastContainer.svelte          # Main container component
├── Toast.svelte                   # Individual toast component
├── toastStore.ts                  # Svelte store for state management
└── types.ts                       # TypeScript interfaces
```

### State Management

The system uses a Svelte writable store to manage the global toast state. This allows any component in the application to trigger toasts without prop drilling or complex event handling.

```typescript
// Toast Store Structure
interface ToastStore {
  toasts: Toast[];
  addToast: (toast: Omit<Toast, 'id' | 'timestamp'>) => void;
  removeToast: (id: string) => void;
  clearAll: () => void;
}
```

### Integration Points

1. **Layout Integration**: The ToastContainer will be added to the root layout (`+layout.svelte`) to ensure global availability
2. **API Integration**: Error handling in `client.ts` will be enhanced to trigger toast notifications
3. **Component Migration**: Existing error displays in `ApiQuizQuestion.svelte` and `ApiQuizResult.svelte` will be replaced with toast calls

## Components and Interfaces

### Toast Interface

```typescript
interface Toast {
  id: string;
  type: 'success' | 'error' | 'warning' | 'info';
  message: string;
  duration?: number; // milliseconds, default 5000
  dismissible?: boolean; // default true
  timestamp: number;
}
```

### ToastContainer Component

**Purpose**: Manages the display and positioning of all toast notifications

**Key Features**:
- Fixed positioning at bottom-right corner
- Z-index management to appear above all content
- Stack management (max 3 visible toasts)
- Responsive design for mobile and desktop
- Auto-dismiss timers
- Manual dismiss handlers

**daisyUI Components Used**:
- `toast` class for positioning
- `stack` component for managing multiple toasts
- `alert` component for individual toast styling

### Toast Component

**Purpose**: Individual toast notification with type-specific styling

**Key Features**:
- Type-based styling (success, error, warning, info)
- Close button with accessibility support
- Icon integration for visual distinction
- Smooth enter/exit animations
- Progress indicator for auto-dismiss countdown

**daisyUI Components Used**:
- `alert` with semantic color variants
- `btn-ghost` for close button
- Icons from Heroicons (already used in the app)

### Toast Store

**Purpose**: Centralized state management for all toast notifications

**Key Methods**:
```typescript
// Add a new toast
addToast(toast: Omit<Toast, 'id' | 'timestamp'>): void

// Remove specific toast
removeToast(id: string): void

// Clear all toasts
clearAll(): void

// Convenience methods
showSuccess(message: string, duration?: number): void
showError(message: string, duration?: number): void
showWarning(message: string, duration?: number): void
showInfo(message: string, duration?: number): void
```

## Data Models

### Toast State Model

```typescript
interface ToastState {
  toasts: Toast[];
  maxVisible: number; // 3
  defaultDuration: number; // 5000ms
}
```

### Toast Configuration

```typescript
interface ToastConfig {
  type: ToastType;
  message: string;
  duration?: number;
  dismissible?: boolean;
  persistent?: boolean; // for critical errors
}
```

### Toast Type Definitions

```typescript
type ToastType = 'success' | 'error' | 'warning' | 'info';

interface ToastTypeConfig {
  alertClass: string;
  icon: string;
  defaultDuration: number;
}

const TOAST_CONFIGS: Record<ToastType, ToastTypeConfig> = {
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
```

## Error Handling

### API Error Integration

The existing `apiRequest` function in `client.ts` will be enhanced to automatically trigger error toasts:

```typescript
// Enhanced error handling
async function apiRequest(endpoint: string, options: RequestInit = {}): Promise<Response> {
  try {
    // ... existing logic
    
    if (!response.ok) {
      // Trigger error toast instead of throwing
      toastStore.showError(`Ошибка API: ${response.status} ${response.statusText}`);
      throw new Error(`API request failed: ${response.status} ${response.statusText}`);
    }
    
    return response;
  } catch (error) {
    // Handle network errors
    toastStore.showError('Ошибка сети. Проверьте подключение к интернету.');
    throw error;
  }
}
```

### Component Error Migration

Current inline error displays will be replaced with toast notifications:

1. **ApiQuizQuestion**: Remove inline error alerts, use toast for API errors
2. **ApiQuizResult**: Replace error states with toast notifications
3. **Form Validation**: Any future form validation will use toast system

### Error Categories

- **Network Errors**: Connection issues, timeouts
- **API Errors**: Server responses with error status codes
- **Validation Errors**: Client-side form validation failures
- **Application Errors**: Unexpected application state errors

## Testing Strategy

### Unit Testing

1. **Toast Store Tests**:
   - Test toast addition and removal
   - Verify auto-dismiss timers
   - Test stack management (max 3 visible)
   - Validate state updates

2. **Component Tests**:
   - Toast rendering with different types
   - Close button functionality
   - Animation behavior
   - Accessibility attributes

3. **Integration Tests**:
   - API error handling triggers toasts
   - Toast positioning and stacking
   - Cross-component toast triggering

### Accessibility Testing

1. **Screen Reader Compatibility**:
   - ARIA live regions for toast announcements
   - Proper role and label attributes
   - Keyboard navigation support

2. **Visual Accessibility**:
   - Color contrast compliance (WCAG 2.1 AA)
   - Icon + text combination for meaning
   - Focus management for dismissible toasts

3. **Keyboard Accessibility**:
   - ESC key to dismiss toasts
   - Tab navigation to close buttons
   - Focus trap management

### Manual Testing Scenarios

1. **Toast Display**:
   - Multiple toast types simultaneously
   - Stack behavior with >3 toasts
   - Auto-dismiss timing accuracy
   - Manual dismiss functionality

2. **Responsive Behavior**:
   - Mobile device positioning
   - Tablet layout adaptation
   - Desktop multi-monitor scenarios

3. **Integration Testing**:
   - API error scenarios
   - Component error migration
   - Cross-page navigation with active toasts

## Implementation Considerations

### Performance

- **Lazy Loading**: Toast components only render when needed
- **Memory Management**: Automatic cleanup of dismissed toasts
- **Animation Optimization**: CSS transitions over JavaScript animations
- **Store Efficiency**: Minimal re-renders through selective subscriptions

### Browser Compatibility

- **Modern Browser Support**: ES2020+ features (crypto.randomUUID already used)
- **Fallback Handling**: Graceful degradation for older browsers
- **CSS Grid/Flexbox**: Already used in existing components

### Internationalization

- **Message Localization**: Support for Russian text (current app language)
- **RTL Support**: Future consideration for right-to-left languages
- **Cultural Considerations**: Toast positioning and behavior preferences

### Security

- **XSS Prevention**: Proper text escaping in toast messages
- **Content Sanitization**: Safe handling of dynamic toast content
- **Rate Limiting**: Prevent toast spam from malicious code

## Migration Strategy

### Phase 1: Core Implementation
1. Create toast store and types
2. Implement Toast and ToastContainer components
3. Add ToastContainer to root layout
4. Create utility functions for easy toast triggering

### Phase 2: API Integration
1. Enhance apiRequest function with toast error handling
2. Update existing API error handling patterns
3. Test API error scenarios

### Phase 3: Component Migration
1. Remove inline error displays from ApiQuizQuestion
2. Remove inline error displays from ApiQuizResult
3. Replace with appropriate toast calls
4. Update any other components with error displays

### Phase 4: Enhancement
1. Add success toasts for positive actions
2. Implement warning toasts for user guidance
3. Add info toasts for system notifications
4. Fine-tune timing and user experience

This design ensures a robust, accessible, and user-friendly toast notification system that integrates seamlessly with the existing Easy Quizy application architecture while providing a foundation for future enhancements.