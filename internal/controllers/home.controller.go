package controllers

import (
	"kub/dashboardES/internal/templates"
	"kub/dashboardES/internal/utils"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return utils.Render(c, templates.Index())
}
