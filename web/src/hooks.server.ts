import type { Handle } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public';

// Simple logger utility
const logger = {
	info: (message: string, data?: any) => {
		console.log(`[PROXY] ${new Date().toISOString()} - ${message}`, data ? JSON.stringify(data, null, 2) : '');
	},
	error: (message: string, error?: any) => {
		console.error(`[PROXY ERROR] ${new Date().toISOString()} - ${message}`, error);
	}
};

export const handle: Handle = async ({ event, resolve }) => {
	// In production, proxy API requests to the Go backend
	if (event.url.pathname.startsWith('/api')) {
		// In Docker, both services run in the same container, so use localhost
		// In development, this can be overridden with BACKEND_URL env var
		const backendUrl = PUBLIC_API_BASE_URL || 'http://localhost:8080';
		const apiUrl = `${backendUrl}${event.url.pathname}${event.url.search}`;
		
		// Log incoming request
		logger.info(`Proxying ${event.request.method} request`, {
			originalUrl: event.url.pathname + event.url.search,
			targetUrl: apiUrl,
			userAgent: event.request.headers.get('user-agent'),
			contentType: event.request.headers.get('content-type')
		});
		
		try {
			const requestInit: RequestInit = {
				method: event.request.method,
				headers: {
					...Object.fromEntries(event.request.headers),
					'host': new URL(backendUrl).host
				}
			};

			// Only add body for non-GET/HEAD requests
			if (event.request.method !== 'GET' && event.request.method !== 'HEAD') {
				requestInit.body = await event.request.arrayBuffer();
			}

			const response = await fetch(apiUrl, requestInit);

			// Log successful response
			logger.info(`Proxy response received`, {
				status: response.status,
				statusText: response.statusText,
				contentType: response.headers.get('content-type'),
				contentLength: response.headers.get('content-length')
			});

			// Create response with proper headers for CORS and content type
			const responseHeaders = new Headers(response.headers);
			
			return new Response(response.body, {
				status: response.status,
				statusText: response.statusText,
				headers: responseHeaders
			});
		} catch (error) {
			console.error('API proxy error:', error);
			return new Response(JSON.stringify({ error: 'Internal Server Error' }), { 
				status: 500,
				headers: {
					'Content-Type': 'application/json'
				}
			});
		}
	}

	return resolve(event);
};