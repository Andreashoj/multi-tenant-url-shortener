import type { Tenant } from '../models/tenant';
import { writable } from 'svelte/store';
import { createTenant, type CreateTenantRequest, deleteTenant, getTenants } from '../api/tenants';

interface TenantsState {
	tenants: Tenant[];
	loading: boolean;
	error: unknown | null
}

function createTenantStore() {
	const {subscribe, update} = writable<TenantsState>({
		tenants: [],
		loading: false,
		error: null,
	})

	return {
		subscribe,

		async getAll() {
			try {
				update(state =>({ ...state, loading: true }))
				const tenants = await getTenants()
				update(state =>({ ...state, tenants }))
			} catch(e) {
				update(state => ({...state, error: e}))
			} finally {
				update(state =>({ ...state, loading: false }))
			}
		},
		async create(name: string, dbType: string) {
			const payload: CreateTenantRequest = {
				name: name,
				type: dbType
			}

			try {
				update(state =>({ ...state, loading: true }))
				const tenant = await createTenant(payload)
				update(state =>({ ...state, tenants: [...state.tenants, tenant] }))
			} catch(e) {
				update(state => ({...state, error: e}))
			} finally {
				update(state =>({ ...state, loading: false }))
			}
		},
		async remove(id: number) {
			try {
				update(state => ({...state, loading: true}))
				await deleteTenant(id)
				update(state => ({
					...state,
					tenants: state.tenants.filter(item => item.id !== id)
				}))
			} catch(e) {
				update(state => ({ ...state, error: e}))
			} finally {
				update(state => ({...state, loading: false}))
			}
		}
	}
}

export const tenantStore = createTenantStore()