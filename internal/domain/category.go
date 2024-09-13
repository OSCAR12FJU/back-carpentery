package domain

type Category struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}
