package transport

import (
	service "cinema/pkg/cinema/application"
	"cinema/pkg/cinema/model"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Server struct {
	movieService service.MovieService
}

func NewServer(service service.MovieService) *Server {
	return &Server{
		service,
	}
}

func Router(srv *Server) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/movie", srv.addMovie).Methods(http.MethodPost)
	return logMiddleware(r)
}

func (srv *Server) addMovie(w http.ResponseWriter, r *http.Request) {
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, err)
		return
	}
	defer r.Body.Close()
	var addMovieInput model.AddMovieInput
	err = json.Unmarshal(requestData, &addMovieInput)
	if err != nil {
		processError(w, err)
		return
	}

	err = srv.movieService.AddMovie(&addMovieInput)
	if err != nil {
		processError(w, err)
		return
	}
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

func processError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
