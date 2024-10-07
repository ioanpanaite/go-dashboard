package models

import (
	"database/sql"
	"time"
)

//Internal Invoice Model
// [internal_invoices_id internal_invoices_uuid internal_invoices_month internal_invoices_year
// internal_invoices_date companies_id companies_name internal_invoices_created_by
// internal_invoices_created_at internal_invoices_updated_at internal_invoices_currency_types_id
// internal_invoices_amount internal_invoices_currency_rate client_invoices_count]

type InternalInvoice struct {
	ID                  uint64          `db:"internal_invoices_id"`
	UUID                string          `db:"internal_invoices_uuid"`
	Month               int             `db:"internal_invoices_month"`
	Year                int             `db:"internal_invoices_year"`
	Date                string          `db:"internal_invoices_date"`
	CompanyID           uint64          `db:"companies_id"`
	CompanyName         string          `db:"companies_name"`
	CreatedBy           uint64          `db:"internal_invoices_created_by"`
	CreatedAt           string          `db:"internal_invoices_created_at"`
	UpdatedAt           string          `db:"internal_invoices_updated_at"`
	CurrencyTypesID     uint64          `db:"internal_invoices_currency_types_id"`
	Amount              float64         `db:"internal_invoices_amount"`
	CurrencyRate        sql.NullFloat64 `db:"internal_invoices_currency_rate"`
	ClientInvoicesCount uint64          `db:"client_invoices_count"`
	IsClientInvoices    bool
}

type ClientInvoice struct {
	ID                uint64         `db:"client_invoices_id"`
	UUID              string         `db:"client_invoices_uuid"`
	Date              string         `db:"client_invoices_date"`
	CompanyID         uint64         `db:"companies_id"`
	InternalInvoiceID uint64         `db:"internal_invoices_id"`
	CreatedBy         uint64         `db:"client_invoices_created_by"`
	CreatedAt         string         `db:"client_invoices_created_at"`
	UpdatedAt         string         `db:"client_invoices_updated_at"`
	CurrencyRate      sql.NullString `db:"client_invoices_currency_rate"`
	InoviceNumber     string         `db:"invoice_number"`
	CompanyName       string
	PaidAt            string
	Template          string
}

// TODO:
type ClientInvoiceTemplate struct {
	ID                uint64    `db:"client_invoice_id"`
	UUID              string    `db:"client_invoice_uuid"`
	InternalInvoiceID uint64    `db:"internal_invoice_id"`
	CreatedBy         uint64    `db:"created_by"`
	Currency          string    `db:"currency"`
	CurrencyRate      string    `db:"currency_rate"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	TemplateTypeID    string    `db:"template_type_id"`
	InvoiceNumber     string    `db:"invoice_number"`
	InvoiceDate       time.Time `db:"invoice_date"`
	CompanyID         uint64    `db:"company_id"`
	InvoiceableUUID   string    `db:"invoiceable_uuid"`
	InvoiceableType   string    `db:"invoiceable_type"`
}

type InvoiceSign struct {
	ChainID    int64  `db:"companies_user_signatures_chain_id"`
	ChainUUID  string `db:"companies_user_signatures_chain_uuid"`
	CompanyID  int64  `db:"companies_id"`
	UserID     int64  `db:"users_id"`
	ChainOrder int    `db:"companies_user_signatures_chain_order"`
	Timestamp  string `db:"internal_invoices_users_signatures_timestamp"`
	Comment    string `db:"internal_invoices_users_signatures_comment"`
}

type InvoiceCharge struct {
	ID                  uint32         `db:"id"`
	UUID                string         `db:"uuid"`
	ChargesDefID        uint64         `db:"charges_definition_id"`
	ChargesDefTypesName string         `db:"charges_definition_types_name"`
	ChargesDefTypesDesc sql.NullString `db:"charges_definition_types_description"`
	Year                uint32         `db:"year"`
	Month               uint32         `db:"month"`
	Date                string         `db:"date"`
	Amount              float64        `db:"amount"`
	IsManual            bool           `db:"is_manual"`
	PaidAt              sql.NullString `db:"paid_at"`
	PaidBy              sql.NullInt64  `db:"paid_by"`
	ReferenceUUID       string         `db:"reference_uuid"`
	IsPrepay            bool           `db:"is_prepay"`
	ReferenceID         int64          `db:"reference_id"`
	ReferenceType       string         `db:"reference_type"`
	InvoiceID           sql.NullInt64  `db:"internal_invoices_id"`
	IsPayroll           bool
}

type InvoiceTemplate struct {
	ID           int
	UUID         string
	Type         string
	TemplateName string
	Description  string
	DisplayName  string
}

type InvoicePayroll struct {
	ID         uint64        `db:"payrolls_id"`
	UUID       string        `db:"payrolls_uuid"`
	Year       int           `db:"payrolls_year"`
	Month      int           `db:"payrolls_month"`
	Status     string        `db:"payrolls_status"`
	TotalPrice float64       `db:"payrolls_total_price"`
	CreatedAt  string        `db:"payrolls_created_at"`
	UpdatedAt  string        `db:"payrolls_updated_at"`
	InvoiceID  sql.NullInt64 `db:"internal_invoices_id"`
}
