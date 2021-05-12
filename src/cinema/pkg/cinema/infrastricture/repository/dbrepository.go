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
	query := "SELECT created_at, name, description FROM movie WHERE id = ? "
	err := movieRepo.db.QueryRow(query, id).Scan(&movie.CreatedAt, &movie.Name, &movie.Description)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (movieRepo *DatabaseRepository) Update(movieData model.Movie) error {
	query := "UPDATE movie SET name = ?, description = ?, updated_at = ? WHERE id = ?"
	_, err := movieRepo.db.Exec(query, movieData.Name, movieData.Description, movieData.UpdatedAt, movieData.Id)
	if err != nil {
		return err
	}
	return nil
}

func (movieRepo *DatabaseRepository) Delete(movieData model.Movie) error {
	query := "UPDATE movie SET deleted_at = ? WHERE id = ?"
	_, err := movieRepo.db.Exec(query, movieData.DeletedAt, movieData.Id)
	if err != nil {
		return err
	}
	return nil
}
