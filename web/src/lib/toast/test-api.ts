// Test file to verify toast API functionality
import toast, { showSuccess, showError, showWarning, showInfo, addToast, removeToast, clearToasts } from './index';
import type { ToastType, ToastConfig } from './index';

// Test the main toast object API
function testToastObjectAPI() {
  // Test convenience methods
  const successId = toast.success('Success message');
  const errorId = toast.error('Error message', 10000);
  const warningId = toast.warning('Warning message');
  const infoId = toast.info('Info message');

  // Test custom toast
  const customId = toast.add({
    type: 'success',
    message: 'Custom toast',
    duration: 2000,
    dismissible: true
  });

  // Test removal
  toast.remove(successId);
  
  // Test clear all
  toast.clear();

  console.log('Toast object API test passed');
}

// Test direct function imports
function testDirectFunctionAPI() {
  const successId = showSuccess('Direct success');
  const errorId = showError('Direct error');
  const warningId = showWarning('Direct warning');
  const infoId = showInfo('Direct info');

  const customId = addToast({
    type: 'error',
    message: 'Direct custom toast',
    duration: 5000
  });

  removeToast(successId);
  clearToasts();

  console.log('Direct function API test passed');
}

// Test TypeScript type safety
function testTypeSafety() {
  // These should compile without errors
  const validType: ToastType = 'success';
  
  const validConfig: ToastConfig = {
    type: 'error',
    message: 'Test message',
    duration: 3000,
    dismissible: true,
    persistent: false
  };

  // Test that functions return correct types
  const toastId: string = toast.success('Test');
  const voidResult: void = toast.remove(toastId);
  const voidClear: void = toast.clear();

  console.log('Type safety test passed');
}

// Export test functions for potential use
export { testToastObjectAPI, testDirectFunctionAPI, testTypeSafety };