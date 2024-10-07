package controllers

import (
	"kub/dashboardES/internal/database"
	"kub/dashboardES/internal/models"
	"kub/dashboardES/internal/templates/components"
	insurance_view "kub/dashboardES/internal/templates/panels/insurance"
	missing_view "kub/dashboardES/internal/templates/panels/missing"
	offboarding_view "kub/dashboardES/internal/templates/panels/offboarding"
	onboarding_view "kub/dashboardES/internal/templates/panels/onboarding"
	"kub/dashboardES/internal/utils"
	"strconv"
	"strings"

	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

// Insurance Controllers "/panels/insurance"
func InsuranceHandler(c echo.Context) error {
	data := utils.GetTemplateData(map[string]interface{}{
		"module": "dashboard",
		"page":   "insurance",
	})
	dataMap, ok := data.Data.(map[string]interface{})
	if !ok {
		log4go.LOGGER("error").Error("Error parsing data")
		return utils.RenderError500(c, "Error parsing data", false)
	}
	module, _ := dataMap["module"].(string)
	page, _ := dataMap["page"].(string)
	return utils.Render(c, insurance_view.InsuranceView(data.ShowCompanySearch, module, page))
}
func InsuranceSearchHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("search!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	page, _ := strconv.Atoi(c.QueryParam("page"))

	log4go.LOGGER("info").Info("Query parameters: %s", c.Request().URL.Query())
	searchQuery := c.QueryParam("searchInput")
	searchEmployee := c.QueryParam("searchEmployee")
	searchCompany := c.QueryParam("searchCompany")
	searchStatus := c.QueryParam("searchStatus")
	searchCategory := c.QueryParam("searchCategory")

	companyIds := c.QueryParam("companyIds")
	employeeIds := c.QueryParam("employeeIds")

	seePending := 1
	if c.QueryParam("seePending") == "true" {
		seePending = 1
	}

	log4go.LOGGER("info").Info("%d", seePending)
	items, err := database.Repo.GetDasboardSearch(userID, "", page, 100, searchQuery, companyIds, employeeIds, 1, 1, 0, seePending, "1", searchEmployee, searchCompany, searchStatus, searchCategory, "", "", "") // Implement SearchDasboard in your repository
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// Data to be passed to the template
	return utils.Render(c, insurance_view.InsuranceTBody("", nil, items))
}
func InsuranceGetByIdHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("get!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	id := c.Param("id")
	log4go.LOGGER("info").Info(id)

	// get edit from header
	edit := c.QueryParam("edit")
	log4go.LOGGER("info").Info(edit)

	id_ := c.QueryParam("id")
	log4go.LOGGER("info").Info(id_)
	companyId := c.QueryParam("companyId")
	log4go.LOGGER("info").Info(companyId)
	// Fetch the current state of the item using the id
	// Return the fetched daSta or handle errors

	items, err := database.Repo.GetDasboardSearch(userID, id, 1, 1, "", companyId, id_, -1, -1, -1, 1, "1", "", "", "", "", "", "", "") // Implement SearchDasboard in your repository
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	BusinessCategories, err := database.Repo.GetBusinessCategories(userID)
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	log4go.LOGGER("info").Info("BusinessCategories")
	// Data to be passed to the template

	// Render the template with the fetched data
	return utils.Render(c, insurance_view.InsuranceTBody(edit, BusinessCategories, items))
}
func InsuranceUpdateHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("update!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	id := c.QueryParam("id")
	cid := c.QueryParam("companyId")
	cuid := c.QueryParam("companyUserId")
	log4go.LOGGER("info").Info("id=%v,cid=%v,cuid=%v", id, cid, cuid)

	cardNumber := c.Request().FormValue("cardNumber")
	additionDate := c.Request().FormValue("additionDate")
	cancellationDate := c.Request().FormValue("cancellationDate")
	visaExpirationDate := c.Request().FormValue("visaExpirationDate")
	//category := r.FormValue("category")
	log4go.LOGGER("info").Info("Update Data:: cardNumber=%v,additionDate=%v,cancellationDate=%v,visaExpirationDate=%v", cardNumber, additionDate, cancellationDate, visaExpirationDate)
	// Call CrudDasboards function with extracted variables
	_, err := database.Repo.CrudDasboards(
		userID,             // User ID, assuming it's defined elsewhere
		id,                 // ID, assuming it's defined elsewhere
		cid,                // Company ID, assuming it's not provided in form
		cuid,               // Company user ID, assuming it's not provided in form
		"",                 // Emirates ID
		"",                 // Passport
		"",                 // Email
		"",                 // Phone
		"",                 // Ejari Number
		"",                 // Ejari Expiration
		"",                 // Labor Card ID
		visaExpirationDate, // Labor Card Expiration
		"",                 // Visa Number
		"",                 // Visa Expiration
		"",                 // Labour ID, assuming it's not provided in form
		cardNumber,         // Card Number, assuming it's not provided in form
		additionDate,       // Insurances Start Date, assuming it's not provided in form
		cancellationDate,   // Insurances End Date, assuming it's not provided in form
		"",                 // Business Category ID, assuming it's not provided in form
		"",                 // Dewa Number
		"",                 // Dewa Expiration
	)

	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	items, err := database.Repo.GetDasboardSearch(userID, id, 1, 1, "", cid, id, -1, -1, -1, 1, "1", "", "", "", "", "", "", "") // Implement SearchDasboard in your repository
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// Render the template with the fetched data
	return utils.Render(c, insurance_view.InsuranceTBody("", nil, items))
}

// Onboarding Controllers "/panels/onboarding"
func OnboardingHandler(c echo.Context) error {
	data := utils.GetTemplateData(map[string]interface{}{
		"module": "dashboard",
		"page":   "onboarding",
	})
	dataMap, ok := data.Data.(map[string]interface{})
	if !ok {
		log4go.LOGGER("error").Error("Error parsing data")
		return utils.RenderError500(c, "Error parsing data", false)
	}
	module, _ := dataMap["module"].(string)
	page, _ := dataMap["page"].(string)
	return utils.Render(c, onboarding_view.OnboardingView(data.ShowCompanySearch, module, page))
}
func OnboardingSearchHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("onboarding search!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	page, _ := strconv.Atoi(c.QueryParam("page"))

	searchQuery := c.QueryParam("searchInput")
	searchEmployee := c.QueryParam("searchEmployee")
	searchCompany := c.QueryParam("searchCompany")
	searchStatus := c.QueryParam("searchStatus")
	searchAccMan := c.QueryParam("searchAccMan")
	searchSalesStaff := c.QueryParam("searchSalesStaff")

	companyIds := c.QueryParam("companyIds")
	employeeIds := c.QueryParam("employeeIds")

	seePending := 1
	if c.QueryParam("seePending") == "true" {
		seePending = 1
	}

	items, err := database.Repo.GetDasboardSearch(userID, "", page, 100, searchQuery, companyIds, employeeIds, 1, 1, 0, seePending, "2", searchEmployee, searchCompany, searchStatus, "", searchAccMan, searchSalesStaff, "") // Implement SearchDasboard in your repository
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)

	}

	// Render the template with the fetched data
	return utils.Render(c, onboarding_view.OnboardingTBody(items))
}

