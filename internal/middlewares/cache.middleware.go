package middlewares

import (
	"github.com/labstack/echo/v4"
)

func DisableCachingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Disable caching by setting Cache-Control headers
		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		return next(c)
	}
}
