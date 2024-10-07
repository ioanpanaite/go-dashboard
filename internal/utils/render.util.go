package utils

import (
	"kub/dashboardES/internal/models"
	"kub/dashboardES/internal/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(c.Request().Context(), c.Response())
}

func GetTemplateData(data interface{}) models.TemplateData {
	// Check if the data contains a 'companies' element
	_, ok := data.(map[string]interface{})["companies"]

	// Add the base panels to the template data
	templateData := models.TemplateData{
		ShowCompanySearch: ok,
		Data:              data,
	}
	return templateData

}
func RenderError404(c echo.Context, fullPage bool) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return templates.Error404(fullPage).Render(c.Request().Context(), c.Response())
}

func RenderError500(c echo.Context, errorMessage string, fullPage bool) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return templates.Error500(errorMessage, fullPage).Render(c.Request().Context(), c.Response())
}
