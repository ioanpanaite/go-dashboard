package models

// DashboardRow represents a data model (modify according to your schema)
type DashboardRow struct {
	Id int

	// Common
	Name       string
	ClientName string

	RefNumber int

	Status string

	CompanyId     int
	CompanyUserId int

	Pending int

	// Insurance
	CardNumber           string
	AdditionDate         string
	ExpirationDate       string
	CancellationDate     string
	VisaExpirationDate   string
	Flag_dateendexpired  int
	Flag_datevisaexpired int
	Category             int
	CategoryName         string
	Plan                 string
	InsuranceType        int
	InsuranceTypeName    string

	// Missing
	EID                    string
	Passport               string
	Email                  string
	Phone                  string
	EJARI                  string
	EJARI_expiration       string
	DEWA                   string
	DEWA_expiration        string
	VisaNumber             string
	LaborCard              string
	Flag_emiratesID        int
	Flag_emiratesIDNumber  int
	Flag_passaportID       int
	Flag_passaportIDNumber int
	Flag_visaNumber        int
	Flag_ejariNumber       int
	Flag_ejariExpiration   int
	Flag_ejariDate         int

	// Onboarding
	CompanyUserStartDate string
	FinanceApproved      string
	AccountManager       string
	SalesStaff           string

	// Offboarding
	RequestDate     string
	EmployeeEndDate string
	SignedForm      string
	FNFAmount       string
	InvoiceRaised   int

	//
	Details int
}

type Company struct {
	// ids
	Id   int
	Uuid string
	// data
	Name   string
	Action string
	//
	Details int
}

type ChargeDef struct {
	// ids
	Id   int
	Uuid string
	//
	ReferenceId   string
	ReferenceUuid string
	ReferenceType string
	ReferenceName string
	InvoiceId     string
	// data
	Name        string
	TypeId      int
	TypeName    string
	Concept     string
	Period      string
	Amount      float64
	Monthly     float64
	Months      int
	Month       string
	Year        string
	DateStart   string
	DateEnd     string
	Prepay      float64
	PrepayMonth int
	Lack        string
	Renew       bool
	Status      string
	//
	Details int
}

type Charge struct {
	// ids
	Id   int
	Uuid string
	//
	CompanyId   string
	CompanyName string
	//
	ReferenceId   string
	ReferenceUuid string
	ReferenceType string
	InvoiceId     string
	// data
	ChargeDefId       string
	ChargeDefTypeName string
	ChargeDefTypeId   string
	Date              string
	Month             string
	Year              string
	Amount            float64
	Prepaid           bool
	Invoice           string
	PaidAt            string
	PaidBy            string
	//
	Details int
}

type ChargeType struct {
	Id   int
	Uuid string
	//
	Name         string
	Descriptions string
}
type BusinessCategory struct {
	Id   int
	Uuid string
	//
	Name                 string
	Description          string
	HasDesignation       bool
	HasPayroll           bool
	IsMainCompanyEnabled bool
}

type Employee struct {
	// ids
	Id   int
	Uuid string
	//
	CompanyId   string
	CompanyName string
	//
	CompanyUserId   int
	CompanyUserUuid string
	// data
	Name        string
	State       string
	Actibe      bool
	Terminated  bool
	Future      bool
	Onboarded   bool
	StartDate   string
	EndDate     string
	ProfileName string
	ProfileLast string
	ProfileFul  string
	//
	Details int
}

type ChargedefsByEmployee struct {
	Employee Employee
	Charges  []ChargeDef
	//
	Details int
}

type ChargedefsByCompany struct {
	Company   Company
	Charges   []ChargeDef
	Employees []ChargedefsByEmployee
	//
	Details int
}

type Repository interface {
	//
	GetDasboardSearch(UserId string, Id string, page int, nrows int, query string, companies_id_list string, user_id_list string, active int, future int, terminated int, flag int, dashboard_id string, searchEmployee string, searchCompany string, searchStatus string, searchCategory string, searchAccMan string, searchSalesStaff string, searchEmail string) ([]DashboardRow, error)
	//
	GetChargesSearch(UserId string, page int, nrows int, Id string, CompanyId string) ([]Charge, error)
	//
	GetChargeDefsSearch(UserId string, page int, nrows int, Id string, CompanyId string, Employees bool, Filter string, RefsId string) ([]ChargeDef, error)
	//
	CrudChargeDef(
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
		desc string) (string, error)

	DeleteChargeDef(
		UserId string,
		id string, // for edit
	) (string, error)
	//
	GetCompaniesSearch(UserId string, page int, nrows int, Id string, CompanyId string, filter string) ([]Company, error)
	//
	GetEmployeesSearchFast(UserId string, page int, nrows int, Id string, CompanyId string, filter string) ([]Employee, error)
	//
	GetEmployeesSearch(UserId string, page int, nrows int, Id string, CompanyId string, filter string) ([]Employee, error)
	//
	GetChargeTypes(UserId string) ([]ChargeType, error)

	// special cruds
	CrudInsuranceDasboard(UserId string, Id string, dewa string, dewa_expiration string) (string, error)

	GetBusinessCategories(UserId string) ([]BusinessCategory, error)
	CrudDasboards(
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
		dewaExpiration string) (string, error)
}
