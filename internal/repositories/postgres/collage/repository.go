package collage

import (
	"database/sql"
)

type Repository struct {
	Db *sql.DB
}
