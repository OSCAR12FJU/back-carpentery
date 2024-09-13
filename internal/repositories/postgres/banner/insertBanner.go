package banner

import (
	"backend-carpintery/internal/domain"
	"fmt"
)

func (r Repository) InsertBanner(banner domain.Banner) (domain.Banner, error) {
	query := `INSERT INTO banner (title, image_url) VALUES ($1,$2) RETURNING id`

	err := r.Db.QueryRow(query, banner.Title, banner.ImageURL).Scan(&banner.ID)

	if err != nil {
		return domain.Banner{}, fmt.Errorf("error al crear el banner: %w", err)
	}
	return banner, nil

}
