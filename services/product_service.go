package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(p *model.Product) error {
	return s.repo.Create(p)
}

func (s *ProductService) GetByID(id int) (*model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(p *model.Product) error {
	return s.repo.Update(p)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
