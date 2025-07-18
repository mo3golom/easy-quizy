# Implementation Plan

- [x] 1. Create core toast system types and interfaces
  - Create `web/src/lib/toast/types.ts` with Toast interface, ToastType, and ToastConfig definitions
  - Define TypeScript interfaces for toast state management and configuration
  - _Requirements: 5.5_

- [x] 2. Implement toast store for state management
  - Create `web/src/lib/toast/toastStore.ts` with Svelte writable store
  - Implement addToast, removeToast, clearAll methods
  - Add convenience methods: showSuccess, showError, showWarning, showInfo
  - Include auto-dismiss timer logic and stack management (max 3 visible)
  - _Requirements: 5.1, 5.2, 4.1, 4.2, 4.3_

- [x] 3. Create individual Toast component
  - Create `web/src/lib/toast/Toast.svelte` component
  - Implement type-based styling using daisyUI alert classes (alert-success, alert-error, alert-warning, alert-info)
  - Add close button with accessibility support and proper ARIA attributes
  - Include appropriate icons for each toast type using existing Heroicons
  - Implement smooth enter/exit animations and auto-dismiss countdown
  - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 7.1, 7.2, 7.4_

- [x] 4. Create ToastContainer component for global positioning
  - Create `web/src/lib/toast/ToastContainer.svelte` component
  - Implement fixed positioning at bottom-right corner with proper z-index
  - Use daisyUI toast class for positioning and stack component for multiple toasts
  - Handle responsive design for mobile and desktop viewports
  - Implement stack management showing only top 3 toasts with indication of additional stacked toasts
  - _Requirements: 3.1, 3.2, 3.3, 3.4, 4.1, 4.2, 4.3, 4.4, 4.5_

- [x] 5. Integrate ToastContainer into application layout
  - Modify `web/src/routes/+layout.svelte` to include ToastContainer component
  - Ensure ToastContainer is available globally across all pages
  - Test that toasts appear correctly on different routes
  - _Requirements: 5.3, 5.4_

- [x] 6. Create toast utility functions and API
  - Create `web/src/lib/toast/index.ts` as main export file
  - Export toast store methods and utility functions for easy component integration
  - Create simple API functions that components can import and use
  - Ensure TypeScript type safety for all exported functions
  - _Requirements: 5.1, 5.2, 5.5_

- [x] 7. Enhance API client with toast error handling
  - Modify `web/src/lib/api/client.ts` apiRequest function to trigger error toasts
  - Replace thrown errors with toast notifications for network and API errors
  - Handle different error types (network errors, API status errors) with appropriate toast messages
  - Maintain existing error throwing behavior for component error handling
  - _Requirements: 6.3_

- [x] 8. Migrate main page error handling to toast system
  - Remove inline error display elements from `web/src/routes/+page.svelte` (both daily quiz and game ID errors)
  - Replace error handling with toast notifications using the new toast system
  - Import and use toast utility functions for error scenarios
  - Test that main page errors now appear as toasts instead of inline messages
  - _Requirements: 6.1, 6.2_

- [x] 9. Migrate ApiQuizQuestion component error handling
  - Remove inline explanation alert from `web/src/lib/components/ApiQuizQuestion.svelte` (currently shows as alert-dash)
  - Replace any error handling with toast notifications using the new toast system
  - Import and use toast utility functions for error scenarios
  - Note: Component currently shows explanations inline, evaluate if this should remain or move to toast
  - _Requirements: 6.1_