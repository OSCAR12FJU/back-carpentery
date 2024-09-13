package collage

import (
	"backend-carpintery/internal/domain"
	"fmt"
)

func (r Repository) InsertCollage(collage domain.Collage) (domain.Collage, error) {
	query := `INSERT INTO collage (title, image_url) VALUES ($1,$2) RETURNING id`

	err := r.Db.QueryRow(query, collage.Title, collage.ImageURL).Scan(&collage.ID)
	if err != nil {
		return domain.Collage{}, fmt.Errorf("error al crear el collage: %w", err)
	}
	return collage, nil
}
