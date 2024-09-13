package category

import "backend-carpintery/internal/ports"

type Handler struct {
	CategoryService ports.CategoryService
}
