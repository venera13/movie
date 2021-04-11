package main

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	log.SetFormatter(&log.JSONFormatter{})

	killSignalChat := getKillSignalChan()
	waitForKillSignal(killSignalChat)

	config, err := parseEnv()
	if err != nil {
		log.Fatal(err)
	}
	srv, err = startServer(config)
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
	router :=

	return nil, nil
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