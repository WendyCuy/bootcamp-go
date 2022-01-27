package product

import (
	"github.com/WendyCuy/bootcamp-go/storage-implementacion/clase-storage/internal/models"
)

// Se implementa la interface Servicio con sus métodos
type Service interface {
	GetByName(name string) (models.Product, error)
	Store(name, productType string, count int, price float64) (models.Product, error)
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

func (s *service) Store(name string, productType string, count int, price float64) (models.Product, error) {

	p := models.Product{Name: name, Type: productType, Count: count, Price: price}

	id, err := s.repository.Store(p)
	if err != nil {
		return models.Product{}, err
	}
}
