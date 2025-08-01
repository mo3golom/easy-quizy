import type { Theme, ThemeColors } from '../types';

const defaultLightColors: ThemeColors = {
	'base-100': 'oklch(100% 0 0)',
	'base-200': 'oklch(98% 0 0)',
	'base-300': 'oklch(95% 0 0)',
	'base-content': 'oklch(21% 0.006 285.885)',
	primary: 'oklch(45% 0.24 277.023)',
	'primary-content': 'oklch(100% 0 0)',
	secondary: 'oklch(65% 0.241 354.308)',
	'secondary-content': 'oklch(100% 0 0)',
	accent: 'oklch(77% 0.152 181.912)',
	'accent-content': 'oklch(38% 0.063 188.416)',
	neutral: 'oklch(100% 0 0)',
	'neutral-content': 'oklch(0% 0 0)',
	info: 'oklch(74% 0.16 232.661)',
	'info-content': 'oklch(29% 0.066 243.157)',
	success: 'oklch(76% 0.177 163.223)',
	'success-content': 'oklch(37% 0.077 168.94)',
	warning: 'oklch(82% 0.189 84.429)',
	'warning-content': 'oklch(41% 0.112 45.904)',
	error: 'oklch(71% 0.194 13.428)',
	'error-content': 'oklch(27% 0.105 12.094)'
};

export const defaultTheme: Theme = {
	name: 'default',
	colors: defaultLightColors
};

export class DefaultThemeProvider {
	async getTheme(): Promise<Theme> {
		return Promise.resolve(defaultTheme);
	}
}