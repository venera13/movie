package model

type MovieRepository interface {
	Add(movieData Movie) error
}

type RatingRepository interface {
}

type AddMovieInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Movie struct {
	Id          string
	Time        int64
	Name        string
	Description string
}
