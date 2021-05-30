package domain

type MovieRepository interface {
	Add(movieData Movie) error
	Get(id string) (*Movie, error)
	Update(movieData Movie) error
	Delete(movieData Movie) error
}

type RatingRepository interface {
}

type Movie struct {
	ID          string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
	Name        string
	Description string
}
