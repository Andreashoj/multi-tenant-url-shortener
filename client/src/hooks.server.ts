export async function handle({ event, resolve }) {
	const PROTECTED_ROUTES = ["/dashboard"]
	const API_URL = "http://host.docker.internal:8080"
	const cookies = event.request.headers.get('cookie');

	console.log(cookies)

	try {
		const res = await fetch(`${API_URL}/api/auth/me`, {
			method: 'GET',
			headers: {
				cookie: cookies || ''
			}
		})

		console.log("here..", res)

		if (res.ok) {
			const user = await res.json()
			console.log('Response status:', user);
			event.locals.user = user
		} else {

			if (PROTECTED_ROUTES.includes(event.url.pathname || "")) {
				return new Response(null, {
					status: 302,
					headers: {
						location: '/login'
					}
				});
			}
		}
	} catch(e) {
		console.error('Auth check failed:', e)
	}

	return await resolve(event);
}