package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"os"
)

func Migrate() {
	postgres := os.Getenv("POSTGRES_URL")
	if postgres == "" {
		log.Fatal("POSTGRES_URL Environment Variable not set.. Exiting!!")
	}
	log.Info("Database URI is: ", os.Getenv("POSTGRES_URL"))
	m, err := migrate.New("file://../db/migrations", postgres)
	if err != nil {
		log.Fatal("Cannot initialize golang-migrate: ", err)
	}
	log.Info("Starting Migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Cannot apply database migrations: ", err)
	}
}
