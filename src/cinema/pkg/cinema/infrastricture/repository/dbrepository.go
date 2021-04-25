package repository

import (
	"cinema/pkg/cinema/model"
	"database/sql"
)

func CreateMovieRepository(db *sql.DB) model.MovieRepository {
	return &DatabaseRepository{
		db: db,
	}
}

type DatabaseRepository struct {
	db *sql.DB
}

func (movieRepo *DatabaseRepository) Add(movieData model.Movie) error {
	query := "INSERT INTO movie(id, time, name, description) VALUES (?, ?, ?, ?)"
	_, err := movieRepo.db.Exec(query, movieData.Id, movieData.Time, movieData.Name, movieData.Description)
	if err != nil {
		return err
	}
	return nil
}
