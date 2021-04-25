package service

import (
	"cinema/pkg/cinema/model"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MovieService interface {
	AddMovie(request *model.AddMovieInput) error
	GetMovie(id string) (*model.Movie, error)
	UpdateMovie(id string, request *model.UpdateMovieInput) error
	DeleteMovie(id string) error
}

type movieService struct {
	movieRepository model.MovieRepository
}

func NewMovieService(movieRepo model.MovieRepository) MovieService {
	return &movieService{
		movieRepository: movieRepo,
	}
}

func (m *movieService) AddMovie(request *model.AddMovieInput) error {
	if len(request.Name) == 0 {
		return fmt.Errorf("the name of the movie is required")
	}
	movieId := uuid.NewString()
	timestamp := time.Now().Unix()
	movieData := model.Movie{
		Id:          movieId,
		CreatedAt:   timestamp,
		Name:        request.Name,
		Description: request.Description,
	}
	err := m.movieRepository.Add(movieData)
	if err != nil {
		return err
	}
	return nil
}

func (m *movieService) GetMovie(id string) (*model.Movie, error) {
	movie, err := m.movieRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *movieService) UpdateMovie(id string, request *model.UpdateMovieInput) error {
	movie, err := m.movieRepository.Get(id)
	if movie == nil {
		return fmt.Errorf("movie %s not found", id)
	}
	if err != nil {
		return err
	}
	if request.Name != "" {
		movie.Name = request.Name
	}
	if request.Description != "" {
		movie.Description = request.Description
	}
	movie.UpdatedAt = time.Now().Unix()
	err = m.movieRepository.Update(*movie)
	if err != nil {
		return err
	}
	return nil
}

func (m *movieService) DeleteMovie(id string) error {
	movie, err := m.movieRepository.Get(id)
	if movie == nil {
		return fmt.Errorf("movie %s not found", id)
	}
	err = m.movieRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
