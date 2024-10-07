package controllers

import (
	"fmt"
	"kub/dashboardES/internal/database"
	"kub/dashboardES/internal/models"
	charge_view "kub/dashboardES/internal/templates/invoicing/charge"
	invoice_view "kub/dashboardES/internal/templates/invoicing/invoice"
	"kub/dashboardES/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

func ChargesHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("charges id !")
	// Fetch user-id from the headers
	// userID := r.Header.Get("user-id")

	module := "invoicing"
	page := "charges"

	// Data to be passed to the template
	data := charge_view.ChargedefsViewData{
		Module: module,
		Page:   page,
		Route:  "/panels/invoicing/charges",
	}

	// Render the template with the fetched data
	return utils.Render(c, charge_view.ChargedefsView(data))
}

func ChargeDefSearchHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("chargedefs search !")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")

	// Retrieve the value of the chargeID parameter from the URL

	typeParam := c.QueryParam("type")

	module := "invoicing"
	page := c.QueryParam("page")

	companyIds := c.QueryParam("companyIds")
	employeeIds := c.QueryParam("employeeIds")

	var selectedCompanies []models.Company
	var chargedefsCompany []models.ChargeDef
	var chargedefsEmp []models.ChargeDef
	var Employees []models.Employee
	var chargedefsEmployees []models.ChargedefsByEmployee
	var chargedefsCompanies []models.ChargedefsByCompany

	//uniqueString := "companies_" + userID                             // Concatenate userID with a prefix
	selectedCompanies, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyIds, "") // Assign the value to items in the outer scope
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	if typeParam == "company" {

		chargedefsCompany, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, "", companyIds, false, "", "")
		if err != nil {
			// Handle error
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		data := charge_view.ChargedefsTBodyData{
			Module:      module,
			Page:        page,
			CompanyIds:  companyIds,
			CompanyData: selectedCompanies,
			Chargedefs:  chargedefsCompany,
			Route:       "/panels/invoicing/charges",
		}
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return utils.Render(c, charge_view.ChargedefsTBody(data))
	}

	if typeParam == "employee" {
		// Fetching charges for companies
		chargedefsCompany, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, "", companyIds, false, "", "")
		if err != nil {
			// Handle error
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		// Map for quick lookup of company charges by CompanyId
		tempCompanyChargeMap := make(map[int][]models.ChargeDef)
		for _, charge := range chargedefsCompany {
			refID, err := strconv.Atoi(charge.ReferenceId)
			if err != nil {
				// Handle error
				continue
			}
			tempCompanyChargeMap[refID] = append(tempCompanyChargeMap[refID], charge)
		}

		// Fetching employees
		Employees, err = database.Repo.GetEmployeesSearch(userID, 0, 0, employeeIds, companyIds, "")
		if err != nil {
			// Handle error
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		// Convert the list of employees into a string of IDs separated by commas
		var employeeIDs []string
		for _, employee := range Employees {
			employeeIDs = append(employeeIDs, strconv.Itoa(employee.Id))
		}
		if employeeIds == "" {
			employeeIds = strings.Join(employeeIDs, ",")
		}

		// Fetching charges for employees
		chargedefsEmp, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, "", companyIds, true, "", employeeIds)
		if err != nil {
			// Handle error
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		// Map for quick lookup of charges by ReferenceId
		tempChargeMap := make(map[int][]models.ChargeDef)
		for _, charge := range chargedefsEmp {
			refID, err := strconv.Atoi(charge.ReferenceId)
			if err != nil {
				// Handle error
				continue
			}
			tempChargeMap[refID] = append(tempChargeMap[refID], charge)
		}

		// Populate the slice with Employee data and corresponding Charges

		for _, emp := range Employees {
			ec := models.ChargedefsByEmployee{
				Employee: emp,
				Charges:  tempChargeMap[emp.Id],
			}
			chargedefsEmployees = append(chargedefsEmployees, ec)
		}

		for _, comp := range selectedCompanies {

			var chargedefsEmployeesComp []models.ChargedefsByEmployee
			for _, emp := range Employees {
				companyId, err := strconv.Atoi(emp.CompanyId)
				if err != nil {
					// Handle error
					continue
				}
				if companyId == comp.Id {
					ec := models.ChargedefsByEmployee{
						Employee: emp,
						Charges:  tempChargeMap[emp.Id],
					}
					chargedefsEmployeesComp = append(chargedefsEmployeesComp, ec)
				}
			}

			cc := models.ChargedefsByCompany{
				Company:   comp,
				Charges:   tempCompanyChargeMap[comp.Id],
				Employees: chargedefsEmployeesComp,
			}
			chargedefsCompanies = append(chargedefsCompanies, cc)
		}

		data := charge_view.ChargedefsTBodyStructuredData{
			Module:              module,
			Page:                page,
			CompanyIds:          companyIds,
			ChargedefsCompany:   chargedefsCompany,
			ChargedefsCompanies: chargedefsCompanies,
			Route:               "/panels/invoicing/charges",
			ChargedefsEmployees: chargedefsEmployees,
		}
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return utils.Render(c, charge_view.ChargedefsTBodyStructured(data))
	}

	log4go.LOGGER("error").Error("Invalid type parameter")
	return utils.RenderError500(c, "Invalid type parameter", false)
}

