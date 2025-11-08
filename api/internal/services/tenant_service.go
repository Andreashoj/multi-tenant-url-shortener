package services

import (
	"api/internal/models"
	"api/internal/repos"
	"errors"
	"fmt"
)

type TenantService interface {
	Create(userID uint, name string) (*models.Tenant, error)
	Get(userID uint, tenantID uint) (*models.Tenant, error)
	GetAll(userID uint) ([]*models.Tenant, error)
	Update(userID uint, tenantID uint, name string) (*models.Tenant, error)
	Delete(userID uint, tenantID uint) error
}

type tenantService struct {
	tenantRepo  repos.TenantRepo
	userService UserService
}

func NewTenantService(tenantRepo repos.TenantRepo, userService UserService) TenantService {
	return &tenantService{
		tenantRepo:  tenantRepo,
		userService: userService,
	}
}

func (t tenantService) GetAll(userID uint) ([]*models.Tenant, error) {
	user, err := t.userService.Me(userID)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving the user: %w", err)
	}

	if user.Role != models.RoleAdmin {
		return nil, errors.New("the users role doesn't have permission to create tenant")
	}

	tenants, err := t.tenantRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tenants: %w", err)
	}

	return tenants, nil
}

func (t tenantService) Create(userID uint, name string) (*models.Tenant, error) {
	user, err := t.userService.Me(userID)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving the user: %w", err)
	}

	if user.Role != models.RoleAdmin {
		return nil, errors.New("the users role doesn't have permission to create tenant")
	}

	tenant, err := t.tenantRepo.Create(name) // TODO: Store the user who created the tenant ?
	if err != nil {
		return nil, fmt.Errorf("failed creating the tenant: %w", err)
	}

	return tenant, nil
}

func (t tenantService) Get(userID uint, tenantID uint) (*models.Tenant, error) {
	//TODO implement me
	panic("implement me")
}

func (t tenantService) Update(userID uint, tenantID uint, name string) (*models.Tenant, error) {
	//TODO implement me
	panic("implement me")
}

func (t tenantService) Delete(userID uint, tenantID uint) error {
	//TODO implement me
	panic("implement me")
}
