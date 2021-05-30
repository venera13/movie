package transport

import (
	service "cinema/pkg/movie/application"
	"cinema/pkg/movie/model"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
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
	s.HandleFunc("/movie/{ID}", srv.getMovie).Methods(http.MethodGet)
	s.HandleFunc("/movie/{ID}", srv.updateMovie).Methods(http.MethodPut)
	s.HandleFunc("/movie/{ID}/delete", srv.deleteMovie).Methods(http.MethodPut)
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

func (srv *Server) getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["ID"] == "" {
		http.Error(w, "id of the movie is required", http.StatusBadRequest)
		return
	}

	movie, err := srv.movieService.GetMovie(vars["ID"])
	if err != nil {
		writeResponse(w, http.StatusNotFound, "Movie not found")
		return
	}
	var b []byte
	b, err = json.Marshal(movie)
	writeResponse(w, http.StatusOK, string(b))
}

func (srv *Server) updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["ID"] == "" {
		http.Error(w, "id of the movie is required", http.StatusBadRequest)
		return
	}

	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, err)
		return
	}
	defer r.Body.Close()
	var updateMovieInput model.UpdateMovieInput
	err = json.Unmarshal(requestData, &updateMovieInput)
	if err != nil {
		processError(w, err)
		return
	}

	err = srv.movieService.UpdateMovie(vars["ID"], &updateMovieInput)
	if err != nil {
		processError(w, err)
		return
	}
}

func (srv *Server) deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["ID"] == "" {
		http.Error(w, "id of the movie is required", http.StatusBadRequest)
		return
	}

	err := srv.movieService.DeleteMovie(vars["ID"])
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

func writeResponse(w http.ResponseWriter, status int, response string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_, err := io.WriteString(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}