func ChargesByCompanyAndIdHandler(c echo.Context) error {

	log4go.LOGGER("info").Info("charges id !")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	// Retrieve the value of the chargeID parameter from the URL
	companyId := c.Param("companyId")
	chargedefId := c.Param("chargedefId")

	chargeType := c.QueryParam("reftype")
	refname := c.QueryParam("name")
	if refname == "" {
		refname = "selection"
	}
	refid := c.QueryParam("refid")
	refuuid := c.QueryParam("refuuid")

	var data charge_view.ChargeDefsFormData

	var selectedCompany []models.Company
	var chargedefItem []models.ChargeDef
	var chargeTypes []models.ChargeType
	var err error
	if companyId == "_" {
		companyId = ""
	}

	if companyId != "" {
		selectedCompany, err = database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyId, "") // Assign the value to items in the outer scope
		if err != nil {
			// Handle error
			log4go.LOGGER("error").Debug(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}
	} else {
		companyId = "_"
	}
	if chargedefId != "new" {
		chargedefItem, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, chargedefId, "", false, "", "") // not optimized, if company id was there would be faster but it does not wor for employees
		if err != nil {
			// Handle error
			log4go.LOGGER("error").Debug(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}
		if len(chargedefItem) == 0 {
			// Handle error when chargedefItem is empty
			log4go.LOGGER("error").Debug(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}
		if chargeType == "user" {
			data.Mode = "Edit user charge definition: " + chargedefItem[0].Name
		} else {
			data.Mode = "Edit company charge definition: " + chargedefItem[0].Name
		}
	} else {
		if chargeType == "user" {
			data.Mode = "Add new " + chargeType + " charge definition to " + refname
		} else {
			data.Mode = "Add new " + chargeType + " charge definition to " + refname
		}

	}

	chargeTypes, err = database.Repo.GetChargeTypes(userID)
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Debug(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	data.Types = append(data.Types, chargeTypes...)

	// Data to be passed to the template

	data.Company = selectedCompany
	data.CompanyId = companyId
	data.ChargedefId = chargedefId
	data.Types = chargeTypes
	data.Reftype = chargeType

	if chargeType == "company" {
		data.IsUser = "0"
		data.Refuuid = refuuid
		data.Refid = refid
	} else {
		data.IsUser = "1"
		data.Refuuid = refuuid
		data.Refid = refid
	}

	if len(chargedefItem) != 0 {
		data.InitData = chargedefItem[0]
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	// Render the template with the fetched data
	return utils.Render(c, charge_view.ChargedefsForm(data))
}

func ChargeCreateHandler(c echo.Context) error {
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	refuuid := c.QueryParam("ref-uuid")
	refid := c.QueryParam("ref-id")
	reftype := c.QueryParam("ref-type")

	refUuids := strings.Split(c.Request().Header.Get("refuuids"), ",")
	refIds := strings.Split(c.Request().Header.Get("refids"), ",")
	//Safe check for reftype
	refType := c.Request().Header.Get("reftype")
	if reftype == "" {
		reftype = refType
	}

	chargedefId := c.Param("chargedefId")

	const (
		refTypeUser    = "App////Contacts////Models////User"
		refTypeCompany = "App////Contacts////Models////Company"
	)

	if reftype == "company" {
		reftype = refTypeCompany
	} else {
		reftype = refTypeUser
	}

	if chargedefId == "new" {
		chargedefId = ""
	}

	if refUuids[0] == "" && refuuid == "" {
		log4go.LOGGER("error").Error("Server Error: Ref UUID is empty or invalid")
		return c.String(500, "Server Error: Ref UUID is empty or invalid")
	}

	totalAmount := c.FormValue("totalAmount")
	monthlyAmount := c.FormValue("monthlyAmount")
	months := c.FormValue("months")
	prepayAmount := c.FormValue("prepayAmount")
	monthsLack := c.FormValue("monthsLack")
	autoRenew := c.FormValue("autoRenew")

	if autoRenew == "on" {
		autoRenew = "1"
	} else {
		autoRenew = "0"
	}

	name := c.FormValue("name")
	typeParam := c.FormValue("type")
	concept := c.FormValue("concept")
	contractStart := c.FormValue("contractStart")
	contractEnd := c.FormValue("contractEnd")

	var result string
	var err error

	if refuuid == "" {
		if len(refIds) != len(refUuids) {
			log4go.LOGGER("error").Error("Server Error: Ref IDs and Ref UUIDs are not equal")
			return c.String(500, "Server Error: Ref IDs and Ref UUIDs are not equal")
		}
		// Store multiple charges for each company
		for i := range refIds {
			result, err = database.Repo.CrudChargeDef(userID, chargedefId, reftype, refUuids[i], refIds[i], contractStart, contractEnd, totalAmount, monthlyAmount, months, prepayAmount, autoRenew, monthsLack, "", typeParam, "", name, concept)
			if err != nil {
				log4go.LOGGER("error").Error(err.Error())
				return c.String(500, err.Error())
			}
		}
	} else {
		result, err = database.Repo.CrudChargeDef(userID, chargedefId, reftype, refuuid, refid, contractStart, contractEnd, totalAmount, monthlyAmount, months, prepayAmount, autoRenew, monthsLack, "", typeParam, "", name, concept)
		if err != nil {
			log4go.LOGGER("error").Error(err.Error())
			return c.String(500, err.Error())
		}
	}

	return c.String(200, result)
}

func ChargeDeleteHandler(c echo.Context) error {

	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	//refid := c.QueryParam("ref-id")
	//refuuid := c.QueryParam("ref-uuid")
	reftype := c.QueryParam("ref-type")

	// reftype := c.QueryParam("reftype")
	// refuuid := c.QueryParam("refuuid")
	// Retrieve the value of the chargeID parameter from the URL
	//companyId := c.Param("companyId")
	chargedefId := c.Param("chargedefId")

	//

	const (
		refTypeUser    = "App\\Contacts\\Models\\User"
		refTypeCompany = "App\\Contacts\\Models\\Company"
	)

	if reftype == "company" {
		reftype = refTypeCompany

	} else {
		reftype = refTypeUser
	}

	_, err := database.Repo.DeleteChargeDef(userID, chargedefId) // Implement CrudChargeDef in your repository

	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return c.String(500, err.Error())
	}

	return c.String(200, "Delete Successful!")
}

func GetChargesHandler(c echo.Context) error {
	companyId := c.QueryParam("companyId")
	panelId := c.QueryParam("panelId")
	chargeId := c.QueryParam("chargeId")
	userID := c.Request().Header.Get("user-id")

	charges, err := database.Repo.GetChargesSearch(userID, 0, 0, chargeId, companyId)

	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, charge_view.SingChargeTable(charges, panelId))
}

// /NEW Filter controllers
func FilterCompanyChargesHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("Company charges filter")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	// Retrieve the value of the chargeID parameter from the URL
	filter := c.QueryParam("filter") //User,Company or All
	module := "invoicing"
	page := c.QueryParam("page")

	companyIds := c.QueryParam("companyIds")

	var selectedCompanies []models.Company
	var chargedefsCompany []models.ChargeDef

	//uniqueString := "companies_" + userID                             // Concatenate userID with a prefix
	selectedCompanies, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyIds, "") // Assign the value to items in the outer scope
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	chargedefsCompany, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, "", companyIds, false, "", "")

	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	if filter != "All" {
		// filter chargedefsCompany by ReferenceType
		chargedefsCompany = utils.FilterChargeDefsByReferenceType(chargedefsCompany, filter)
	}

	data := charge_view.ChargedefsTBodyData{
		Module:      module,
		Page:        page,
		CompanyIds:  companyIds,
		CompanyData: selectedCompanies,
		Chargedefs:  chargedefsCompany,
		Route:       "/panels/invoicing/charges",
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, charge_view.ChargedefsTBody(data))

}

