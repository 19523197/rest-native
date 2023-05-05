package repository

import "database/sql"

type BaseRepository struct {
	Sql *sql.DB
}

func NewBaseRepository(s *sql.DB) *BaseRepository {
	return &BaseRepository{Sql: s}
}
