package models

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleOwner  Role = "owner"
)

type User struct {
	ID       uint   `json:"id"`
	TenantID uint   `json:"tenant_id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     Role   `json:"role"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