func FilterEmployeeChargesHandler(c echo.Context) error {
	log4go.LOGGER("info").Info("chargedefs search !")
	// Fetch user-id from the headers
	userID := c.Request().Header.Get("user-id")
	filter := c.QueryParam("filter") //User,Company or All
	// Retrieve the value of the chargeID parameter from the URL

	module := "invoicing"
	page := c.QueryParam("page")

	companyIds := c.QueryParam("companyIds")
	employeeIds := c.QueryParam("employeeIds")

	var selectedCompanies []models.Company
	var chargedefsCompany []models.ChargeDef
	var chargedefsEmp []models.ChargeDef
	var Employees []models.Employee
	var chargedefsEmployees []models.ChargedefsByEmployee
	var chargedefsCompanies []models.ChargedefsByCompany

	//uniqueString := "companies_" + userID                             // Concatenate userID with a prefix
	selectedCompanies, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyIds, "") // Assign the value to items in the outer scope
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// Fetching charges for companies
	chargedefsCompany, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, "", companyIds, false, "", "")
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	if filter != "All" {
		// filter chargedefsCompany by ReferenceType
		chargedefsCompany = utils.FilterChargeDefsByReferenceType(chargedefsCompany, filter)
	}

	// Map for quick lookup of company charges by CompanyId
	tempCompanyChargeMap := make(map[int][]models.ChargeDef)
	for _, charge := range chargedefsCompany {
		refID, err := strconv.Atoi(charge.ReferenceId)
		if err != nil {
			// Handle error
			continue
		}
		tempCompanyChargeMap[refID] = append(tempCompanyChargeMap[refID], charge)
	}

	// Fetching employees
	Employees, err = database.Repo.GetEmployeesSearch(userID, 0, 0, employeeIds, companyIds, "")
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	// Convert the list of employees into a string of IDs separated by commas
	var employeeIDs []string
	for _, employee := range Employees {
		employeeIDs = append(employeeIDs, strconv.Itoa(employee.Id))
	}
	if employeeIds == "" {
		employeeIds = strings.Join(employeeIDs, ",")
	}

	// Fetching charges for employees
	chargedefsEmp, err = database.Repo.GetChargeDefsSearch(userID, 0, 0, "", companyIds, true, "", employeeIds)
	if err != nil {
		// Handle error
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	if filter != "All" {
		// filter chargedefsEmp by ReferenceType
		chargedefsEmp = utils.FilterChargeDefsByReferenceType(chargedefsEmp, filter)
	}
	// Map for quick lookup of charges by ReferenceId
	tempChargeMap := make(map[int][]models.ChargeDef)
	for _, charge := range chargedefsEmp {
		refID, err := strconv.Atoi(charge.ReferenceId)
		if err != nil {
			// Handle error
			continue
		}
		tempChargeMap[refID] = append(tempChargeMap[refID], charge)
	}

	// Populate the slice with Employee data and corresponding Charges

	for _, emp := range Employees {
		ec := models.ChargedefsByEmployee{
			Employee: emp,
			Charges:  tempChargeMap[emp.Id],
		}
		chargedefsEmployees = append(chargedefsEmployees, ec)
	}

	for _, comp := range selectedCompanies {

		var chargedefsEmployeesComp []models.ChargedefsByEmployee
		for _, emp := range Employees {
			companyId, err := strconv.Atoi(emp.CompanyId)
			if err != nil {
				// Handle error
				continue
			}
			if companyId == comp.Id {
				ec := models.ChargedefsByEmployee{
					Employee: emp,
					Charges:  tempChargeMap[emp.Id],
				}
				chargedefsEmployeesComp = append(chargedefsEmployeesComp, ec)
			}
		}

		cc := models.ChargedefsByCompany{
			Company:   comp,
			Charges:   tempCompanyChargeMap[comp.Id],
			Employees: chargedefsEmployeesComp,
		}
		chargedefsCompanies = append(chargedefsCompanies, cc)
	}

	data := charge_view.ChargedefsTBodyStructuredData{
		Module:              module,
		Page:                page,
		CompanyIds:          companyIds,
		ChargedefsCompany:   chargedefsCompany,
		ChargedefsCompanies: chargedefsCompanies,
		Route:               "/panels/invoicing/charges",
		ChargedefsEmployees: chargedefsEmployees,
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, charge_view.ChargedefsTBodyStructured(data))
}

