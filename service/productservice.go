package service

import (
	"g-mysql/entity"
	"g-mysql/repository"
)

type IProductService interface {
	SaveProduct(product entity.Product) (entity.Product, error)
	FindById(Id int) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(Id int) (entity.Product, error)
	FindAllProduct(product []entity.Product) ([]entity.Product, error)
}

type Service struct {
	repository repository.IProductRepository
}

func NewService(r repository.IProductRepository) *Service {
	return &Service{r}
}

func (s *Service) SaveProduct(product entity.Product) (entity.Product, error) {
	product, err := s.repository.SaveProduct(product)
	return product, err
}

func (s *Service) FindById(Id int) (entity.Product, error) {
	product, err := s.repository.FindById(Id)
	return product, err
}

func (s *Service) UpdateProduct(product entity.Product) (entity.Product, error) {
	pr, err := s.repository.UpdateProduct(product)
	return pr, err
}

func (s *Service) DeleteProduct(Id int) (entity.Product, error) {
	return s.repository.DeleteProduct(Id)
}

func (s *Service) FindAllProduct(product []entity.Product) ([]entity.Product, error) {
	return s.repository.FindAllProduct(product)
}
