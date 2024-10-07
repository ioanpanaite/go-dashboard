package database

import (
	"database/sql"
	"fmt"
	"kub/dashboardES/internal/models"
	"kub/dashboardES/internal/utils"
	"strconv"
	"strings"

	"github.com/jeanphorn/log4go"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0 // default value in case of an error
	}
	return num
}

func nullStringToString(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	}
	return ""
}

func nullIntToInt(nullInt sql.NullInt64) int {
	if nullInt.Valid {
		return int(nullInt.Int64)
	}
	return 0
}

func nullFloatToFloat(nullInt sql.NullFloat64) float64 {
	if nullInt.Valid {
		return nullInt.Float64
	}
	return 0
}

func checkEmptyString(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func convertpageoffset(page int, len int) int {
	offset := (page - 1) * len
	return offset
}

func ParseSQLnull_onlystring(sql string, params []interface{}) string {
	// Splitting the SQL string at each '?'
	parts := strings.Split(sql, "?")

	if len(parts)-1 != len(params) {
		return "Error: The number of placeholders and parameters do not match."
	}

	var sb strings.Builder

	for i, part := range parts {
		sb.WriteString(part)
		if i < len(params) {
			param := params[i]
			if param == nil {
				sb.WriteString("null")
			} else if str, ok := param.(string); ok && str == "" {
				sb.WriteString("null")
			} else {
				sb.WriteString(fmt.Sprintf("'%v'", param))
			}
		}
	}

	return sb.String()
}

func ParseSQLnull(sql string, params []interface{}) string {
	// Splitting the SQL string at each '?'
	parts := strings.Split(sql, "?")

	if len(parts)-1 != len(params) {
		return "Error: The number of placeholders and parameters do not match."
	}

	var sb strings.Builder

	for i, part := range parts {
		sb.WriteString(part)
		if i < len(params) {
			param := params[i]
			switch v := param.(type) {
			case nil:
				sb.WriteString("null")
			case string:
				if v == "" {
					sb.WriteString("null")
				} else {
					sb.WriteString(fmt.Sprintf("'%v'", v))
				}
			case int:
				if v == -1 {
					// If the integer is -1, write null
					sb.WriteString("null")
				} else {
					// Directly format int type without quotes
					sb.WriteString(fmt.Sprintf("%v", v))
				}
			default:
				// For all other types, continue using single quotes
				sb.WriteString(fmt.Sprintf("'%v'", v))
			}
		}
	}

	return sb.String()
}

func (r *MySQLRepository) GetDasboardSearch(UserId string, Id string, page int, nrows int, query string, companies_id_list string, user_id_list string, active int, future int, terminated int, flag int, dashboard_id string, searchEmployee string, searchCompany string, searchStatus string, searchCategory string, searchAccMan string, searchSalesStaff string, searchEmail string) ([]models.DashboardRow, error) {
	var items []models.DashboardRow

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	offset := convertpageoffset(page, nrows)

	// CREATE DEFINER=`root`@`%` PROCEDURE `sp-insurance-dashboard`(
	// IN 		_login_id 				BIGINT,
	// IN 		_debug					INT,
	// INOUT 	_return					TEXT,
	// IN 		_permissions_id_list	TEXT,
	// IN 		_pag_begin				BIGINT,
	// IN 		_pag_offset				BIGINT,
	// IN 		_filter					TEXT,
	// IN 		_companies_id_list 		TEXT,
	// IN 		_users_id_list 			TEXT,
	// IN 		_type_list 				TEXT,
	// IN 		_role_list				TEXT,
	// IN 		_roles_id_list 			TEXT,
	// IN       _business_category_id_list
	// IN 		_date 					DATE,
	// IN 		_active 				BIT,
	// IN 		_terminated 			BIT,
	// IN 		_future 				BIT,
	// IN 		_payrolls_template		BIT,
	// IN 		_dashboard				INT,
	// IN 		_filter_employee		TEXT,
	// IN 		_filter_company			TEXT,
	// IN 		_filter_status			TEXT,
	// IN 		_filter_business_category		TEXT,
	// IN 		_filter_account_manager	TEXT,
	// IN 		_filter_sales_staff		TEXT,
	// IN 		_filter_email			TEXT
	// )
	// Call the stored procedure

	procedure := "`sp-dashboard`"
	parameters := []interface{}{
		UserId,
		offset,
		nrows,
		query,
		companies_id_list,
		user_id_list, // Id
		active,
		future,
		terminated,
		dashboard_id,
		//flag,
		searchEmployee,
		searchCompany,
		searchStatus,
		searchCategory,
		searchAccMan,
		searchSalesStaff,
		searchEmail,
	}

	callStatement := "CALL " + procedure + "(?, NULL, @return_value, null, ?, ?, ?, ?, ?, 'employee', NULL, NULL, NULL, NULL, ?, ?, ?, NULL, ?, null, ?, ?, ?, ?, ?, ?, ?)"

	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)

	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.DashboardRow
		var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			clientNameNull sql.NullString

			additionDateNull       sql.NullString
			cancellationDateNull   sql.NullString
			expirationDateNull     sql.NullString
			visaExpirationDateNull sql.NullString
			statusNull             sql.NullString
			nameNull               sql.NullString
			categoryNull           sql.NullInt64
			categoryNameNull       sql.NullString
			cardNull               sql.NullString
			InsuranceTypeNull      sql.NullInt64
			InsuranceTypeNameNull  sql.NullString

			// Missing
			EIDNull             sql.NullString
			PassportNull        sql.NullString
			EmailNull           sql.NullString
			PhoneNull           sql.NullString
			EJARINull           sql.NullString
			EJARIExpirationNull sql.NullString
			DEWANull            sql.NullString
			DEWA_expirationNull sql.NullString
			VisaNumberNull      sql.NullString
			LaborCardNull       sql.NullString

			// Onboarding
			CompanyUserStartDateNull sql.NullString
			FinanceAprovedNull       sql.NullString
			AcountManagerNull        sql.NullString
			SalesStaffNull           sql.NullString

			// offboarding
			RequestDateNull     sql.NullString
			EmployeeEndDateNull sql.NullString
			SignedFormNull      sql.NullString
			FNFAmountNull       sql.NullString
			InvoiceRaisedNull   sql.NullInt64
		)
		if err := rows.Scan(
			&ignore,                      // company_user_uuid
			&item.CompanyUserId,          // company_user_id
			&item.CompanyId,              // company_id
			&clientNameNull,              // company_name
			&ignore,                      // company_uuid
			&item.Id,                     // user_id
			&ignore,                      // user_name
			&ignore,                      // user_uuid
			&ignore,                      // user_type
			&ignore,                      // date
			&ignore,                      // company_user_deleted_at
			&CompanyUserStartDateNull,    // company_user_start_date
			&ignore,                      // company_user_start_month
			&EmployeeEndDateNull,         // company_user_end_date
			&ignore,                      // company_user_end_month
			&ignore,                      // user_created_at
			&ignore,                      // user_created_month
			&ignore,                      // user_created_at_year
			&ignore,                      // user_created_at_month
			&ignore,                      // user_active
			&ignore,                      // user_terminated
			&ignore,                      // user_future
			&ignore,                      // user_created
			&ignore,                      // user_onboarded
			&statusNull,                  // user_status
			&ignore,                      // iban
			&ignore,                      // routing_code
			&LaborCardNull,               // labour_id
			&ignore,                      // working_permit_id
			&ignore,                      // profile_name
			&ignore,                      // profile_last_name
			&nameNull,                    // profile_full_name
			&PhoneNull,                   // profile_phone
			&EmailNull,                   // profile_email
			&EIDNull,                     // emirates_id
			&PassportNull,                // passport_id
			&EJARINull,                   // ejari_number
			&EJARIExpirationNull,         // ejari_expiration_date
			&ignore,                      // payroll_template_id
			&ignore,                      // payroll_template_basic_salary
			&ignore,                      // payroll_template_housing_allowance
			&ignore,                      // payroll_template_transportation_allowance
			&ignore,                      // payroll_template_mobile_allowance
			&ignore,                      // payroll_template_other_allowances
			&ignore,                      // payroll_template_total
			&ignore,                      // is_sif_exportable
			&ignore,                      // payroll_payment_methods_name
			&ignore,                      // payroll_payment_methods_type
			&ignore,                      // payroll_payment_methods_id
			&ignore,                      // payroll_payment_day
			&visaExpirationDateNull,      // visa_expiration
			&categoryNull,                // business_category_id
			&categoryNameNull,            // business_category
			&ignore,                      // user_visa_status
			&ignore,                      // user_visa_status_at
			&ignore,                      // visa_validity_start_date
			&ignore,                      // visa_validity_duration
			&ignore,                      // visa_validity_end_date
			&ignore,                      // date_of_entry
			&ignore,                      // visa_duration
			&ignore,                      // date_of_exit
			&VisaNumberNull,              // visa_number
			&InsuranceTypeNull,           // insurance_types_id
			&InsuranceTypeNameNull,       // insurance_types_name
			&cardNull,                    // employee_insurances_card_number
			&additionDateNull,            // employee_insurances_start_date
			&expirationDateNull,          // employee_insurances_expiry_date
			&ignore,                      // gratuity
			&ignore,                      // leave_available_days
			&ignore,                      // company_user_fee_amount
			&ignore,                      // health_card_number
			&ignore,                      // health_card_expiration_date
			&ignore,                      // disabled_at
			&DEWANull,                    // dewa_number
			&DEWA_expirationNull,         // dewa_expiration
			&ignore,                      // missing_user_bank_information
			&item.Flag_emiratesID,        // missing_emirates_id
			&item.Flag_emiratesIDNumber,  // wrong_emirates_id
			&item.Flag_passaportID,       // missing_passport_id
			&item.Flag_passaportIDNumber, // wrong_passport_id
			&ignore,                      // missing_insurance_card_number
			&ignore,                      // missing_insurance_types
			&ignore,                      // missing_insurances_start_date
			&ignore,                      // expiration_insurances_start_date
			&ignore,                      // missing_insurances_expiry_date
			&item.Flag_dateendexpired,    // expiration_insurances_expiry_date
			&ignore,                      // missing_visa_expiration_date
			&item.Flag_datevisaexpired,   // expiration_visa_expiration_date
			&item.Flag_visaNumber,        // missing_visa_number
			&item.Flag_ejariNumber,       // missing_ejari_number
			&item.Flag_ejariExpiration,   // missing_ejari_expiration_date
			&item.Flag_ejariDate,         // expiration_ejari_date
			&ignore,                      // missing_dewa_number
			&ignore,                      // missing_dewa_expiration_date
			&ignore,                      // expiration_dewa_date
			&ignore,                      // cancel_insurance
			&item.Pending,                // pending

		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Name = nullStringToString(nameNull)
		item.ClientName = nullStringToString(clientNameNull)

		item.Status = nullStringToString(statusNull)
		// Insurance
		item.CardNumber = nullStringToString(cardNull)
		item.AdditionDate = nullStringToString(additionDateNull)
		item.CancellationDate = nullStringToString(cancellationDateNull)
		item.ExpirationDate = nullStringToString(expirationDateNull)
		item.VisaExpirationDate = nullStringToString(visaExpirationDateNull)
		//
		item.Category = nullIntToInt(categoryNull)
		item.CategoryName = nullStringToString(categoryNameNull)
		item.InsuranceType = nullIntToInt(InsuranceTypeNull)
		item.InsuranceTypeName = nullStringToString(InsuranceTypeNameNull)
		// Missing
		item.EID = nullStringToString(EIDNull)
		item.Passport = nullStringToString(PassportNull)
		item.Email = nullStringToString(EmailNull)
		item.Phone = nullStringToString(PhoneNull)
		item.EJARI = nullStringToString(EJARINull)
		item.EJARI_expiration = nullStringToString(EJARIExpirationNull)
		item.DEWA = nullStringToString(DEWANull)
		item.DEWA_expiration = nullStringToString(DEWA_expirationNull)
		item.VisaNumber = nullStringToString(VisaNumberNull)
		item.LaborCard = nullStringToString(LaborCardNull)

		// Onboarding
		item.CompanyUserStartDate = nullStringToString(CompanyUserStartDateNull)
		item.EmployeeEndDate = nullStringToString(EmployeeEndDateNull)

		item.FinanceApproved = nullStringToString(FinanceAprovedNull)
		item.AccountManager = nullStringToString(AcountManagerNull)
		item.SalesStaff = nullStringToString(SalesStaffNull)

		// Offboarding
		item.RequestDate = nullStringToString(RequestDateNull)

		item.SignedForm = nullStringToString(SignedFormNull)
		item.FNFAmount = nullStringToString(FNFAmountNull)
		item.InvoiceRaised = nullIntToInt(InvoiceRaisedNull)

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...

	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetCompaniesSearch(UserId string, page int, nrows int, Id string, CompanyId string, filter string) ([]models.Company, error) {
	var items []models.Company

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}
	// Call the stored procedure

	const procedure = "`sp-companies-select`"
	parameters := []interface{}{
		UserId,
		"196,197", // permissions ids
		filter,
		CompanyId,
		nil,
	}
	callStatement := "CALL " + procedure + "(?,0,@return_value,?,null,null,?,?,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callStatement, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Company
		var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			nameNull sql.NullString
			UuidNull sql.NullString
		)
		if err := rows.Scan(
			&item.Id,  // companies_id
			&nameNull, // companies_name
			&ignore,   // companies_email
			&ignore,   // companies_phone
			&ignore,   // companies_legal_name
			&ignore,   // companies_legal_uid
			&ignore,   // companies_legal_address
			&UuidNull, // companies_uuid
			&ignore,   // contact_id
		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Name = nullStringToString(nameNull)
		item.Uuid = nullStringToString(UuidNull)

		item.Action = "/" + strconv.Itoa(item.Id)

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetEmployeesSearch(UserId string, page int, nrows int, Id string, CompanyId string, filter string) ([]models.Employee, error) {
	var items []models.Employee

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}
	// Call the stored procedure

	const procedure = "`sp-employees-select`"
	parameters := []interface{}{
		UserId,     //
		"196,197",  // permissions ids
		filter,     // filter
		CompanyId,  //
		Id,         //
		"employee", //type_list
		nil,        // roles_list
		nil,        // roles_id_list
		nil,        // date
		nil,        // active
		nil,        // terminated
		nil,        // future
		nil,        //
	}
	callStatement := "CALL " + procedure + "(?,0,@return_value,?,null,null,?,?,?,?,?,?,?,?,?,?,?,null,null,null,null,null,null,null)"

	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callStatement, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Employee
		var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			CompanyNameNull sql.NullString
			ProfileNameNull sql.NullString
			ProfileLastNull sql.NullString
			ProfileFulNull  sql.NullString
		)
		if err := rows.Scan(
			&item.CompanyUserUuid, // company_user_uuid
			&item.CompanyUserId,   // company_user_id
			&item.CompanyId,       // company_id
			&CompanyNameNull,      // company_name
			&ignore,               // company_uuid
			&item.Id,              // user_id
			&ignore,               // user_name
			&item.Uuid,            // user_uuid
			&ignore,               // user_type
			&ignore,               // date
			&ignore,               // company_user_deleted_at
			&ignore,               // company_user_start_date
			&ignore,               // company_user_start_month
			&ignore,               // company_user_end_date
			&ignore,               // company_user_end_month
			&ignore,               // user_created_at
			&ignore,               // user_created_month
			&ignore,               // user_created_at_year
			&ignore,               // user_created_at_month
			&ignore,               // user_active
			&item.Terminated,      // user_terminated
			&ignore,               // user_future
			&ignore,               // user_created
			&ignore,               // user_onboarded
			&ignore,               // user_state
			&ignore,               // user_status
			&ignore,               // iban
			&ignore,               // routing_code
			&ignore,               // labour_id
			&ignore,               // working_permit_id
			&ProfileNameNull,      // profile_name
			&ProfileLastNull,      // profile_last_name
			&ProfileFulNull,       // profile_full_name
			&ignore,               // profile_phone
			&ignore,               // profile_email
			&ignore,               // emirates_id
			&ignore,               // passport_id
			&ignore,               // ejari_number
			&ignore,               // ejari_expiration_date
			&ignore,               // missing_user_bank_information
			&ignore,               // payroll_template_id
			&ignore,               // payroll_template_basic_salary
			&ignore,               // payroll_template_housing_allowance
			&ignore,               // payroll_template_transportation_allowance
			&ignore,               // payroll_template_mobile_allowance
			&ignore,               // payroll_template_other_allowances
			&ignore,               // payroll_template_total
			&ignore,               // is_sif_exportable
			&ignore,               // payroll_payment_methods_name
			&ignore,               // payroll_payment_methods_type
			&ignore,               // payroll_payment_methods_id
			&ignore,               // payroll_payment_day
			&ignore,               // visa_expiration
			&ignore,               // business_category_id
			&ignore,               // business_category
			&ignore,               // user_visa_status
			&ignore,               // user_visa_status_at
			&ignore,               // visa_validity_start_date
			&ignore,               // visa_validity_duration
			&ignore,               // visa_validity_end_date
			&ignore,               // date_of_entry
			&ignore,               // visa_duration
			&ignore,               // date_of_exit
			&ignore,               // visa_number
			&ignore,               // insurance_types_id
			&ignore,               // insurance_types_name
			&ignore,               // employee_insurances_card_number
			&ignore,               // employee_insurances_start_date
			&ignore,               // employee_insurances_expiry_date
			&ignore,               // gratuity
			&ignore,               // leave_available_days
			&ignore,               // company_user_fee_amount
			&ignore,               // health_card_number
			&ignore,               // health_card_expiration_date
			&ignore,               // disabled_at
			&ignore,               // dewa_number
			&ignore,               // dewa_expiration
		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Name = nullStringToString(ProfileFulNull)
		item.CompanyName = nullStringToString(CompanyNameNull)
		item.ProfileName = nullStringToString(ProfileNameNull)
		item.ProfileLast = nullStringToString(ProfileLastNull)
		item.ProfileFul = nullStringToString(ProfileFulNull)

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetEmployeesSearchFast(UserId string, page int, nrows int, Id string, CompanyId string, filter string) ([]models.Employee, error) {
	var items []models.Employee

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}
	// Call the stored procedure

	const procedure = "`sp-companies_employees_cache-select`"
	parameters := []interface{}{
		UserId,    //
		filter,    // filter
		CompanyId, //
		"196,197", // permissions ids
	}

	callStatement := "CALL " + procedure + "(?,0,@return_value,?,?,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Employee
		//var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			ProfileFulNull      sql.NullString
			CompanyNameNull     sql.NullString
			CompanyUserUuidNull sql.NullString
			StateNull           sql.NullString
		)
		if err := rows.Scan(
			&item.Id, // company_id
			&item.Uuid,
			&ProfileFulNull,
			&item.CompanyId,
			&item.CompanyUserUuid,
			&CompanyNameNull,
			&item.CompanyUserId,
			&CompanyUserUuidNull,
			&StateNull,
		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Name = nullStringToString(ProfileFulNull)
		item.CompanyName = nullStringToString(CompanyNameNull)
		item.State = nullStringToString(StateNull)
		item.CompanyUserUuid = nullStringToString(CompanyUserUuidNull)

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// check the number of items
	if len(items) > 0 {
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func extractLastWord(input string) string {
	// Replace backslashes with forward slashes for uniformity
	uniformInput := strings.ReplaceAll(input, "\\", "/")
	// Split the string by the slash
	parts := strings.Split(uniformInput, "/")
	// Return the last part
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

func (r *MySQLRepository) GetChargesSearch(UserId string, page int, nrows int, Id string, CompanyId string) ([]models.Charge, error) {
	var items []models.Charge

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	// Call the stored procedure

	const procedure = "`sp-charges_payment-select`"
	parameters := []interface{}{
		UserId,
		nil,       // filter
		CompanyId, // company_id
		Id,        // def_id_list
		nil,       // date_range
		nil,       // ref_uuid_list
		nil,       // invoice_id_list
		nil,       // invoice_status
	}
	callStatement := "CALL " + procedure + "(?,0,@return,null,null,null,?,?,?,?,?,?,?)"

	// new query call
	rows, err := r.db.Query(callStatement, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// row interpolation
	for rows.Next() {
		var item models.Charge
		var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			UuidNull              sql.NullString
			ChargeDefIdNull       sql.NullString
			ChargeDefTypeNameNull sql.NullString
			ChargeDefTypeIdNull   sql.NullString
			DateNull              sql.NullString
			MonthNull             sql.NullString
			YearNull              sql.NullString
			PaidAtNull            sql.NullString
			PaidByNull            sql.NullString
			ReferenceIdNull       sql.NullString
			ReferenceUuidNull     sql.NullString
			ReferenceTypeNull     sql.NullString
			InvoiceIdNull         sql.NullString
		)
		if err := rows.Scan(
			&item.Id,               // id
			&UuidNull,              // uuid
			&ChargeDefIdNull,       // charges_definition_id
			&ChargeDefTypeNameNull, // charges_definition_types_name
			&ChargeDefTypeIdNull,   // charges_definition_types_description
			&YearNull,              // year
			&MonthNull,             // month
			&DateNull,              // date
			&item.Amount,           // amount
			&ignore,                // is_manual
			&PaidAtNull,            // paid_at
			&PaidByNull,            // paid_by
			&ReferenceUuidNull,     // reference_uuid
			&item.Prepaid,          // is_prepay
			&ReferenceIdNull,       // reference_id
			&ReferenceTypeNull,     // reference_type
			&InvoiceIdNull,         // internal_invoices_id
		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Uuid = nullStringToString(UuidNull)
		item.ChargeDefId = nullStringToString(ChargeDefIdNull)
		item.ChargeDefTypeName = nullStringToString(ChargeDefTypeNameNull)
		item.ChargeDefTypeId = nullStringToString(ChargeDefTypeIdNull)
		item.Date = nullStringToString(DateNull)
		item.Year = nullStringToString(YearNull)
		item.Month = nullStringToString(MonthNull)
		item.PaidAt = nullStringToString(PaidAtNull)
		item.PaidBy = nullStringToString(PaidByNull)
		item.ReferenceId = nullStringToString(ReferenceIdNull)
		item.ReferenceUuid = nullStringToString(ReferenceUuidNull)
		item.ReferenceType = extractLastWord(nullStringToString(ReferenceTypeNull))
		item.InvoiceId = nullStringToString(InvoiceIdNull)

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetChargeDefsSearch(UserId string, page int, nrows int, Id string, CompanyId string, Employees bool, Filter string, RefsId string) ([]models.ChargeDef, error) {
	var items []models.ChargeDef

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	const procedure = "`sp-charges_definition-select`"

	parameters := []interface{}{
		UserId,    // user_id
		"196,197", // permissions ids
		//checkEmptyString(Filter),  // filter
		Id,        // id_list
		CompanyId, // ref_id_list company
		//nil,                       // ref_uuid_list all
		RefsId, // ref_id_list user
	}

	if Employees {
		parameters = append(parameters, nil) // ref_id_list employees
		parameters[3] = nil                  // ref_id_list user
	} else {
		parameters = append(parameters, nil) // ref_id_list employees
	}

	// date should be appended here?

	callStatement := "CALL " + procedure + "(?,0,@return_value,?,null,null,null,?,null,?,null,?,?,null)"

	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)

	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.ChargeDef
		var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			UuidNull        sql.NullString
			NameNull        sql.NullString
			TypeIdNull      sql.NullInt64
			TypeNameNull    sql.NullString
			ConceptNull     sql.NullString
			TotalAmountNull sql.NullFloat64
			MonthlyNull     sql.NullFloat64
			MonthsNull      sql.NullInt64
			PrepayNull      sql.NullFloat64
			PrepayMonthNull sql.NullInt64
			LackNull        sql.NullString
			//RenewNull         sql.NullString
			DateStartNull     sql.NullString
			DateEndNull       sql.NullString
			ReferenceNameNull sql.NullString
			ReferenceIdNull   sql.NullString
			ReferenceUuidNull sql.NullString
			ReferenceTypeNull sql.NullString
			StatusNull        sql.NullString
			// InvoiceIdNull     sql.NullString

		)
		if err := rows.Scan(
			&item.Id,           // id
			&UuidNull,          // uuid
			&DateStartNull,     // start_date
			&DateEndNull,       // end_date
			&TotalAmountNull,   // total_amount
			&MonthlyNull,       // apportionment_amount
			&MonthsNull,        // apportionment_months
			&PrepayNull,        // prepay
			&item.Renew,        // is_renewable
			&LackNull,          // months_lack
			&StatusNull,        // status
			&TypeIdNull,        // charges_definition_types_id
			&TypeNameNull,      // charges_definition_types_name
			&ignore,            // charges_definition_types_description
			&ReferenceTypeNull, // reference_type
			&ReferenceUuidNull, // reference_uuid
			&ReferenceNameNull, // reference_name
			&PrepayMonthNull,   // prepay_month
			&ReferenceIdNull,   // reference_id
			&NameNull,          // name
			&ConceptNull,       // description
		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Uuid = nullStringToString(UuidNull)
		item.Name = nullStringToString(NameNull)
		item.TypeId = nullIntToInt(TypeIdNull)
		item.TypeName = nullStringToString(TypeNameNull)
		item.Concept = nullStringToString(ConceptNull)
		item.Amount = nullFloatToFloat(TotalAmountNull)
		item.Monthly = nullFloatToFloat(MonthlyNull)
		item.Months = nullIntToInt(MonthsNull)
		item.Prepay = nullFloatToFloat(PrepayNull)
		item.PrepayMonth = nullIntToInt(PrepayMonthNull)
		item.Lack = nullStringToString(LackNull)
		item.DateStart = nullStringToString(DateStartNull)
		item.DateEnd = nullStringToString(DateEndNull)
		item.ReferenceId = nullStringToString(ReferenceIdNull)
		item.ReferenceUuid = nullStringToString(ReferenceUuidNull)
		item.ReferenceType = extractLastWord(nullStringToString(ReferenceTypeNull))
		item.ReferenceName = nullStringToString(ReferenceNameNull)
		item.Status = nullStringToString(StatusNull)

		if item.DateEnd != "" {
			item.Period = item.DateStart + "\n" + item.DateEnd
		} else {
			item.Period = item.DateStart
		}

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func (r *MySQLRepository) CrudChargeDef(
	UserId string,
	id string, // for edit
	ref_type string, // for new, type is employee or company
	ref_uuid string,
	ref_id string,
	start_date string,
	end_date string,
	total_amount string,
	app_amount string,
	app_month string,
	prepay string,
	renew string,
	lack string,
	status string,
	chargetype string,
	prepay_month string,
	name string,
	desc string) (string, error) {

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return "error", err
	}

	const procedure = "`sp-charges_definition-crupdate`"
	parameters := []interface{}{
		UserId, // user_id
		id,
		start_date,
		end_date,
		total_amount,
		app_amount,
		app_month,
		prepay,
		renew,
		lack,
		status,
		chargetype,
		ref_type,
		ref_uuid,
		ref_id,
		prepay_month,
		name,
		desc,
	}

	callStatement := "CALL " + procedure + "(?,0,@return,?,null,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)

	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return "error", err
	}
	defer rows.Close()

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return "error", err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
		fmt.Println(returnValue.String)
	} else {
		// Handle NULL value
		// ...
	}

	return returnValue.String, rows.Err()
}

func (r *MySQLRepository) DeleteChargeDef(
	UserId string,
	id string, // for edit
) (string, error) {

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return "error", err
	}

	const procedure = "`sp-charges_definition-delete`"
	parameters := []interface{}{
		UserId, // user_id
		id,     // to remove
	}

	callStatement := "CALL " + procedure + "(?,0,@return,?,null)"

	// new query call
	rows, err := r.db.Query(callStatement, parameters...)
	if err != nil {
		return "error", err
	}
	defer rows.Close()

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return "error", err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
		fmt.Println(returnValue.String)
	} else {
		// Handle NULL value
		// ...
	}

	return "deleted", rows.Err()
}

func (r *MySQLRepository) GetBusinessCategories(UserId string) ([]models.BusinessCategory, error) {
	var items []models.BusinessCategory

	// Set @return_value
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	const procedure = "`sp-business_category-select`"
	parameters := []interface{}{
		UserId, // user_id

	}
	callStatement := "CALL " + procedure + "(?,0,@return,null,null)"

	rows, err := r.db.Query(callStatement, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BusinessCategory
		var idNull sql.NullInt64
		var nameNull sql.NullString
		var descriptionNull sql.NullString
		var hasDesignationNull sql.NullBool
		var hasPayrollNull sql.NullBool
		var isMainCompanyEnabledNull sql.NullBool

		if err := rows.Scan(
			&idNull,
			&nameNull,
			&descriptionNull,
			&hasDesignationNull,
			&hasPayrollNull,
			&isMainCompanyEnabledNull,
		); err != nil {
			return nil, err
		}

		// Convert nullable values to regular Go types
		if idNull.Valid {
			item.Id = int(idNull.Int64)
		}
		item.Name = nameNull.String
		item.Description = descriptionNull.String
		item.HasDesignation = hasDesignationNull.Bool
		item.HasPayroll = hasPayrollNull.Bool
		item.IsMainCompanyEnabled = isMainCompanyEnabledNull.Bool

		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Handle returnValue if needed
	// ...

	return items, rows.Err()

}

func (r *MySQLRepository) GetChargeTypes(UserId string) ([]models.ChargeType, error) {
	var items []models.ChargeType

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}
	// Call the stored procedure

	const procedure = "`sp-charges_definition_types-select`"
	parameters := []interface{}{
		UserId, // user_id
	}
	callStatement := "CALL " + procedure + "(?,0,@return,null,null,null)"

	// new query call
	rows, err := r.db.Query(callStatement, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.ChargeType
		var ignore sql.RawBytes // Use this for columns you want to ignore
		var (
			UuidNull         sql.NullString
			NameNull         sql.NullString
			DescriptionsNull sql.NullString
		)

		// charges_definition_types_id,charges_definition_types_uuid,charges_definition_types_name,charges_definition_types_description,charges_definition_types_isbillable,charges_definition_types_tax_percert,charges_definition_types_invoice_description,charges_definition_types_deleted_at,charges_definition_types_is_profit

		if err := rows.Scan(
			&item.Id,          // charges_definition_types_id
			&UuidNull,         // charges_definition_types_uuid
			&NameNull,         // charges_definition_types_name
			&DescriptionsNull, // charges_definition_types_description
			&ignore,           // charges_definition_types_isbillable
			&ignore,           // charges_definition_types_tax_percert
			&ignore,           // charges_definition_types_invoice_description
			&ignore,           // charges_definition_types_deleted_at
			&ignore,           // charges_definition_types_is_profit
		); err != nil {
			// Handling any error that occurs during row scanning
			return nil, err
		}

		item.Uuid = nullStringToString(UuidNull)
		item.Name = nullStringToString(NameNull)
		item.Descriptions = nullStringToString(DescriptionsNull)

		// add the item to the list of rows
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
	} else {
		// Handle NULL value
		// ...
	}

	return items, rows.Err()
}

func (r *MySQLRepository) CrudInsuranceDasboard(UserId string, Id string, dewa string, dewa_expiration string) (string, error) {

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return "error", err
	}

	const procedure = "`sp-users_additional_information-crupdate`"
	parameters := []interface{}{
		UserId, // user_id
		Id,
		dewa,
		dewa_expiration,
	}

	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?)"

	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return "error", err
	}
	defer rows.Close()

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return "error", err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
		fmt.Println(returnValue.String)
	} else {
		// Handle NULL value
		// ...
	}

	return "updated", rows.Err()
}