/// Invoice Controllers

func InvoiceHandeler(c echo.Context) error {
	module := "invoicing"
	page := "invoice"
	companyIDs := c.QueryParam("companyIds")
	userID := c.Request().Header.Get("user-id")

	// search companies
	companies, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyIDs, "")
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	// Data to be passed to the template
	data := invoice_view.InvoiceViewData{
		Module:    module,
		Page:      page,
		Companies: companies,
		Route:     "/panels/invoicing/charges",
	}
	return utils.Render(c, invoice_view.InvoiceView(data))
}

func GetInternalInvoicesHandler(c echo.Context) error {
	userID := c.Request().Header.Get("user-id")
	companyID := c.QueryParam("companyIds")
	invoices, err := database.Repo.GetInternalInvoiceSearch(userID, 0, 0, "", companyID)

	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, invoice_view.CompanyInvoiceTBody(invoices))
}

func GetClientInvoiceSearch(c echo.Context) error {
	userID := c.Request().Header.Get("user-id")
	companyID := c.QueryParam("companyIds")

	companies, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyID, "")
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	// Store company data by IDs in map
	companyMap := make(map[int]models.Company)
	for _, com := range companies {
		companyMap[com.Id] = com
	}
	clientInvoices, err := database.Repo.GetClientInvoiceSearch(userID, 0, 0, "", companyID)
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	//Assoicated company details to client invoices
	for i, inv := range clientInvoices {
		if com, ok := companyMap[int(inv.CompanyID)]; ok {
			clientInvoices[i].CompanyName = com.Name
		}
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, invoice_view.ClientInvoiceTBody(clientInvoices))
}

