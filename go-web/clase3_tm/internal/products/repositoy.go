package products

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// Para obtener los productos
func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

// Para obtener el último ID almacenado
func (r *repository) LastID() (int, error) {
	return lastID, nil
}

// Guarda la información de producto, asignará el último ID a la variable y nos retorna la entidad Producto
func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

/* Se implementa la funcionalidad para actualizar el producto en memoria,
en caso que coincida con ID enviado, caso contrario retorna un error */

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
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
	return nil
}
