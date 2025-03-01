package server

import (
	"fmt"
	"kub/dashboardES/internal/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jeanphorn/log4go"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log4go.LOGGER("info").Info("App Started on port: %d", port)
	return server
}
