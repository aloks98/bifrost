package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"os"
)

func Migrate() {
	log.Info("Database URI is: ", os.Getenv("POSTGRES_URL"))
	m, err := migrate.New("file://../db/migrations", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Cannot initialize golang-migrate: ", err)
	}
	log.Info("Starting Migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Cannot apply database migrations: ", err)
	}
}
