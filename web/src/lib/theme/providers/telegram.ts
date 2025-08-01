import { themeParams } from '@telegram-apps/sdk';
import type { Theme, ThemeColors } from '../types';

const TELEGRAM_COLOR_MAPPING: { [key in keyof ThemeColors]?: () => string | undefined } = {
	'base-100': themeParams.backgroundColor,
	'base-content': themeParams.textColor,
	primary: themeParams.headerBackgroundColor,
	'primary-content': themeParams.buttonTextColor,
	'base-200': themeParams.secondaryBackgroundColor,
	'base-300': themeParams.headerBackgroundColor,
	accent: themeParams.accentTextColor,
	info: themeParams.linkColor,
	error: themeParams.destructiveTextColor,
	neutral: themeParams.buttonTextColor,
	'neutral-content': themeParams.subtitleTextColor
};

export class TelegramThemeProvider {
	async getTheme(): Promise<Theme> {
		if (!themeParams.isMounted()) {
			themeParams.mountSync();
		}

		const colors: Partial<ThemeColors> = {};
		for (const [colorKey, getColorValue] of Object.entries(TELEGRAM_COLOR_MAPPING)) {
			const value = getColorValue?.();
			if (value) {
				colors[colorKey as keyof ThemeColors] = value;
			}
		}

		return {
			name: 'telegram',
			colors
		};
	}
}