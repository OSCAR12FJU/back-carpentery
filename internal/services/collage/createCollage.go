package collage

import "backend-carpintery/internal/domain"

func (s Services) CreateCollage(collage domain.Collage) (domain.Collage, error) {
	return s.Repo.InsertCollage(collage)
}

// func (s Services) CreateProduct(product domain.Product) (domain.Product, error) {
// 	if err := s.validateProduct(product); err != nil {
// 		return domain.Product{}, err
// 	}
// 	return s.Repo.Insert(product)
// }
