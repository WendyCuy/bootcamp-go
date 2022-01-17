package products

import (
	"fmt"

	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/pkg/store"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

// Con la inclusión de db store.Store se provee del archivo Store que esta en file.go
type repository struct {
	db store.Store
}

// Con db: db, permite inicializar el repository, se agrega como argumento el store
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// Para obtener los productos
/* Se declara una slice de Productos dentro del scope del métdodo ,
se le pasará esa variable al método Read q es la responsable de cargarla
con la información del archivo, la cual finalmente será retornada por el
método */

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

// Para obtener el último ID almacenado.
/* Al no tener más variables globales de último ID y productos, se debe
obtener el último ID del archivo.

Se obtendrá la información de productos guardada, en caso de no existir el
archivo, retornará com último ID cero. */
func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].ID, nil
}

// Guarda la información de producto, asignará el último ID a la variable y nos retorna la entidad Producto
func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {

	var ps []Product
	r.db.Read(&ps)
	p := Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

/* Se implementa la funcionalidad para actualizar el producto en memoria,
en caso que coincida con ID enviado, caso contrario retorna un error */

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}
	fmt.Printf("%+v", ps)
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return err
	}

	deleted := false
	var index int

	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil
}
