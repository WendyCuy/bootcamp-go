package products

// Se implementa la interface Servicio con sus metodos
type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
}
type service struct {
	repository Repository
}

// Se implementa una función que recibe un repositorio y nos devuelve
// el servicio que se utilizará, instanciado
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

/* Se implementa el método GetAll que se encargará de pasarle la tarea
al repositorio y nos retorna un array de Productos */
func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

/* El método Store se encargará de pasarle la tarea de obtener el ultimo ID
y guardar el producto al Repositorio, el servicio se encargará de incrementar
el ID */

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++
	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}