// CREATE DEFINER=`root`@`%` PROCEDURE `sp-dashboard-crupdate`(
// 	IN 	_login_id 							BIGINT,
// 	IN 	_debug								INT,
// 	INOUT _return							TEXT,
//     IN _user_id								BIGINT,
//     IN _company_id							BIGINT,
// 	IN _company_user_id						BIGINT,
// 	IN _emirates_id							VARCHAR(30), 	#profiles (user_id)
//     IN _passport							VARCHAR(30), 	#profiles (user_id)
// 	IN _email								VARCHAR(150), 	#profiles (user_id)
//     IN _phone								VARCHAR(255),	#profiles (user_id)
//     IN _ejari_number						VARCHAR(255),	#profiles (user_id)
// 	IN _ejari_expiration_date				DATE,			#profiles (user_id)
// 	IN _labor_card_id						VARCHAR(255),	#profiles (user_id)
// 	IN _labor_card_id_expiration_date		DATE,			#profiles (user_id)
// 	IN _visa_number							VARCHAR(255),	#employee_visas (user_id)
// 	IN _visa_expiration						DATE,			#company_user (id)
// 	IN _labour_id							VARCHAR(120),	#user_bank_information (user_id)
// 	IN _card_number							VARCHAR(255),	#employee_insurances (company_user_id)
// 	IN _employee_insurances_start_date		DATE,			#employee_insurances (company_user_id)
// 	IN _employee_insurances_end_date		DATE,			#employee_insurances (company_user_id)
//     IN _business_category_id				BIGINT,			#company_user (id),
//     IN _dewa_number							VARCHAR(45),	#users_additional_information (user_id)
//     IN _dewa_expiration						DATE			#users_additional_information (user_id)
// )

