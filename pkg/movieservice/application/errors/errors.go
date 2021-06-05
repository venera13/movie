package serviceerrors

import "errors"

var RequiredNameError = errors.New("the name of the movie is required")
var NotFoundMovieError = errors.New("movie not found")
