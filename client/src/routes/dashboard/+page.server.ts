import { type Actions, redirect } from '@sveltejs/kit';

export const actions = {
	logout: async ({ cookies, fetch }) => {
		const accessToken = cookies.get('access_token');
		const refreshToken = cookies.get('refresh_token');

		const response = await fetch('http://host.docker.internal:8080/api/auth/logout', {
			method: 'POST',
			headers: {
				cookie: `access_token=${accessToken}; refresh_token=${refreshToken}`
			}
		});

		cookies.delete('access_token', { path: '/' });
		cookies.delete('refresh_token', { path: '/' });

		throw redirect(303, '/login');
	}
} satisfies Actions;