package product

import (
	"backend-carpintery/internal/domain"
	"fmt"
)

func (s Services) CreateProduct(product domain.Product) (domain.Product, error) {
	if err := s.validateProduct(product); err != nil {
		return domain.Product{}, err
	}
	return s.Repo.InsertProduct(product)
}

func (s Services) validateProduct(product domain.Product) error {
	if product.Price <= 0 {
		return fmt.Errorf("precio inválido")
	}
	return nil
}
