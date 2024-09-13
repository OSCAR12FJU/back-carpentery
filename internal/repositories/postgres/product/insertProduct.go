package product

import (
	"backend-carpintery/internal/domain"
	"fmt"
)

func (r Repository) InsertProduct(product domain.Product) (domain.Product, error) {

	query := `INSERT INTO products (title, image_url, price) VALUES ($1,$2,$3) RETURNING id`

	err := r.Db.QueryRow(query, product.Title, product.ImageURL, product.Price).Scan(&product.ID)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error al crear el producto: %w", err)
	}
	return product, nil

}
