package data

type AddMovieInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateMovieInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
