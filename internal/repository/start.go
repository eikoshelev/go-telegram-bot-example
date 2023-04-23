package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Start struct {
	db pgxpool.Pool
}

/*
	there should be object management methods at the repository level (storage state)
*/

func (s *Start) SomeAction() {}