func (r *MySQLRepository) CrudDasboards(
	UserId string,
	Id string,
	CompanyId string,
	CUId string,
	emiratesID string,
	passport string,
	email string,
	phone string,
	ejariNumber string,
	ejariExpiration string,
	laborCardID string,
	laborCardExpiration string,
	visaNumber string,
	visaExpiration string,
	labourID string,
	cardNumber string,
	insurancesStartDate string,
	insurancesEndDate string,
	businessCategoryID string,
	dewaNumber string,
	dewaExpiration string) (string, error) {

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return "error", err
	}

	const procedure = "`sp-dashboard-crupdate`"
	parameters := []interface{}{
		UserId,              // _user_id
		Id,                  // _id
		CompanyId,           // _company_id
		CUId,                // _company_user_id
		emiratesID,          // _emirates_id
		passport,            // _passport
		email,               // _email
		phone,               // _phone
		ejariNumber,         // _ejari_number
		ejariExpiration,     // _ejari_expiration_date
		laborCardID,         // _labor_card_id
		laborCardExpiration, // _labor_card_id_expiration_date
		visaNumber,          // _visa_number
		visaExpiration,      // _visa_expiration
		labourID,            // _labour_id
		cardNumber,          // _card_number
		insurancesStartDate, // _employee_insurances_start_date
		insurancesEndDate,   // _employee_insurances_end_date
		businessCategoryID,  // _business_category_id
		dewaNumber,          // _dewa_number
		dewaExpiration,      // _dewa_expiration
	}

	// Construct the CALL statement with named parameters
	//1					2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	callStatement := "CALL " + procedure + "(?, 0, @return_value,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	callstring := ParseSQLnull(callStatement, parameters)

	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return "error", err
	}
	defer rows.Close()

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return "error", err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
		fmt.Println(returnValue.String)
	} else {
		// Handle NULL value
		// ...
	}

	return "updated", rows.Err()
}

