package main

import (
	"fmt"
	"kub/dashboardES/internal/logger"
	"kub/dashboardES/internal/middlewares"
	"kub/dashboardES/internal/server"

	"github.com/jeanphorn/log4go"
)

func main() {
	logger.LoggerInit()
	server := server.NewServer()

	middlewares.InitCache()

	err := server.ListenAndServe()
	if err != nil {
		log4go.LOGGER("error").Error("cannot start server: %s", err)
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