// Offboarding Controllers "/panels/offboarding"
func OffboardingHandler(c echo.Context) error {
	data := utils.GetTemplateData(map[string]interface{}{
		"module": "dashboard",
		"page":   "offboarding",
	})
	dataMap, ok := data.Data.(map[string]interface{})
	if !ok {
		log4go.LOGGER("error").Error("Error parsing data")
		return utils.RenderError500(c, "Error parsing data", false)
	}
	module, _ := dataMap["module"].(string)
	page, _ := dataMap["page"].(string)
	return utils.Render(c, offboarding_view.OffboardingView(data.ShowCompanySearch, module, page))
}
func OffboardingSearchHandler(c echo.Context) error {
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	page, _ := strconv.Atoi(c.QueryParam("page"))

	searchQuery := c.QueryParam("searchInput")
	searchEmployee := c.QueryParam("searchEmployee")
	searchCompany := c.QueryParam("searchCompany")
	searchStatus := c.QueryParam("searchStatus")

	companyIds := c.QueryParam("companyIds")
	employeeIds := c.QueryParam("employeeIds")

	seePending := 1
	if c.QueryParam("seePending") == "true" {
		seePending = 1
	}
	items, err := database.Repo.GetDasboardSearch(userID, "", page, 100, searchQuery, companyIds, employeeIds, 1, 1, 1, seePending, "3", searchEmployee, searchCompany, searchStatus, "", "", "", "") // Implement SearchDasboard in your repository
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)

	}

	// Render the template with the fetched data
	return utils.Render(c, offboarding_view.OffboardingTBody(items))
}

