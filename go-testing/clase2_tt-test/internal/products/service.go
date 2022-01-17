package products

// Se implementa la interface Servicio con sus metodos
type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
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

/* Dentro del servicio se llama al repositorio para que proceda a
actualizar el producto. */

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {

	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
