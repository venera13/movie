package transport

import (
	"cinema/pkg/cinema/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	movieRepository  model.MovieRepository
	ratingRepository model.RatingRepository
}

func NewServer(movieRepo model.MovieRepository, ratingRepo model.RatingRepository) *Server {
	return &Server{
		movieRepository:  movieRepo,
		ratingRepository: ratingRepo,
	}
}

func Router(s *Server) http.Handler {
	r := mux.NewRouter()
	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
			"time":       time.Since(start),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}
