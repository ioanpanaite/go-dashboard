package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

func LoggerInit() {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Define the path to the "logs" folder in the root directory
	logsFolderPath := filepath.Join(workingDir, "logs")

	// Check if the "logs" folder already exists
	if _, err := os.Stat(logsFolderPath); os.IsNotExist(err) {
		// Create the "logs" folder if it doesn't exist
		err := os.Mkdir(logsFolderPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("The 'logs' folder has been created.")
	} else if err != nil {
		log.Fatal(err)
	}
	log4go.LoadConfiguration("./log.json")

}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		ReqLogger(false, "[%v] Requested resource: %v", req.Method, req.RequestURI)
		// log4go.LOGGER("request").Info("REQ [%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		err := next(c)
		if res.Status >= 400 {
			ReqLogger(true, "[%v]:[%d] Response: %v", req.Method, res.Status, req.RequestURI)
		} else {
			ReqLogger(false, "[%v]:[%d] Response: %v", req.Method, res.Status, req.RequestURI)
		}
		return err
	}
}

func ReqLogger(err bool, str string, a ...interface{}) {
	message := fmt.Sprintf(str, a...)
	if err {
		log4go.LOGGER("request").Error(message)
	} else {
		log4go.LOGGER("request").Info(message)
	}
}
