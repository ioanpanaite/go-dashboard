package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeanphorn/log4go"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *sql.DB
}

var (
	databaseHost = os.Getenv("DATABASE_HOST")
	databaseUser = os.Getenv("DATABASE_USER")
	databasePass = os.Getenv("DATABASE_PASS")
	databaseName = os.Getenv("DATABASE_NAME")
	DbClient     *sql.DB
	Repo         MySQLRepository
)

func New() Service {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", databaseUser, databasePass, databaseHost, databaseName))
	if err != nil {
		log4go.LOGGER("error").Error("db down: %v", err)
		log.Fatal(err)
	}
	// db.SetConnMaxLifetime(0)
	// db.SetMaxIdleConns(50)
	// db.SetMaxOpenConns(50)
	DbClient = db
	Repo = *NewMySQLRepository(DbClient)
	s := &service{db: db}
	return s

}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err := DbClient.PingContext(ctx)
	if err != nil {
		log4go.LOGGER("error").Error("db down: %v", err)
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
