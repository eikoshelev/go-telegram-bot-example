package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Start   *Start
	Profile *Profile
}

// function to initialize all repositories
func Init(dbPool pgxpool.Pool) *Repository {
	return &Repository{
		Start:   &Start{db: dbPool},
		Profile: &Profile{db: dbPool},
	}
}
