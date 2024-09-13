package ports

import "backend-carpintery/internal/domain"

type BannerService interface {
	CreateBanner(banner domain.Banner) (domain.Banner, error)
	// Edit()
	// Delete()
}

type BannerRepository interface {
	InsertBanner(banner domain.Banner) (domain.Banner, error)
	// Modify()
	// Delete()
}