func GetInvoiceFormHandler(c echo.Context) error {
	userID := c.Request().Header.Get("user-id")
	invoiceId := c.Param("invoiceId")
	var data invoice_view.InvoiceFormViewData
	// Creating new invoice
	if invoiceId == "new" {
		//Params
		companyID := c.Request().Header.Get("companyId")
		uuidList := c.Request().Header.Get("uuidList")
		// dateRange := c.Request().Header.Get("dateRange")
		fromDate := c.QueryParam("fromDate")
		toDate := c.QueryParam("toDate")
		fmt.Println("fromDate", fromDate)
		fmt.Println("toDate", toDate)
		dateRange := fromDate + ">" + toDate
		var companyCharges []models.InvoiceCharge
		var employeeCharges []models.InvoiceCharge
		var payrollCharges []models.InvoiceCharge
		var employeeTableIds []string

		// Get employee uuids
		emps, err := database.Repo.GetEmployeesSearch(userID, 0, 0, "", companyID, "")
		if err != nil {
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		// Get employee charges and build empProfileName
		empProfileName := make(map[string]string)
		empUuidMap := make(map[string]bool)
		empIds := make(map[string]string)

		// var empUuids []string
		for _, e := range emps {
			empProfileName[e.Uuid] = e.ProfileFul
			if empProfileName[e.Uuid] == "" {
				empProfileName[e.Uuid] = e.ProfileName
			}
			empUuidMap[e.Uuid] = true
			empIds[e.Uuid] = strconv.Itoa(e.Id)
			employeeTableIds = append(employeeTableIds, e.Uuid)
		}
		//charges
		allCharges, err := database.Repo.GetInvoiceCharges(userID, "", "", companyID, dateRange, "", "")
		if err != nil {
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}
		employeeTableInfo := make(map[string]invoice_view.EmployeeTableInfo)
		for _, inv := range allCharges {
			if inv.ReferenceUUID == uuidList && !inv.InvoiceID.Valid {
				//company charges
				companyCharges = append(companyCharges, inv)
			} else if empUuidMap[inv.ReferenceUUID] && !inv.InvoiceID.Valid {

				employeeTable, exists := employeeTableInfo[inv.ReferenceUUID]
				//per employee charges
				if !exists {
					employeeTable = invoice_view.EmployeeTableInfo{
						Name:    empProfileName[inv.ReferenceUUID],
						TableId: inv.ReferenceUUID,
						Charges: []models.InvoiceCharge{},
					}
				}
				employeeTable.Charges = append(employeeTable.Charges, inv)
				employeeTableInfo[inv.ReferenceUUID] = employeeTable
				//employee list charges
				employeeCharges = append(employeeCharges, inv)
			}
		}
		//TODO: Optimize
		for uuid, id := range empIds {
			//Payroll
			payroll, err := database.Repo.GetPayrollCharges(userID, "", "", companyID, dateRange, id, "")
			if err != nil {
				log4go.LOGGER("error").Error(err.Error())
				return utils.RenderError500(c, err.Error(), false)
			}
			for _, p := range payroll {
				if empUuidMap[uuid] && !p.InvoiceID.Valid {
					//employee list charges
					employeeTable, exists := employeeTableInfo[uuid]
					//per employee charges
					if !exists {
						employeeTable = invoice_view.EmployeeTableInfo{
							Name:           empProfileName[uuid],
							TableId:        uuid,
							PayrollCharges: []models.InvoiceCharge{},
						}
					}
					employeeTable.PayrollCharges = append(employeeTable.PayrollCharges, p)
					employeeTableInfo[uuid] = employeeTable
					payrollCharges = append(payrollCharges, p)
				}
			}
		}

		data = invoice_view.InvoiceFormViewData{
			CompanyCharges:    companyCharges,
			EmployeeTableIds:  employeeTableIds,
			EmployeeCharges:   employeeCharges,
			EmployeeNames:     empProfileName,
			EmployeeTableInfo: employeeTableInfo,
			PayrollCharges:    payrollCharges,
		}
	} else {
		// Editing invoice
		// dateRange := c.Request().Header.Get("dateRange")
		fromDate := c.QueryParam("fromDate")
		toDate := c.QueryParam("toDate")
		dateRange := fromDate + ">" + toDate
		if toDate == "" && fromDate == "" {
			dateRange = ">" + time.Now().Format("2006-01-01")
		}
		companyId := c.Request().Header.Get("companyId")
		companyUuid := c.Request().Header.Get("uuidList")

		// Get employee uuids
		emps, err := database.Repo.GetEmployeesSearch(userID, 0, 0, "", companyId, "")
		if err != nil {
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		// Get employee charges and build empProfileName
		empProfileName := make(map[string]string)
		empUuidMap := make(map[string]bool)
		empIds := make(map[string]string)
		for _, e := range emps {
			empProfileName[e.Uuid] = e.ProfileFul
			if empProfileName[e.Uuid] == "" {
				empProfileName[e.Uuid] = e.ProfileName
			}
			empIds[e.Uuid] = strconv.Itoa(e.Id)
			empUuidMap[e.Uuid] = true
		}

		// Fetch all charges
		allCharges, err := database.Repo.GetInvoiceCharges(userID, "", "", companyId, dateRange, "", "")
		if err != nil {
			log4go.LOGGER("error").Error(err.Error())
			return utils.RenderError500(c, err.Error(), false)
		}

		var (
			companySelRows  []string
			employeeSelRows []string
			selPayroll      []string
			companyCharges  []models.InvoiceCharge
			employeeCharges []models.InvoiceCharge
			payrollCharges  []models.InvoiceCharge

			employeeTableIds []string
		)

		intInvoiceId, _ := strconv.Atoi(invoiceId)

		for _, inv := range allCharges {
			if inv.ReferenceUUID == companyUuid {
				if inv.InvoiceID.Int64 == int64(intInvoiceId) {
					companySelRows = append(companySelRows, strconv.Itoa(len(companyCharges)))
					companyCharges = append(companyCharges, inv)
				} else if !inv.InvoiceID.Valid {
					companyCharges = append(companyCharges, inv)
				}
			} else if empUuidMap[inv.ReferenceUUID] {
				if inv.InvoiceID.Int64 == int64(intInvoiceId) {
					employeeSelRows = append(employeeSelRows, strconv.Itoa(len(employeeCharges)))
					employeeCharges = append(employeeCharges, inv)
				} else if !inv.InvoiceID.Valid {
					employeeCharges = append(employeeCharges, inv)
				}
			}
		}
		employeeTableInfo := make(map[string]invoice_view.EmployeeTableInfo)
		//per employee charges
		for _, inv := range allCharges {
			if empUuidMap[inv.ReferenceUUID] {
				//per employee charges
				employeeTable, exists := employeeTableInfo[inv.ReferenceUUID]
				if !exists {
					employeeTable = invoice_view.EmployeeTableInfo{
						Name:         empProfileName[inv.ReferenceUUID],
						TableId:      inv.ReferenceUUID,
						Charges:      []models.InvoiceCharge{},
						SelectedRows: []string{},
					}
				}
				if inv.InvoiceID.Int64 == int64(intInvoiceId) {
					employeeTable.SelectedRows = append(employeeTable.SelectedRows, strconv.Itoa(len(employeeTable.Charges)))
					employeeTable.Charges = append(employeeTable.Charges, inv)
					employeeTableInfo[inv.ReferenceUUID] = employeeTable
				} else if !inv.InvoiceID.Valid {
					employeeTable.Charges = append(employeeTable.Charges, inv)
					employeeTableInfo[inv.ReferenceUUID] = employeeTable
				}
			}
		}
		//TODO: Optimize
		for uuid, id := range empIds {
			//Payroll
			payroll, err := database.Repo.GetPayrollCharges(userID, "", "", companyId, dateRange, id, "")
			if err != nil {
				log4go.LOGGER("error").Error(err.Error())
				return utils.RenderError500(c, err.Error(), false)
			}
			for _, p := range payroll {
				if empUuidMap[uuid] {
					//employee list
					if p.InvoiceID.Int64 == int64(intInvoiceId) {
						selPayroll = append(selPayroll, strconv.Itoa(len(payrollCharges)))
						payrollCharges = append(payrollCharges, p)
					} else if !p.InvoiceID.Valid {
						payrollCharges = append(payrollCharges, p)
					}
					//per employee
					employeeTable, exists := employeeTableInfo[p.ReferenceUUID]
					if !exists {
						employeeTable = invoice_view.EmployeeTableInfo{
							Name:            empProfileName[p.ReferenceUUID],
							TableId:         p.ReferenceUUID,
							PayrollCharges:  []models.InvoiceCharge{},
							SelectedPayroll: []string{},
						}
					}
					if p.InvoiceID.Int64 == int64(intInvoiceId) {
						employeeTable.SelectedPayroll = append(employeeTable.SelectedPayroll, strconv.Itoa(len(employeeTable.PayrollCharges)))
						employeeTable.PayrollCharges = append(employeeTable.Charges, p)
						employeeTableInfo[p.ReferenceUUID] = employeeTable
					} else if !p.InvoiceID.Valid {
						employeeTable.PayrollCharges = append(employeeTable.Charges, p)
						employeeTableInfo[p.ReferenceUUID] = employeeTable
					}
				}

			}
		}

		for k := range employeeTableInfo {
			employeeTableIds = append(employeeTableIds, k)
		}

		data = invoice_view.InvoiceFormViewData{
			CompanyCharges:    companyCharges,
			EmployeeCharges:   employeeCharges,
			EmployeeNames:     empProfileName,
			EmployeeSelRows:   employeeSelRows,
			CompanySelRows:    companySelRows,
			EmployeeTableIds:  employeeTableIds,
			PayrollCharges:    payrollCharges,
			SelectedPayroll:   selPayroll,
			EmployeeTableInfo: employeeTableInfo,
		}

	}
	return utils.Render(c, invoice_view.InvoiceFormView(data))
}

func ManageInvoiceHandler(c echo.Context) error {
	invoiceID := c.Param("invoiceId")
	userID := c.Request().Header.Get("user-id")
	companyID := c.Param("companyId")
	dateRange := ">" + time.Now().Format("2006-01-01")

	// Get company details and employee UUIDs and names
	company, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyID, "")
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	emps, err := database.Repo.GetEmployeesSearch(userID, 0, 0, "", companyID, "")
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	employeeNames := make(map[string]string)
	for _, e := range emps {
		if e.ProfileFul != "" {
			employeeNames[e.Uuid] = e.ProfileFul
		} else {
			employeeNames[e.Uuid] = e.ProfileName
		}
	}

	// Get company and employee charges
	companyCharges, employeeCharges := []models.InvoiceCharge{}, []models.InvoiceCharge{}
	allCharges, err := database.Repo.GetInvoiceCharges(userID, "", "", companyID, dateRange, "", invoiceID)
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}

	for _, inv := range allCharges {
		if inv.ReferenceUUID == company[0].Uuid && inv.InvoiceID.Valid {
			companyCharges = append(companyCharges, inv)
		} else if inv.InvoiceID.Valid {
			employeeCharges = append(employeeCharges, inv)
		}
	}
	data := invoice_view.ManageInvoiceViewData{
		CompanyCharges:  companyCharges,
		EmployeeCharges: employeeCharges,
		Company:         company[0],
		Module:          "Edit Internal Invoice",
		Page:            fmt.Sprintf("%v - Invoice %v", company[0].Name, invoiceID),
		EmployeeNames:   employeeNames,
	}

	return utils.Render(c, invoice_view.ManageInvoiceView(data))

}

