package models

type Tenant struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateTenantRequest struct {
	Name string `json:"name"`
}