///New Queries Addded

func (r *MySQLRepository) GetInternalInvoiceSearch(UserId string, page int, nrows int, Id string, CompanyId string) ([]models.InternalInvoice, error) {
	var items []models.InternalInvoice

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	// Call the stored procedure

	const procedure = "`sp-internal_invoices-select`"
	parameters := []interface{}{
		UserId,
		nil, // filter
		nil,
		CompanyId, // company_id
	}
	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?,0)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}

	// row interpolation
	defer rows.Close()

	for rows.Next() {
		var item models.InternalInvoice
		err := rows.Scan(
			&item.ID,
			&item.UUID,
			&item.Month,
			&item.Year,
			&item.Date,
			&item.CompanyID,
			&item.CompanyName,
			&item.CreatedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.CurrencyTypesID,
			&item.Amount,
			&item.CurrencyRate,
			&item.ClientInvoicesCount,
		)
		if err != nil {
			return nil, err
		}
		item.IsClientInvoices = item.ClientInvoicesCount != 0
		item.CreatedAt = strings.Split(item.CreatedAt, " ")[0]
		items = append(items, item)
	}
	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetClientInvoiceSearch(UserId string, page int, nrows int, InoviceId string, CompanyId string) ([]models.ClientInvoice, error) {
	var items []models.ClientInvoice

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	// Call the stored procedure

	const procedure = "`sp-client_invoices-select`"
	parameters := []interface{}{
		UserId,    //userId
		nil,       //filter
		InoviceId, //InvoiceId
		CompanyId, //CompanyId
		1,
	}
	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	// row interpolation
	defer rows.Close()
	for rows.Next() {
		var item models.ClientInvoice
		err := rows.Scan(
			&item.ID,
			&item.UUID,
			&item.Date,
			&item.CompanyID,
			&item.InternalInvoiceID,
			&item.CreatedBy,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.CurrencyRate,
		)

		if err != nil {
			return nil, err
		}
		item.Template = "test"
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetInvoiceCharges(UserId string, Filter string, DefIdList string,
	CompanyId string, DateRange string, RefUUIDList string, InvoiceIDList string) ([]models.InvoiceCharge, error) {
	var items []models.InvoiceCharge

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	// Call the stored procedure

	const procedure = "`sp-charges_payment-select`"
	// IN 	_login_id 							BIGINT,
	// IN 	_debug								INT,
	// INOUT _return							TEXT,
	// IN _charges_definition_id_list			TEXT,
	// IN _date_range							TEXT, 	#2023-07-23>2023-07-24
	// IN _filter								TEXT,
	// IN _companies_user_uuid_list			TEXT,
	// IN _internal_invoices_id_list			TEXT,
	// IN _notin_invoice						BIT     #el payroll no esta en un internal invoice
	parameters := []interface{}{
		UserId,
		Filter,
		CompanyId,
		DefIdList,
		DateRange,
		RefUUIDList,
		InvoiceIDList,
		nil,
	}
	callStatement := "CALL " + procedure + "(?,0,@return,null,null,null,?,?,?,?,?,?,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	// Get Columns And Types
	// fmt.Println("Column Types:")
	// types, _ := rows.ColumnTypes()
	// for _, t := range types {
	// 	fmt.Printf("%s: %s\n", t.Name(), t.DatabaseTypeName())
	// }
	// row interpolation
	defer rows.Close()
	for rows.Next() {
		var item models.InvoiceCharge
		err := rows.Scan(
			&item.ID,
			&item.UUID,
			&item.ChargesDefID,
			&item.ChargesDefTypesName,
			&item.ChargesDefTypesDesc,
			&item.Year,
			&item.Month,
			&item.Date,
			&item.Amount,
			&item.IsManual,
			&item.PaidAt,
			&item.PaidBy,
			&item.ReferenceUUID,
			&item.IsPrepay,
			&item.ReferenceID,
			&item.ReferenceType,
			&item.InvoiceID,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}

	return items, rows.Err()
}

func (r *MySQLRepository) GetPayrollCharges(UserId string, Filter string, DefIdList string,
	CompanyId string, DateRange string, EmployeeId string, InvoiceId string) ([]models.InvoiceCharge, error) {
	var items []models.InvoiceCharge

	// offset := convertpageoffset(page, len)

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	// Call the stored procedure

	const procedure = "`sp-payrolls-select`"
	// IN 	_login_id 							BIGINT,
	// IN 	_debug								INT,
	// INOUT _return							TEXT,
	// IN _payrolls_id_list					TEXT,
	// IN _payrolls_uuid_list					TEXT,
	// IN _status_list							TEXT,	#accepted
	// IN _date_range							TEXT, 	#2023-07-23>2023-07-24
	// IN _user_id_employee_list				TEXT,
	// IN _internal_invoices_id_list			TEXT,
	// IN _company_id_list						TEXT,
	// IN _filter								TEXT,
	// IN _notin_invoice						BIT     #el payroll no esta en un internal invoice
	parameters := []interface{}{
		UserId,
		nil,
		nil,
		nil,
		DateRange,
		EmployeeId,
		InvoiceId,
		CompanyId,
		Filter,
	}
	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?,?,?,?,?,?,1)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	// Get Columns And Types
	// fmt.Println("Column Types:")
	// types, _ := rows.ColumnTypes()
	// for _, t := range types {
	// 	fmt.Printf("%s: %s\n", t.Name(), t.DatabaseTypeName())
	// }
	// row interpolation
	defer rows.Close()

	for rows.Next() {
		var p models.InvoicePayroll
		err := rows.Scan(
			&p.ID,
			&p.UUID,
			&p.Year,
			&p.Month,
			&p.Status,
			&p.TotalPrice,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.InvoiceID,
		)
		if err != nil {
			return nil, err
		}
		if p.Status == "accepted" {
			item := models.InvoiceCharge{
				ID:        uint32(p.ID),
				UUID:      p.UUID,
				Date:      utils.CreateDateFromMonthAndYear(p.Month, p.Year),
				Amount:    p.TotalPrice,
				InvoiceID: p.InvoiceID,
				IsPayroll: true,
			}
			items = append(items, item)
		}
	}

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return nil, err
	}
	return items, rows.Err()
}

func (r *MySQLRepository) CreateInvoice(UserId string, invoiceId string, month string, year string,
	companyId string, uuidList string, currencyRate string, currencyTypeId string) (string, error) {

	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return "error", err
	}

	const procedure = "`sp-internal_invoices-crupdate`"
	parameters := []interface{}{
		UserId, // user_id
		invoiceId,
		month,
		year,
		companyId,
		currencyRate,
		currencyTypeId,
		nil,
		nil,
		uuidList,
	}

	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?,?,?,?,?,?,?)"

	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return "error", err
	}
	defer rows.Close()

	// Retrieve the value of the INOUT parameter
	var returnValue sql.NullString
	err = r.db.QueryRow("SELECT @return_value").Scan(&returnValue)
	if err != nil {
		return "error", err
	}

	// Check if returnValue is valid (not NULL)
	if returnValue.Valid {
		// Use returnValue.String as needed
		// ...
		fmt.Println(returnValue.String)
	} else {
		// Handle NULL value
		// ...
	}

	return "created", rows.Err()
}

