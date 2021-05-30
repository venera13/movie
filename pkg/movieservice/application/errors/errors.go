package serviceerrors

import "errors"

var RequiredNameError = errors.New("the name of the movieservice is required")
var NotFoundMovieError = errors.New("movieservice not found")
