package serviceerrors

import "errors"

var ErrorRequiredName = errors.New("the name of the movie is required")
var ErrorNotFoundMovie = errors.New("movie not found")
