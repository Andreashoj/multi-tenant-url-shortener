import type { RequestHandler } from '@sveltejs/kit';

const API_URL = "http://host.docker.internal:8080";

/* PROXY */
/* Handle all client requests through client server, easier handle cross-origin cookies */
async function proxyRequest(event: any, method: string) {
	const { path } = event.params;
	const cookies = event.request.headers.get('cookie');
	const url = `${API_URL}/api/${path}`;
	// Create new request and attach cookies
	const options: RequestInit = {
		method,
		headers: {
			'content-type': 'application/json',
			cookie: cookies || ''
		}
	};

	// Attach body to request for anything but GET
	if (method !== 'GET') {
		options.body = await event.request.text();
	}

	try {
		// Make request with proxy URL
		const response = await fetch(url, options);
		const data = await response.text();

		// Create headers and cookies for client, based off response
		const headers = new Headers();
		headers.set('content-type', response.headers.get('content-type') || 'application/json');
		const setCookie = response.headers.get('set-cookie');
		if (setCookie) headers.set('set-cookie', setCookie);

		// Return response with headers and cookies attached
		return new Response(data, {
			status: response.status,
			headers
		});
	} catch (error) {
		console.error('Proxy error:', error);
		return new Response(
			JSON.stringify({ error: 'Failed to connect to API' }),
			{
				status: 500,
				headers: { 'content-type': 'application/json' }
			}
		);
	}
}

// Export handlers for each HTTP method
export const GET: RequestHandler = (event) => proxyRequest(event, 'GET');
export const POST: RequestHandler = (event) => proxyRequest(event, 'POST');
export const PUT: RequestHandler = (event) => proxyRequest(event, 'PUT');
export const DELETE: RequestHandler = (event) => proxyRequest(event, 'DELETE');
export const PATCH: RequestHandler = (event) => proxyRequest(event, 'PATCH');