import { isTMA } from "@telegram-apps/sdk";
import { hapticFeedback } from "@telegram-apps/sdk";
import type { Action } from 'svelte/action';

/**
 * Type of haptic feedback to trigger
 * - light/medium/heavy: Impact feedback for UI collisions
 * - selection: Feedback for selection changes
 * - success/warning/error: Notification feedback for task completion
 */
export type FeedbackType =
    | "light"
    | "medium"
    | "heavy"
    | "selection"
    | "success"
    | "warning"
    | "error";

/**
 * Default vibration pattern for non-Telegram environments (browser vibration API)
 * Can be a single number (duration in ms) or array of durations
 */
export let vibrationPattern: number | number[] = 50;

/**
 * Whether the feedback is disabled by default
 */
export let disabled = false;

/**
 * Check if we're running in Telegram Mini App environment
 */
function checkIsTelegram(): boolean {
    try {
        return isTMA();
    } catch {
        return false;
    }
}

/**
 * Svelte action that triggers haptic feedback on click/touch events
 * @param node - The DOM element to attach the action to
 * @param feedbackType - Type of feedback to trigger
 */
export const triggerFeedback: Action<HTMLElement, FeedbackType> = (node, feedbackType = "light") => {
    let currentFeedbackType = feedbackType;

    const handleClick = () => {
        triggerFeedbackFunction(currentFeedbackType);
    };

    node.addEventListener('click', handleClick);

    return {
        update(newFeedbackType: FeedbackType) {
            currentFeedbackType = newFeedbackType;
        },
        destroy() {
            node.removeEventListener('click', handleClick);
        }
    };
};
/**
 * Triggers haptic feedback based on the specified feedback type
 * @param type - The type of feedback to trigger
 * @param options - Optional configuration for the feedback
 */
export function triggerFeedbackFunction(
    type: FeedbackType = "light",
    options: {
        disabled?: boolean;
        vibrationPattern?: number | number[];
    } = {}
) {
    const { disabled: isDisabled = disabled, vibrationPattern: pattern = vibrationPattern } = options;

    if (isDisabled) return false;

    const isTelegram = checkIsTelegram();
    
    if (isTelegram) {
        // Telegram haptic feedback with proper availability checks
        try {
            switch (type) {
                case "light":
                    if (hapticFeedback.impactOccurred.isAvailable()) {
                        hapticFeedback.impactOccurred("light");
                        return true;
                    }
                    break;
                case "medium":
                    if (hapticFeedback.impactOccurred.isAvailable()) {
                        hapticFeedback.impactOccurred("medium");
                        return true;
                    }
                    break;
                case "heavy":
                    if (hapticFeedback.impactOccurred.isAvailable()) {
                        hapticFeedback.impactOccurred("heavy");
                        return true;
                    }
                    break;
                case "selection":
                    if (hapticFeedback.selectionChanged.isAvailable()) {
                        hapticFeedback.selectionChanged();
                        return true;
                    }
                    break;
                case "success":
                    if (hapticFeedback.notificationOccurred.isAvailable()) {
                        hapticFeedback.notificationOccurred("success");
                        return true;
                    }
                    break;
                case "warning":
                    if (hapticFeedback.notificationOccurred.isAvailable()) {
                        hapticFeedback.notificationOccurred("warning");
                        return true;
                    }
                    break;
                case "error":
                    if (hapticFeedback.notificationOccurred.isAvailable()) {
                        hapticFeedback.notificationOccurred("error");
                        return true;
                    }
                    break;
            }
        } catch (error) {
            console.warn("Failed to trigger Telegram haptic feedback:", error);
        }
    } else if (typeof navigator !== "undefined" && "vibrate" in navigator) {
        // Browser vibration API fallback
        try {
            navigator.vibrate(pattern);
            return true;
        } catch (error) {
            console.warn("Vibration not supported or failed:", error);
        }
    }

    return false;
}