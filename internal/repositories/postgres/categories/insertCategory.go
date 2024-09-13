package categories

import (
	"backend-carpintery/internal/domain"
	"fmt"
)

func (r Repository) InsertCategory(category domain.Category) (domain.Category, error) {

	query := `INSERT INTO categories (title, image_url) VALUES ($1,$2) RETURNING id`

	err := r.Db.QueryRow(query, category.Title, category.ImageURL).Scan(&category.ID)
	if err != nil {
		return domain.Category{}, fmt.Errorf("error al crear categories: %w", err)
	}
	return category, nil
}
