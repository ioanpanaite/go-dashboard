# Features not working from React Project

- Client Invoice Actions are not working
- Inovice template not properly working
- Client Invoice delete feature not working

- Invoice Get Sign return data is different in database and react app

`
Data schema from database:
Column Types:
companies_user_signatures_chain_id: BIGINT
companies_user_signatures_chain_uuid: CHAR
companies_id: BIGINT
users_id: BIGINT
companies_user_signatures_chain_order: INT
internal_invoices_users_signatures_timestamp: VARCHAR
internal_invoices_users_signatures_comment: VARCHAR

Data schema from react app:
{
      reference: e.internal_invoices_id,
      id: e.internal_invoices_id,
      uuid: e.internal_invoices_uuid,
      company_id: e.companies_id,
      company_name: e.companies_name,
      date: new Date(e.internal_invoices_year, e.internal_invoices_month - 1)
        .toISOString()
        .split("T")[0],
      state: "state",
      charges: 0,
      payroll: 0,
      company: 0,
      amount: e.internal_invoices_amount,
      created_by: e.internal_invoices_created_by,
      created_at: e.internal_invoices_created_at.split(" ")[0],
      updated_at: e.internal_invoices_updated_at.split(" ")[0],
      currency_rate: e.internal_invoices_currency_rate,
      // client invoice info
      client_invoices: e.client_invoices_count != 0,
      client_invoices_n: e.client_invoices_count,
      // permissiones
      user_can_create: true,
      user_can_edit: true,
      user_can_sign: true,
    }
`
