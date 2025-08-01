import { readable, writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { Environment, Theme } from './types';
import { applyTheme } from './utils/css-vars';
import { DefaultThemeProvider, defaultTheme } from './providers/default';
import { TelegramThemeProvider } from './providers/telegram';
import { isTMA } from '@telegram-apps/sdk';

export const environment = readable<Environment>(detectEnvironment());
export const loading = writable<boolean>(true);
export const currentTheme = writable<Theme>(defaultTheme);

function detectEnvironment(): Environment {
    if (isTMA()) {
		return 'telegram';
	}
	
	return 'browser';
}

async function initializeTheme() {
	loading.set(true);
	const env = detectEnvironment();

	try {
		if (env === 'telegram') {
			const provider = new TelegramThemeProvider();

			const applyTelegramTheme = async () => {
				const theme = await provider.getTheme();
				const finalColors = { ...defaultTheme.colors, ...theme.colors };
				const finalTheme: Theme = { name: 'telegram', colors: finalColors };

				currentTheme.set(finalTheme);
				applyTheme(finalTheme.colors);
			};

			await applyTelegramTheme();

		} else {
			const provider = new DefaultThemeProvider();
			const theme = await provider.getTheme();
			currentTheme.set(theme);
			applyTheme(theme.colors);
		}
	} catch (error) {
		console.error('Failed to initialize theme, falling back to default.', error);
		const theme = await new DefaultThemeProvider().getTheme();
		currentTheme.set(theme);
		applyTheme(theme.colors);
	} finally {
		loading.set(false);
	}
}

export const theme = {
	subscribe: currentTheme.subscribe,
	initialize: initializeTheme
};