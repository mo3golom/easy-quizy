export interface ThemeColors {
	primary: string;
	'primary-content': string;
	secondary: string;
	'secondary-content': string;
	accent: string;
	'accent-content': string;
	'base-100': string;
	'base-200': string;
	'base-300': string;
	'base-content': string;
	info: string;
	'info-content': string;
	success: string;
	'success-content': string;
	warning: string;
	'warning-content': string;
	error: string;
	'error-content': string;
	neutral: string;
	'neutral-content': string;
}

export interface Theme {
	name: string;
	colors: Partial<ThemeColors>;
}

export type Environment = 'browser' | 'telegram';