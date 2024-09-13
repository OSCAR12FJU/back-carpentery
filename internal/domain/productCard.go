package domain

type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}