func PreviewInvoiceHandler(c echo.Context) error {
	invoiceId := c.Param("invoiceId")
	companyId := c.Param("companyId")
	userID := c.Request().Header.Get("user-id")
	// Get company details and employee UUIDs and names
	company, err := database.Repo.GetCompaniesSearch(userID, 0, 0, "", companyId, "")
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	templates := database.Repo.GetInvoiceTemplateList()
	//TODO:
	_, err = database.Repo.PostInovoiceTemplate(userID, 5, "", "2024-02-16", false)
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	data := invoice_view.InvoicePreviewViewData{
		InvoiceTemplates: templates,
		Page:             fmt.Sprintf("%v - Invoice %v", company[0].Name, invoiceId),
		Module:           "Create Client Invoice",
	}
	return utils.Render(c, invoice_view.InvoicePreviewView(data))
}

func CreateUpdateInvoiceHandler(c echo.Context) error {
	//
	userID := c.Request().Header.Get("user-id")
	invoiceId := c.Param("invoiceId")
	companyId := c.Request().Header.Get("companyId")
	month := c.Request().Header.Get("month")
	year := c.Request().Header.Get("year")
	// each charge selected uuids
	selections := c.Request().Header.Get("selections")
	if invoiceId == "new" {
		invoiceId = ""
	}
	_, err := database.Repo.CreateInvoice(userID, invoiceId, month, year, companyId, selections, "", "")
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return c.String(500, err.Error())
	}
	var msg string
	if invoiceId == "" {
		msg = "Invoice Created Successfully"
	} else {
		msg = "Invoice Updated Successfully"

	}
	return c.String(200, msg)
}

func DeleteInvoiceHandler(c echo.Context) error {
	invoiceId := c.Param("invoiceId")
	userID := c.Request().Header.Get("user-id")

	_, err := database.Repo.DeleteInvoice(userID, invoiceId)
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return c.String(200, err.Error())
	}
	return c.String(200, "Invoice deleted successfully")
}

func GetInvoiceSignHandler(c echo.Context) error {
	invoiceId := c.Param("invoiceId")
	userID := c.Request().Header.Get("user-id")
	//TODO: Incomplete
	_, err := database.Repo.GetInternalInvoiceSignatures(userID, invoiceId)
	if err != nil {
		log4go.LOGGER("error").Error(err.Error())
		return utils.RenderError500(c, err.Error(), false)
	}
	return utils.Render(c, invoice_view.InvoiceSignView())
}
