package storage

import (
	"database/sql"
	"github.com/KarinaRdk/WB-L2/tree/main/develop/dev11/Calendar/internal/domain"
)

type Storage struct {
	DB *sql.DB
}

func (s Storage) NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db,}
}

func (s Storage) Insert(e domain.Event) {

}