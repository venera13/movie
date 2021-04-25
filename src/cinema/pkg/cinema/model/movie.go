package model

type MovieRepository interface {
	Add(movieData Movie) error
	Get(id string) (*Movie, error)
	Update(movieData Movie) error
	Delete(id string) error
}

type RatingRepository interface {
}

type AddMovieInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateMovieInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Movie struct {
	Id          string
	CreatedAt   int64
	UpdatedAt   int64
	Name        string
	Description string
}