// Missing Controllers "/panels/missing"
func MissingHandler(c echo.Context) error {
	data := utils.GetTemplateData(map[string]interface{}{
		"module": "dashboard",
		"page":   "missing",
	})
	dataMap, ok := data.Data.(map[string]interface{})
	if !ok {
		log4go.LOGGER("error").Error("Error parsing data")
		return utils.RenderError500(c, "Error parsing data", false)
	}
	module, _ := dataMap["module"].(string)
	page, _ := dataMap["page"].(string)
	return utils.Render(c, missing_view.MissingInfoView(data.ShowCompanySearch, module, page))
}
func MissingSearchHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("missing search!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")

	page, _ := strconv.Atoi(c.QueryParam("page"))

	searchQuery := c.QueryParam("searchInput")
	searchEmployee := c.QueryParam("searchEmployee")
	searchCompany := c.QueryParam("searchCompany")
	searchEmail := c.QueryParam("searchEmail")

	companyIds := c.QueryParam("companyIds")
	employeeIds := c.QueryParam("employeeIds")

	seePending := 1
	if c.QueryParam("seePending") == "true" {
		seePending = 1
	}

	items, err := database.Repo.GetDasboardSearch(userID, "", page, 100, searchQuery, companyIds, employeeIds, 1, 0, 0, seePending, "4", searchEmployee, searchCompany, "", "", "", "", searchEmail) // Implement SearchDasboard in your repository
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// Render the template with the fetched data
	return utils.Render(c, missing_view.MissingInfoTBody("", items))
}
func MissingGetByIdHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("missing get!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	id := c.QueryParam("id")
	cid := c.QueryParam("companyId")
	//cuid := r.Header.Get("Companyuserid")

	// get edit from header
	edit := c.QueryParam("edit")
	// Fetch the current state of the item using the id
	// Return the fetched data or handle errors

	items, err := database.Repo.GetDasboardSearch(userID, id, 1, 1, "", cid, id, -1, -1, -1, 1, "4", "", "", "", "", "", "", "") // Implement SearchDasboard in your repository
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// Data to be passed to the template

	// Render the template with the fetched data

	return utils.Render(c, missing_view.MissingInfoTBody(edit, items))
}
func MissingUpdateHandler(c echo.Context) error {
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	id := c.QueryParam("id")
	cid := c.QueryParam("companyId")
	cuid := c.QueryParam("companyUserId")

	// Extracting the variables from the form data
	eid := c.Request().FormValue("eid")
	passport := c.Request().FormValue("passport")
	email := c.Request().FormValue("email")
	phone := c.Request().FormValue("phone")
	ejari := c.Request().FormValue("ejari")
	ejari_expiration := c.Request().FormValue("ejari_expiration")
	laborCard := c.Request().FormValue("laborCard")
	visaNumber := c.Request().FormValue("visaNumber")
	dewa := c.Request().FormValue("dewa")
	dewa_expiration := c.Request().FormValue("dewa_expiration")

	log4go.LOGGER("info").Info("Update Data:: eid=%v,passport=%v,email=%v,phone=%v,ejari=%v,ejari_expiration=%v,laborCard=%v,visaNumber=%v,dewa=%v,dewa_expiration=%v", eid, passport, email, phone, ejari, ejari_expiration, laborCard, visaNumber, dewa, dewa_expiration)

	// Call CrudDasboards function with extracted variables
	_, err := database.Repo.CrudDasboards(
		userID,           // User ID, assuming it's defined elsewhere
		id,               // ID, assuming it's defined elsewhere
		cid,              // Company ID, assuming it's not provided in form
		cuid,             // Company user ID, assuming it's not provided in form
		eid,              // Emirates ID
		passport,         // Passport
		email,            // Email
		phone,            // Phone
		ejari,            // Ejari Number
		ejari_expiration, // Ejari Expiration
		laborCard,        // Labor Card ID
		"",               // Labor Card Expiration
		visaNumber,       // Visa Number
		"",               // Visa Expiration
		"",               // Labour ID, assuming it's not provided in form
		"",               // Card Number, assuming it's not provided in form
		"",               // Insurances Start Date, assuming it's not provided in form
		"",               // Insurances End Date, assuming it's not provided in form
		"",               // Business Category ID, assuming it's not provided in form
		dewa,             // Dewa Number
		dewa_expiration,  // Dewa Expiration
	)
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())

		return utils.RenderError500(c, err.Error(), false)
	}

	items, err := database.Repo.GetDasboardSearch(userID, id, 1, 1, "", cid, id, -1, -1, -1, 1, "4", "", "", "", "", "", "", "") // Implement SearchDasboard in your repository
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())

		return utils.RenderError500(c, err.Error(), false)
	}

	// Data to be passed to the template

	// Render the template with the fetched data

	return utils.Render(c, missing_view.MissingInfoTBody("", items))

}

func CompaniesSearchHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("search!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	Module := c.QueryParam("module")
	Page := c.QueryParam("page")

	//values := c.Request().Header.Get("values")
	values := c.QueryParam("values")
	// Retrieve the value of the chargeID parameter from the URL
	route := c.Param("route")

	log4go.LOGGER("info").Info(route)

	searchQuery := c.QueryParam("search")

	log4go.LOGGER("info").Info(searchQuery)

	var items []models.Company // Declare items variable in the outer scope

	//uniqueString := "companies_" + userID                             // Concatenate userID with a prefix
	items, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", "", searchQuery) // Assign the value to items in the outer scope
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// select the 25 first items
	var filtered []models.Company

	for _, company := range items {
		company.Action = "/panels/" + Module + "/" + Page + "/" + strconv.Itoa(company.Id)
		if len(filtered) >= 25 {
			break
		}
		filtered = append(filtered, company)
	}

	// Render the template with the fetched data
	return utils.Render(c, components.CompanyLinks(values, filtered))
}

func EmployeesSearchHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("search!")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	fast := c.QueryParam("fast")
	//values := c.Request().Header.Get("values")
	values := c.QueryParam("values")
	companyIds := strings.Split(c.Request().Header.Get("companyids"), ",")
	//Module := r.Header.Get("Module")
	//Page := r.Header.Get("Page")
	// Retrieve the value of the chargeID parameter from the URL
	route := c.Param("route")
	log4go.LOGGER("info").Info(route)

	searchQuery := c.QueryParam("search")
	log4go.LOGGER("info").Info(searchQuery)

	var items []models.Employee // Declare items variable in the outer scope
	var err error

	if fast == "true" {
		log4go.LOGGER("info").Info("employee fast query")
		if companyIds[0] != "" {
			var eachEmpData []models.Employee
			//Loop through the companyIds and get the employees
			//TODO: GetEmployeesSearchFast Not returning employee UUIDs
			//Changed to GetEmployeesSearch
			for i := range companyIds {
				eachEmpData, err = database.Repo.GetEmployeesSearchFast(userID, 0, 0, "", companyIds[i], searchQuery) // Assign the value to items in the outer scope
				if err != nil {
					log4go.LOGGER("error").Error(err.Error())
					return utils.RenderError500(c, err.Error(), false)
				}
				items = append(items, eachEmpData...)
			}
		} else {
			items, err = database.Repo.GetEmployeesSearchFast(userID, 0, 0, "", "", searchQuery) // Assign the value to items in the outer scope
			if err != nil {
				log4go.LOGGER("error").Error(err.Error())
				return utils.RenderError500(c, err.Error(), false)
			}
		}
	} else {
		if companyIds[0] != "" {
			var eachEmpData []models.Employee
			//Loop through the companyIds and get the employees
			for i := range companyIds {
				eachEmpData, err = database.Repo.GetEmployeesSearch(userID, 0, 0, "", companyIds[i], searchQuery) // Assign the value to items in the outer scope
				if err != nil {
					log4go.LOGGER("error").Error(err.Error())
					return utils.RenderError500(c, err.Error(), false)
				}
				items = append(items, eachEmpData...)
			}
		} else {
			items, err = database.Repo.GetEmployeesSearch(userID, 0, 0, "", "", searchQuery) // Assign the value to items in the outer scope
			if err != nil {
				log4go.LOGGER("error").Error(err.Error())
				return utils.RenderError500(c, err.Error(), false)
			}
		}
	}

	// select the 25 fiart items
	var filtered []models.Employee

	for _, company := range items {
		if len(filtered) >= 25 {
			break
		}
		filtered = append(filtered, company)
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	// Render the template with the fetched data
	return utils.Render(c, components.EmployeeLinks(values, filtered))
}
