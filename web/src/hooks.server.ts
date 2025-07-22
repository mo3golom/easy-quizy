import type { Handle } from '@sveltejs/kit';
import { PUBLIC_API_BASE_URL } from '$env/static/public';

export const handle: Handle = async ({ event, resolve }) => {
	// In production, proxy API requests to the Go backend
	if (event.url.pathname.startsWith('/api')) {
		// In Docker, both services run in the same container, so use localhost
		// In development, this can be overridden with BACKEND_URL env var
		const backendUrl = PUBLIC_API_BASE_URL || 'http://localhost:8080';
		const apiUrl = `${backendUrl}${event.url.pathname}${event.url.search}`;
		
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