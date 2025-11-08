import type { Tenant } from '../models/tenant';

export interface CreateTenantRequest {
	name: string
}
export async function createTenant(payload: CreateTenantRequest): Promise<Tenant> {
	const res = await fetch("/api/tenant", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(payload),
		credentials: "include"
	})

	if (!res.ok) {
		throw new Error("Failed creating tenant")
	}

	return res.json()
}

export async function getTenants(): Promise<Tenant[]> {
	const res = await fetch("/api/tenant", {
		method: "GET",
		credentials: "include"
	})

	if (!res.ok) {
		throw new Error("Failed getting tenants")
	}

	return res.json()
}