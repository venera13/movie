package repository

import (
	"database/sql"
	"cinema/pkg/cinema/model"
)

func CreateMovieRepository(db *sql.DB) model.MovieRepository {
	return &DatabaseRepository{
		db: db,
	}
}

func CreateRatingRepository(db *sql.DB) model.RatingRepository {
	return &DatabaseRepository{
		db: db,
	}
}

type DatabaseRepository struct {
	db *sql.DB
}