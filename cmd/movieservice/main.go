package main

import (
	service "cinema/pkg/movieservice/application"
	"cinema/pkg/movieservice/infrastructure/repository"
	"cinema/pkg/movieservice/infrastructure/transport"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/sanketplus/go-mysql-lock"
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
	serverURL := config.ServeRESTAddress

	log.WithFields(log.Fields{
		"url": serverURL,
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
		Addr:    serverURL,
		Handler: router,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv, nil
}

func createDBConn(config *config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@%s/%s?multiStatements=true", config.DBUser, config.DBPass, config.DBAddress, config.DBName)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	//err = migrations(db)
	//if err != nil {
	//	log.Fatal(err)
	//
	//	return nil, err
	//}

	err = applyLocker(db)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return db, nil
}

func migrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)

		return err
	}

	var m *migrate.Migrate
	m, err = migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)

		return err
	}

	err = m.Up()

	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

func applyLocker(db *sql.DB) error {
	locker := gomysqllock.NewMysqlLocker(db)

	lock, err := locker.Obtain("foo")

	if err != nil {
		log.Fatal(err)

		return err
	}

	err = lock.Release()

	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