func (r *MySQLRepository) DeleteInvoice(UserId string, invoiceIdList string) (string, error) {

	const procedure = "`sp-internal_invoices-delete`"
	parameters := []interface{}{
		UserId, invoiceIdList,
	}
	callStatement := "CALL " + procedure + "(?,0,@return,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	rows, err := r.db.Query(callstring)
	if err != nil {
		return "error", err
	}
	defer rows.Close()
	return "deleted", rows.Err()
}

func (r *MySQLRepository) GetInvoiceTemplateList() []models.InvoiceTemplate {
	return []models.InvoiceTemplate{
		{
			ID:           1,
			UUID:         "c991623f-e2bf-44b9-a143-6449ba327a6e",
			Type:         "all_inclusive",
			TemplateName: "sp_client_invoice_template_all_inclusive",
			Description:  "1 - All Inclusive (all months together)",
			DisplayName:  "All inclusive",
		},
		{
			ID:           2,
			UUID:         "6452712c-d20b-40c9-9a3f-044b30a6f57d",
			Type:         "all_lines",
			TemplateName: "sp_client_invoice_template_all_lines",
			Description:  "2 - Breakdown of charges and payrolls per employee",
			DisplayName:  "Breakdown of charges and payrolls per employee",
		},
		{
			ID:           3,
			UUID:         "f8b9052f-c2f3-4416-92ec-d5f44661751b",
			Type:         "group_charges",
			TemplateName: "sp_client_invoice_template_group_fees",
			Description:  "3 - Breakdown of charges per employee and Payrolls all employees together",
			DisplayName:  "Breakdown of charges per employee and Payrolls all employees together",
		},
		{
			ID:           4,
			UUID:         "9050c818-69d8-423b-9681-fc64fc95c2eb",
			Type:         "charges_payrolls_groupped",
			TemplateName: "sp_client_invoice_template_charges_payrolls_groupped",
			Description:  "4 - per month - All Charges together and all payrolls together of all employees",
			DisplayName:  "All Charges together and all payrolls together of all employees",
		},
		{
			ID:           5,
			UUID:         "07bd04c4-0556-4dd2-95d9-1c42b789d59d",
			Type:         "all_inclusive_monthly",
			TemplateName: "sp_client_invoice_template_all_inclusive_monthly",
			Description:  "5 - All Inclusive (monthly)",
			DisplayName:  "",
		},
	}
}

