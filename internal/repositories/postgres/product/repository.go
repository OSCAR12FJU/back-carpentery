package product

import (
	"database/sql"
)

type Repository struct {
	Db *sql.DB
}

// type ProductRepository interface {
// 	Insert(product domain.Product) (domain.Product, error)
// }
