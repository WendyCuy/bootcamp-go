package products

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
