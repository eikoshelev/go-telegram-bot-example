package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Welcome *WelcomeRepo
}

func Init(dbPool pgxpool.Pool) *Repository {
	return &Repository{
		Welcome: &WelcomeRepo{db: dbPool},
	}
}
