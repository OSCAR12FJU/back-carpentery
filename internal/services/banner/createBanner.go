package banner

import "backend-carpintery/internal/domain"

func (s Services) CreateBanner(banner domain.Banner) (domain.Banner, error) {
	return s.Repo.InsertBanner(banner)
}
