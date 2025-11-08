export async function login(email: string, password: string): Promise<void> {
	const response = await fetch('/api/auth/login', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			email,
			password
		}),
	})

	if (!response.ok) {
		throw new Error("Authorization failed")
	}

	// User is set in server layer, so no need to set/return it here
}

export async function logout(): Promise<void> {
	const res = await fetch("/api/auth/logout", {
		method: "POST",
		credentials: "include"
	})

	if (!res.ok) {
		throw new Error("Logout failed")
	}
}
