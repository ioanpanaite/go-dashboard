package middlewares

import (
	"encoding/json"
	"fmt"
	"kub/dashboardES/internal/templates"
	"kub/dashboardES/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/allegro/bigcache"
	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

var cache *bigcache.BigCache

func InitCache() {
	var err error
	cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(24 * time.Hour))
	if err != nil {
		log.Fatal(err)
	}
}

// Middleware to check for a specific cookie
func CheckCookieMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the underlying Request from Echo's context
		r := c.Request()

		// Logging the request
		log4go.LOGGER("info").Info("Request received")

		// Attempt to retrieve the cookie
		cookie, err := r.Cookie("session_token") // Replace 'session_token' with your cookie name
		if err != nil {
			if err == http.ErrNoCookie {
				// Log and handle the case where the cookie is not found
				log4go.LOGGER("error").Error("Cookie not found")
				return utils.Render(c, templates.NoLogin())
			}
			// Log and handle other potential errors
			log4go.LOGGER("error").Error("Error retrieving cookie: " + err.Error())
			return err
		}

		// If the cookie is found, log its value (consider if this is safe to do depending on your application's security requirements)
		log4go.LOGGER("info").Info("Cookie found: " + cookie.Value)

		// Check the cache first
		data, err := cache.Get(cookie.Value)
		log4go.LOGGER("info").Info("Cache get")
		if err == nil {
			// Use the cached data
			r.Header.Add("user-id", string(data))
			log4go.LOGGER("info").Info("cache" + string(data))
			return next(c)
		} else if err == bigcache.ErrEntryNotFound {
			// Create a new request to validate the cookie
			req, err := http.NewRequest("GET", "https://accounts-api.connectresources.ae/api/v1/authentication/session", nil)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			// Set the headers
			req.Header.Set("x-tenant", "eservices")
			req.Header.Set("authorization", "Bearer "+cookie.Value)
			req.Header.Set("Content-Type", "application/json")

			// Send the request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			defer resp.Body.Close()

			responseData, err := parseJSONResponse(resp)

			// Check the response status
			if resp.StatusCode != http.StatusOK {
				//TODO: Render nologin page
				// tmpl, err := template.ParseFiles("templates/nologin.html")
				// if err == http.ErrNoCookie {
				// 	// Handle the absence of the cookie, e.g., redirect or return an error
				// 	return c.String(http.StatusUnauthorized, "Cookie required")
				// }

				// tmpl.Execute(w, map[string]interface{}{
				// 	"BaseRoute": "/dash", // Replace with your actual base panels
				// })
				return utils.Render(c, templates.NoLogin())
			}

			userId, ok := responseData["id"]
			if !ok {
				// Handle the case where the id key doesn't exist
				log4go.LOGGER("error").Error("User ID not found in the response data")
				return nil
			}
			userIdStr := fmt.Sprintf("%v", userId)

			// Cache the data
			err = cache.Set(cookie.Value, []byte(userIdStr)) // Replace "38" with the actual user ID
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			// Add the received header to the request
			r.Header.Add("user-id", userIdStr)
			log4go.LOGGER("info").Info("auth" + userIdStr)
			// If the cookie is valid, call the next handler
			return next(c)
		} else {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

}

// Middleware to check for a specific cookie
func AddUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the underlying Request and ResponseWriter from Echo's context
		r := c.Request()
		// w := c.Response().Writer

		log4go.LOGGER("info").Info("Cookie! fix")

		// Add the received header to the request
		r.Header.Add("user-id", "38")

		// If the cookie is valid, call the next handler
		return next(c)
	}
}

func parseJSONResponse(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()

	var data map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
