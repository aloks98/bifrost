package db

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
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

func CreateConn() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		os.Exit(1)
	}
	return conn
}
