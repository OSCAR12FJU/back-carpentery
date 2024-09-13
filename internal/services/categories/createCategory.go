package categories

import (
	"backend-carpintery/internal/domain"
)

func (s Services) CreateCategory(category domain.Category) (domain.Category, error) {
	return s.Repo.InsertCategory(category)
}
