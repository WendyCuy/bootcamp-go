package product

import (
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/models"
)

// Se implementa la interface Servicio con sus métodos
type Service interface {
	GetByName(name string) (models.Product, error)
	Store(name string, typeProduct string, count int, price float64) (models.Product, error)
	GetOne(id int) (models.Product, error)
	Update(id int, name string, typeProduct string, count int, price float64) (models.Product, error)
	GetAll() ([]models.Product, error)
}
type service struct {
	repository Repository
}

// Se implementa la función que recibe el repositorio y devuelve el servicio que se utilizará.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetByName(name string) (models.Product, error) {
	product, err := s.repository.GetByName(name)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *service) Store(name string, typeProduct string, count int, price float64) (models.Product, error) {

	newProduct := models.Product{
		Name:  name,
		Type:  typeProduct,
		Count: count,
		Price: price,
	}

	res, err := s.repository.Store(newProduct)

	if err != nil {
		return models.Product{}, err
	}

	return res, nil
}

func (s *service) GetOne(id int) (models.Product, error) {

	product, err := s.repository.GetOne(id)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

func (s *service) Update(id int, name string, typeProduct string, count int, price float64) (models.Product, error) {

	product, err := s.repository.GetOne(int(id))

	if err != nil {
		return models.Product{}, err
	}

	productToUp := models.Product{Name: name, Type: typeProduct, Count: count, Price: price, ID: id}

	err = s.repository.Update(productToUp)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) GetAll() ([]models.Product, error) {

	products, err := s.repository.GetAll()

	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}
