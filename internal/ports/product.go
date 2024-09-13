package ports

import (
	"backend-carpintery/internal/domain"
)

type ProductService interface {
	CreateProduct(product domain.Product) (domain.Product, error)
	// HandleImageUpload(file io.Reader, handler *multipart.FileHeader) (url string, err error)
	// Edit()
	// Delete()
}

type ProductRepository interface {
	InsertProduct(product domain.Product) (domain.Product, error)
	// Modify()
	// Delete()
}
