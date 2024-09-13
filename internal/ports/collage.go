package ports

import "backend-carpintery/internal/domain"

type CollageService interface {
	CreateCollage(collage domain.Collage) (domain.Collage, error)
	// Edit()
	// Delete()
}

type CollageRepository interface {
	InsertCollage(collage domain.Collage) (domain.Collage, error)
	// Modify()
	// Delete()
}
