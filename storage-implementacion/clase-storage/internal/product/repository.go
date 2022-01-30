package product

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/models"
)

/* Optimización. Una buena práctica en la implementación de repositorios, es abstraer todas las queries
usadas en constantes fuera del código */

const (
	queryByname = "SELECT * FROM products WHERE name = ?;"
	queryStore  = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	queryGetOne = "select * from products where id = ?"
	queryUpdate = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?"
	queryGetAll = "select id, name, type, count, price from products"
	queryDelete = "DELETE FROM products WHERE id = ?"
)

// Repository encapsulates the storage of a product.
type Repository interface {
	GetByName(name string) (models.Product, error)
	Store(product models.Product) (models.Product, error)
	GetOne(id int) (models.Product, error)
	Update(product models.Product) error
	GetAll() ([]models.Product, error)
	Delete(id int) error
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
	row := r.db.QueryRow(queryByname, name)
	p := models.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *repository) Store(product models.Product) (models.Product, error) { // se inicializa la base
	stmt, err := r.db.Prepare(queryStore) // se prepara el SQL
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

func (r *repository) GetOne(id int) (models.Product, error) {
	row := r.db.QueryRow(queryGetOne, id)
	p := models.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price) // rows.Scan: lee los campos de la fila obtenida, y los almacena en las posiciones de memoria de las variables que se indican por parámetros.
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *repository) Update(product models.Product) error {

	stmt, err := r.db.Prepare(queryUpdate)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", product)

	res, err := stmt.Exec(product.Name, product.Type, product.Count,
		product.Price, product.ID)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()

	if err != nil {
		return err
	}

	fmt.Print(ra)

	return nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	var products []models.Product
	db := r.db
	rows, err := db.Query(queryGetAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slide products
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(queryDelete) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()     // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.Exec(id) // retorna un sql.Result y un error
	if err != nil {
		return err
	}
	return nil
}
