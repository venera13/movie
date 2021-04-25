package main

import (
	service "cinema/pkg/cinema/application"
	"cinema/pkg/cinema/infrastricture/repository"
	"cinema/pkg/cinema/infrastricture/transport"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	killSignalChat := getKillSignalChan()

	config, err := parseEnv()
	if err != nil {
		log.Fatal(err)
	}
	var srv *http.Server
	srv, err = startServer(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	waitForKillSignal(killSignalChat)
	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

func startServer(config *config) (*http.Server, error) {
	serverUrl := config.ServeRESTAddress
	log.WithFields(log.Fields{
		"url": serverUrl,
	}).Info("starting the server")

	db, err := createDBConn(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	movieService := service.NewMovieService(repository.CreateMovieRepository(db))
	router := transport.Router(transport.NewServer(
		movieService,
	))
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv, nil
}

func createDBConn(config *config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@/%s", config.DBUser, config.DBPass, config.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
