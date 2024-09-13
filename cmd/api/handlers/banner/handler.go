package banner

import "backend-carpintery/internal/ports"

type Handler struct {
	BannerService ports.BannerService
}
