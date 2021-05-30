package repository

import (
	"cinema/pkg/movieservice/domain"
	"database/sql"
	"errors"
)

func CreateMovieRepository(db *sql.DB) domain.MovieRepository {
	return &DatabaseRepository{
		db: db,
	}
}

type DatabaseRepository struct {
	db *sql.DB
}

func (movieRepo *DatabaseRepository) Add(movieData domain.Movie) error {
	query := "INSERT INTO movieservice(id, created_at, name, description) VALUES (?, ?, ?, ?)"
	_, err := movieRepo.db.Exec(query, movieData.ID, movieData.CreatedAt, movieData.Name, movieData.Description)

	return err
}

func (movieRepo *DatabaseRepository) Get(id string) (*domain.Movie, error) {
	var movie domain.Movie
	movie.ID = id
	query := "SELECT created_at, name, description FROM movieservice WHERE id = ? "
	err := movieRepo.db.QueryRow(query, id).Scan(&movie.CreatedAt, &movie.Name, &movie.Description)

	if errors.Is(err, sql.ErrNoRows) {
		return &movie, domain.ErrorMovieNotFound
	}

	return &movie, err
}

func (movieRepo *DatabaseRepository) Update(movieData domain.Movie) error {
	query := "UPDATE movieservice SET name = ?, description = ?, updated_at = ? WHERE id = ?"
	_, err := movieRepo.db.Exec(query, movieData.Name, movieData.Description, movieData.UpdatedAt, movieData.ID)

	return err
}

func (movieRepo *DatabaseRepository) Delete(movieData domain.Movie) error {
	query := "UPDATE movieservice SET deleted_at = ? WHERE id = ?"
	_, err := movieRepo.db.Exec(query, movieData.DeletedAt, movieData.ID)

	return err
}
