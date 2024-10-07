package server

import (
	"kub/dashboardES/internal/controllers"
	"kub/dashboardES/internal/logger"
	"kub/dashboardES/internal/middlewares"
	"kub/dashboardES/internal/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var isProduction = os.Getenv("PROD") == "true"

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	//Middlewares

	e.Static("/", "public")

	e.Use(middleware.Recover())
	e.Use(logger.LoggerMiddleware)
	//Disable Caching for Development
	if !isProduction {
		e.Use(middlewares.DisableCachingMiddleware)
	}
	e.Use(middlewares.DisableCachingMiddleware)
	//Home Route

	if isProduction {
		e.Use(middlewares.CheckCookieMiddleware)
	} else {
		e.Use(middlewares.AddUserMiddleware)
	}

	home := e.Group("")
	home.GET("", controllers.HomeHandler)
	home.RouteNotFound("/*", func(c echo.Context) error {
		return utils.RenderError404(c, true)
	})

	//Panels Route
	// panels := e.Group("/panels")
	panels := home // for production

	if isProduction {
		panels = home
	} else {
		panels = e.Group("/panels")
	}
	panels.GET("", controllers.HomeHandler)

	panels.GET("", controllers.HomeHandler)
	panels.GET("/onboarding", controllers.OnboardingHandler)
	panels.GET("/onboarding/search", controllers.OnboardingSearchHandler)

	panels.GET("/offboarding", controllers.OffboardingHandler)
	panels.GET("/offboarding/search", controllers.OffboardingSearchHandler)

	panels.GET("/missing", controllers.MissingHandler)
	panels.GET("/missing/search", controllers.MissingSearchHandler)
	panels.GET("/missing/get/:id", controllers.MissingGetByIdHandler)
	panels.POST("/missing/update/:id", controllers.MissingUpdateHandler)

	panels.GET("/insurance", controllers.InsuranceHandler)
	panels.GET("/insurance/search", controllers.InsuranceSearchHandler)
	panels.POST("/insurance/update/:id", controllers.InsuranceUpdateHandler)
	panels.GET("/insurance/get/:id", controllers.InsuranceGetByIdHandler)

	panels.GET("/companies/search", controllers.CompaniesSearchHandler)
	panels.GET("/employee/search", controllers.EmployeesSearchHandler)

	panels.Static("/", "public")
	panels.RouteNotFound("/*", func(c echo.Context) error {
		return utils.RenderError404(c, true)
	})

	//Invoicing Route
	invoicing := panels.Group("/invoicing")

	invoicing.GET("", func(c echo.Context) error {
		return utils.RenderError404(c, true)
	})
	invoicing.RouteNotFound("/*", func(c echo.Context) error {
		return utils.RenderError404(c, true)
	})
	invoicing.GET("/charges", controllers.ChargesHandler)
	invoicing.GET("/chargedef/search", controllers.ChargeDefSearchHandler)
	invoicing.GET("/charges/:companyId/:chargedefId", controllers.ChargesByCompanyAndIdHandler)
	invoicing.POST("/charges/:companyId/:chargedefId", controllers.ChargeCreateHandler)
	invoicing.DELETE("/charges/:companyId/:chargedefId", controllers.ChargeDeleteHandler)
	invoicing.GET("/charges/charges", controllers.GetChargesHandler)
	//Filter company charges
	invoicing.GET("/chargedef/filterCompanyCharges", controllers.FilterCompanyChargesHandler)
	invoicing.GET("/chargedef/filterEmployeeCharges", controllers.FilterEmployeeChargesHandler)

	//Invoicing/invoice
	invoicing.GET("/invoice", controllers.InvoiceHandeler)
	invoicing.GET("/invoice/search", controllers.GetInternalInvoicesHandler)
	invoicing.GET("/clientInvoice/search", controllers.GetClientInvoiceSearch)
	invoicing.GET("/invoice/form/:invoiceId", controllers.GetInvoiceFormHandler)
	invoicing.POST("/invoice/:invoiceId", controllers.CreateUpdateInvoiceHandler)
	invoicing.DELETE("/invoice/:invoiceId", controllers.DeleteInvoiceHandler)
	invoicing.GET("/invoice/manage/:companyId/:invoiceId", controllers.ManageInvoiceHandler)
	invoicing.GET("/invoice/preview/:companyId/:invoiceId", controllers.PreviewInvoiceHandler)
	invoicing.GET("/invoice/sign/:invoiceId", controllers.GetInvoiceSignHandler)

	//Database Ping
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) _RegisterRoutes() http.Handler {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	//Middlewares
	e.Static("/", "public")
	e.Static("/panels", "public")
	e.Use(middleware.Recover())
	e.Use(logger.LoggerMiddleware)
	//Disable Caching for Development
	if !isProduction {
		e.Use(middlewares.DisableCachingMiddleware)
	}
	e.Use(middlewares.DisableCachingMiddleware)
	//Home Route
	if isProduction {
		e.Use(middlewares.CheckCookieMiddleware)
	} else {
		e.Use(middlewares.AddUserMiddleware)
	}

	home := e.Group("")

	home.GET("", controllers.HomeHandler)

	home.GET("/onboarding", controllers.OnboardingHandler)
	home.GET("/onboarding/search", controllers.OnboardingSearchHandler)

	home.GET("/offboarding", controllers.OffboardingHandler)
	home.GET("/offboarding/search", controllers.OffboardingSearchHandler)

	home.GET("/missing", controllers.MissingHandler)
	home.GET("/missing/search", controllers.MissingSearchHandler)
	home.GET("/missing/get/:id", controllers.MissingGetByIdHandler)
	home.POST("/missing/update/:id", controllers.MissingUpdateHandler)

	home.GET("/insurance", controllers.InsuranceHandler)
	home.GET("/insurance/search", controllers.InsuranceSearchHandler)
	home.POST("/insurance/update/:id", controllers.InsuranceUpdateHandler)
	home.GET("/insurance/get/:id", controllers.InsuranceGetByIdHandler)

	home.GET("/companies/search", controllers.CompaniesSearchHandler)
	home.GET("/employee/search", controllers.EmployeesSearchHandler)

	//Invoicing Route
	invoicing := home.Group("/invoicing")

	invoicing.GET("", func(c echo.Context) error {
		return utils.RenderError404(c, true)
	})
	invoicing.RouteNotFound("/*", func(c echo.Context) error {
		return utils.RenderError404(c, true)
	})
	invoicing.GET("/chargedef/search", controllers.ChargeDefSearchHandler)
	invoicing.GET("/charges", controllers.ChargesHandler)
	invoicing.GET("/charges/:companyId/:chargedefId", controllers.ChargesByCompanyAndIdHandler)
	invoicing.POST("/charges/:companyId/:chargedefId", controllers.ChargeCreateHandler)
	invoicing.DELETE("/charges/:companyId/:chargedefId", controllers.ChargeDeleteHandler)

	//Database Ping
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