func (r *MySQLRepository) PostInovoiceTemplate(UserId string, template int, invoiceId string, dateTime string, commit bool) (string, error) {
	options := []string{
		"`sp-client_invoice-templates-all_inclusive`",
		"`sp-client_invoice-templates-all_lines`",
		"`sp-client_invoice-templates-group_fees`",
		"`sp-client_invoice-templates-charges_payrolls_groupped`",
		"`sp-client_invoice-templates-all_inclusive_monthly`",
		"`sp-client_invoice-templates-all_lines_by_employee`",
	}
	procedure := options[template-1]
	set := 0
	if commit {
		set = 1
	}
	//TODO:
	parameters := []interface{}{
		UserId,
		"52", //invoiceId,
		nil,
		"2024-02-16", //dateTime,
		set,
	}

	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	rows, err := r.db.Query(callstring)

	if err != nil {
		return "error", err
	}
	defer rows.Close()
	// Get Columns And Types
	// fmt.Println("Column Types:")
	// types, _ := rows.ColumnTypes()
	// for _, t := range types {
	// 	fmt.Printf("%s: %s\n", t.Name(), t.DatabaseTypeName())
	// }
	return "created", rows.Err()
}

func (r *MySQLRepository) GetInternalInvoiceSignatures(userId string, invoiceId string) ([]models.InvoiceCharge, error) {
	var items []models.InvoiceCharge
	// Declare a user-defined variable for the INOUT parameter
	_, err := r.db.Exec("SET @return_value = ''")
	if err != nil {
		return nil, err
	}

	// Call the stored procedure
	const procedure = "`sp-internal_invoices_users_signatures-select`"
	parameters := []interface{}{
		userId,
		invoiceId,
	}
	callStatement := "CALL " + procedure + "(?,0,@return,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	// new query call
	rows, err := r.db.Query(callstring)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Get Columns And Types
	fmt.Println("Column Types:")
	types, _ := rows.ColumnTypes()
	for _, t := range types {
		fmt.Printf("%s: %s\n", t.Name(), t.DatabaseTypeName())
	}
	return items, nil
}

func (r *MySQLRepository) CrudInvoiceSignature(userId string, invoiceId string) (string, error) {

	//TODO:
	procedure := "`sp-internal_invoices_users_signatures-crupdate"
	parameters := []interface{}{
		userId,
		invoiceId,
		userId,
	}

	callStatement := "CALL " + procedure + "(?,0,@return,?,?,?)"
	callstring := ParseSQLnull(callStatement, parameters)
	log4go.LOGGER("sql").Trace(callstring)
	rows, err := r.db.Query(callstring)

	if err != nil {
		return "error", err
	}
	defer rows.Close()
	// Get Columns And Types
	// fmt.Println("Column Types:")
	// types, _ := rows.ColumnTypes()
	// for _, t := range types {
	// 	fmt.Printf("%s: %s\n", t.Name(), t.DatabaseTypeName())
	// }
	return "created", rows.Err()
}
