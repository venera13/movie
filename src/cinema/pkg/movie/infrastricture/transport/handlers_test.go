package transport

import (
	"cinema/pkg/movie/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type movieService struct{}

const mainMovieID = "kjrhfkjrhkrhr"

func (m movieService) AddMovie(request *model.AddMovieInput) error {
	return nil
}

func (m movieService) GetMovie(id string) (*model.Movie, error) {
	return &model.Movie{
		ID:          "7e7e7a38-b4ac-4ceb-b110-bf55e80f31fb",
		CreatedAt:   1619367396,
		Name:        "Аватар",
		Description: "В 2154 году полезные ископаемые планеты Земля практически исчерпаны.",
	}, nil
}

func (m movieService) UpdateMovie(id string, request *model.UpdateMovieInput) error {
	return nil
}

func (m movieService) DeleteMovie(id string) error {
	return nil
}

func TestAddMovie(t *testing.T) {
	srv := Server{movieService{}}
	w := httptest.NewRecorder()
	bodyReader := strings.NewReader(`{"name": "Аватар", "description": " В 2154 году полезные ископаемые планеты Земля практически исчерпаны."}`)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/movie", bodyReader)
	srv.addMovie(w, r)
	response := w.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d", response.StatusCode, http.StatusOK)
	}
	defer response.Body.Close()
}

func TestGetMovie(t *testing.T) {
	srv := Server{movieService{}}
	w := httptest.NewRecorder()
	movieID := mainMovieID
	r := httptest.NewRequest(http.MethodGet, "/api/v1/movie/{ID}", nil)
	r = mux.SetURLVars(r, map[string]string{"ID": movieID})
	srv.getMovie(w, r)
	response := w.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	var movies model.Movie
	if err = json.Unmarshal(jsonString, &movies); err != nil {
		t.Errorf("Can't parse json: %s response with error %v", jsonString, err)
	}
}

func TestUpdateMovie(t *testing.T) {
	srv := Server{movieService{}}
	w := httptest.NewRecorder()
	bodyReader := strings.NewReader(`{"name": "Аватар", "description": "Новый description"}`)
	movieID := mainMovieID
	r := httptest.NewRequest(http.MethodPut, "/api/v1/movie/{ID}", bodyReader)
	r = mux.SetURLVars(r, map[string]string{"ID": movieID})
	srv.updateMovie(w, r)
	response := w.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d", response.StatusCode, http.StatusOK)
	}
	defer response.Body.Close()
}

func TestDeleteMovie(t *testing.T) {
	srv := Server{movieService{}}
	w := httptest.NewRecorder()
	movieID := mainMovieID
	r := httptest.NewRequest(http.MethodDelete, "/api/v1/movie/{ID}", nil)
	r = mux.SetURLVars(r, map[string]string{"ID": movieID})
	srv.deleteMovie(w, r)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d", response.StatusCode, http.StatusOK)
	}
	defer response.Body.Close()
}
