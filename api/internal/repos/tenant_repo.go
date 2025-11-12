package repos

import (
	"api/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type TenantRepo interface {
	Create(name string, tp string) (*models.Tenant, error)
	Get(id uint) (*models.Tenant, error)
	GetAll() ([]*models.Tenant, error)
	Delete(id uint) error
	Update(id uint, name string) (*models.Tenant, error)
}

type tenantRepo struct {
	db *sql.DB
}

func NewTenantRepo(db *sql.DB) TenantRepo {
	return &tenantRepo{db: db}
}

var ErrNoTenantFound = errors.New("no tenant found")

func (t tenantRepo) Create(name string, tp string) (*models.Tenant, error) {
	tenant, err := models.NewTenant(name, tp)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant: %w", err)
	}

	err = t.db.QueryRow(`INSERT INTO tenants (name, type) VALUES ($1, $2) RETURNING id`, name, tp).Scan(&tenant.Id)
	if err != nil {
		return nil, fmt.Errorf("inserting the teanant went wrong: %w", err)
	}

	return tenant, nil
}

func (t tenantRepo) GetAll() ([]*models.Tenant, error) {
	rows, err := t.db.Query(`SELECT id, name FROM tenants`)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving tenants: %w", err)
	}
	defer rows.Close()

	var tenants []*models.Tenant
	for rows.Next() {
		var tenant models.Tenant
		err = rows.Scan(&tenant.Id, &tenant.Name)
		if err != nil {
			return nil, fmt.Errorf("failed mapping tenant row: %w", err)
		}
		tenants = append(tenants, &tenant)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed iteraing through the rows: %w", err)
	}

	return tenants, nil
}

func (t tenantRepo) Get(id uint) (*models.Tenant, error) {
	//TODO implement me
	panic("implement me")
}

func (t tenantRepo) Delete(id uint) error {
	result, err := t.db.Exec(`DELETE FROM tenants WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed deleting tenant: %d, with error: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed getting rows effected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNoTenantFound
	}

	return nil
}

func (t tenantRepo) Update(id uint, name string) (*models.Tenant, error) {
	//TODO implement me
	panic("implement me")
}
