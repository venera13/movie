package service

import (
	"cinema/pkg/cinema/model"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MovieService interface {
	AddMovie(request *model.AddMovieInput) error
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
		Time:        timestamp,
		Name:        request.Name,
		Description: request.Description,
	}
	err := m.movieRepository.Add(movieData)
	if err != nil {
		return err
	}
	return nil
}
