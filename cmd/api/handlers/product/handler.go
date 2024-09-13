package product

import "backend-carpintery/internal/ports"

type Handler struct {
	ProductService ports.ProductService
}
