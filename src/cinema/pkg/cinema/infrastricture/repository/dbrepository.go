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
	query := "INSERT INTO movie(id, created_at, name, description) VALUES (?, ?, ?, ?)"
	_, err := movieRepo.db.Exec(query, movieData.Id, movieData.CreatedAt, movieData.Name, movieData.Description)
	if err != nil {
		return err
	}
	return nil
}

func (movieRepo *DatabaseRepository) Get(id string) (*model.Movie, error) {
	var movie model.Movie
	movie.Id = id
	query := "SELECT created_at, name, description FROM movie where id = ? "
	err := movieRepo.db.QueryRow(query, id).Scan(&movie.CreatedAt, &movie.Name, &movie.Description)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (movieRepo *DatabaseRepository) Update(movieData model.Movie) error {
	query := "UPDATE movie SET name = ?, description = ?, updated_at = ?"
	_, err := movieRepo.db.Exec(query, movieData.Name, movieData.Description, movieData.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (movieRepo *DatabaseRepository) Delete(id string) error {
	query := "DELETE FROM movie WHERE id = ?"
	_, err := movieRepo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
