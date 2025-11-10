package models

import (
	"fmt"
)

type DBType string

const (
	DBTypeIsolated DBType = "isolated"
	DBTypeSchema   DBType = "schema"
	DBTypeShared   DBType = "shared"
)

type Tenant struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Type DBType `json:"type"`
}

type CreateTenantRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func IsValidTypeDB(t string) bool {
	switch DBType(t) {
	case DBTypeIsolated, DBTypeSchema, DBTypeShared:
		return true
	default:
		return false
	}
}

func NewTenant(name string, t string) (*Tenant, error) {
	if name == "" {
		return nil, fmt.Errorf("name must not be emptyu: %s", t)
	}
	if !IsValidTypeDB(t) {
		return nil, fmt.Errorf("invalid db type passed: %s", t)
	}

	return &Tenant{Name: name, Type: DBType(t)}, nil
}
