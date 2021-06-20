package service

import (
	"errors"
	"github.com/google/uuid"
	"movie/pkg/movieservice/application/data"
	"movie/pkg/movieservice/application/errors"
	"movie/pkg/movieservice/domain"
	"time"
)

type MovieService interface {
	AddMovie(request *data.AddMovieInput) error
	GetMovie(id string) (*domain.Movie, error)
	UpdateMovie(id string, request *data.UpdateMovieInput) error
	DeleteMovie(id string) error
}

type movieService struct {
	movieRepository domain.MovieRepository
}

func NewMovieService(movieRepo domain.MovieRepository) MovieService {
	return &movieService{
		movieRepository: movieRepo,
	}
}

func (m *movieService) AddMovie(request *data.AddMovieInput) error {
	if request.Name == "" {
		return serviceerrors.ErrorRequiredName
	}

	movieID := uuid.NewString()
	timestamp := time.Now().Unix()
	movieData := domain.Movie{
		ID:          movieID,
		CreatedAt:   timestamp,
		Name:        request.Name,
		Description: request.Description,
	}

	err := m.movieRepository.Add(movieData)

	return err
}

func (m *movieService) GetMovie(id string) (*domain.Movie, error) {
	movie, err := m.movieRepository.Get(id)

	if errors.Is(err, domain.ErrorMovieNotFound) {
		return nil, serviceerrors.ErrorNotFoundMovie
	}

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (m *movieService) UpdateMovie(id string, request *data.UpdateMovieInput) error {
	movie, err := m.movieRepository.Get(id)
	if movie == nil {
		return serviceerrors.ErrorNotFoundMovie
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
		return serviceerrors.ErrorNotFoundMovie
	}

	if err != nil {
		return err
	}

	movie.DeletedAt = time.Now().Unix()
	err = m.movieRepository.Delete(*movie)

	if err != nil {
		return err
	}

	return nil
}
