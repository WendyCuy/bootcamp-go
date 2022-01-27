package product

import (
	"database/sql"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/models"
)

// Repository encapsulates the storage of a product.
type Repository interface {
	GetByName(name string) (models.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(name string) (models.Product, error) {
	query := "SELECT * FROM products WHERE name = ?;"
	row := r.db.QueryRow(query, name)
	p := models.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}
