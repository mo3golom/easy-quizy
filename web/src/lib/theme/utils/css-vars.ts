import type { ThemeColors } from '../types';

export function applyTheme(colors: Partial<ThemeColors>) {
	if (typeof document === 'undefined') {
		return;
	}
	const root = document.documentElement;

	for (const [key, value] of Object.entries(colors)) {
		if (value) {
			root.style.setProperty(`--theme-${key}`, value);
		} else {
			root.style.removeProperty(`--theme-${key}`);
		}
	}
}