package ports

import "backend-carpintery/internal/domain"

type CategoryService interface {
	CreateCategory(card domain.Category) (domain.Category, error)
	// Edit()
	// Delete()
}

type CategoryRepository interface {
	InsertCategory(card domain.Category) (domain.Category, error)
	// Modify()
	// Delete()
}
