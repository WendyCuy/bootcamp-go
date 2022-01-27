package product

import (
	"database/sql"
	"log"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/models"
)

// Repository encapsulates the storage of a product.
type Repository interface {
	GetByName(name string) (models.Product, error)
	Store(name, productType string, count int, price float64) (models.Product, error)
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

func (r *repository) Store(product models.Product) (models.Product, error) { // se inicializa la base
	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return models.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	product.ID = int(insertedId)
	return product, nil
}
