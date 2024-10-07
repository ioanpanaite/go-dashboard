// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package charge_view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"
import "kub/dashboardES/internal/templates/layout"

type ChargedefsViewData struct {
	Module string
	Page   string
	Route  string
}

func ChargedefsView(data ChargedefsViewData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\r\n    .active-tab {\r\n      background-color: gray; /* Example color, change as needed */\r\n      color: white;\r\n    }\r\n  </style> <div class=\"mx-auto\"><!-- Tab Headers --><div class=\"bg-gray-300 py-1\"><ul class=\"flex\"><li class=\"ml-1\"><button id=\"tab1\" class=\"tab-button bg-black text-white py-2 px-4 focus:outline-none rounded-md\" role=\"tab\" aria-selected=\"true\" aria-controls=\"panel1\">All Charges\r</button></li><li class=\"ml-1\"><button id=\"tab2\" class=\"tab-button bg-gray-800 text-white py-2 px-4 focus:outline-none rounded-md\" role=\"tab\" aria-selected=\"false\" aria-controls=\"panel2\">Grouped Charges\r</button></li></ul></div><!-- Company Charges Table (Initially Visible) --><div id=\"panel1\" class=\"tab-panel pt-2 container mx-auto\" role=\"tabpanel\" aria-labelledby=\"tab1\"><!-- Selector --><sl-radio-group size=\"small\" name=\"visualizeSelection\" value=\"all\" id=\"companyFilterSelection\" class=\"pt-3 px-2 container mx-auto\"><sl-radio-button data-target-id=\"searchEmployee-CDF\" value=\"all\" class=\"filterable-hx-rows\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/filterCompanyCharges?page=1&filter=All")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 68, Col: 83}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-target=\"#company-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><i class=\"fas fa-all\"></i> All\r</sl-radio-button> <sl-radio-button data-target-id=\"searchCompany-CDF\" value=\"company\" class=\"filterable-hx-rows\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/filterCompanyCharges?page=1&filter=Company")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 80, Col: 87}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-target=\"#company-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><i class=\"fas fa-building pr-2\"></i> Companies</sl-radio-button> <sl-radio-button data-target-id=\"searchEmployee-CDF\" value=\"employee\" class=\"filterable-hx-rows\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/filterCompanyCharges?page=1&filter=User")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 90, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-target=\"#company-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><i class=\"fas fa-user pr-2\"></i>Employees</sl-radio-button></sl-radio-group><div class=\"container w-full m-2 border-black shadow-lg rounded-md overflow-hidden\"><div class=\"bg-black text-white py-2 px-4 flex\"><h2>All Charges</h2><div class=\"grow\"></div><div class=\"flex bg-gray-800\"><sl-button class=\"hidden\" size=\"small\" id=\"prevPageCompany\" type=\"submit\"><i class=\"fas fa-chevron-left\"></i></sl-button> <sl-input id=\"companyPage\" size=\"small\" min=\"0\" step=\"1\" type=\"number\" readonly class=\"w-12 hidden\" value=\"0\"></sl-input> <sl-button class=\"hidden\" size=\"small\" id=\"nextPageCompany\" type=\"submit\"><i class=\"fas fa-chevron-right\"></i></sl-button></div><div class=\"w-2\"></div><!-- Add hx-headers='{\"reftype\":\"company\",\"refuuid\":\"{{ .Data.company.Uuid }}\"}' -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, show())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<sl-button size=\"small\" id=\"Add-charges\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/panels/invoicing/charges/%s/new", "_"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 136, Col: 68}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#addformdialog\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.ComponentScript = show()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Add Charges\r</sl-button><div class=\"w-2\"></div><sl-button size=\"small\" id=\"Refresh-charges\" class=\"filterable-hx-rows\" hx-on::before-request=\"document.getElementById(&#39;companyFilterSelection&#39;).value = &#39;all&#39;;\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/search?page=1&type=company")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 150, Col: 72}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-target=\"#company-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><sl-icon name=\"arrow-clockwise\"></sl-icon></sl-button></div><table class=\"min-w-full leading-normal\"><thead><tr><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-4\">T\r</th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-44\"><span>Reference<br>Name</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-44\"><span>Type<br>Concept</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Amount</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Period</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Monthly<br>N (Lack)</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Prepay<br>P. Month</span></th><th class=\"px-2 py-2 bg-gray-800 text-center text-xs font-semibold text-gray-300 uppercase tracking-wider w-20\"><span>Renew</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-32\"><span>Actions</span></th></tr></thead> <tbody id=\"company-tbody\" class=\"filterable-hx-rows\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/search?page=1&type=company")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 212, Col: 72}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"filterUpdateEvent\" hx-target=\"#company-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><td colspan=\"12\" class=\"px-1 py-1 border-b border-gray-200 text-sm text-center h-20\" id=\"spinner\">Filter to show results\r<!-- <sl-spinner class=\"text-2xl m-5\"></sl-spinner> --></td></tbody></table><div class=\"bg-black min-h-4 p-2 flex\"></div></div></div><!-- Employee Charges Table (Initially Hidden) --><div id=\"panel2\" class=\"tab-panel hidden pt-2 container mx-auto\" role=\"tabpanel\" aria-labelledby=\"tab2\"><!-- Selector --><sl-radio-group size=\"small\" name=\"visualizeSelection\" id=\"employeeFilterSelection\" value=\"all\" class=\"pt-3 px-2 container mx-auto\"><sl-radio-button data-target-id=\"searchEmployee-CDF\" value=\"all\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/filterEmployeeCharges?page=1&filter=All")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 249, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" class=\"filterable-hx-rows\" hx-target=\"#employee-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><i class=\"fas fa-all\"></i> All\r</sl-radio-button> <sl-radio-button data-target-id=\"searchCompany-CDF\" value=\"company\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/filterEmployeeCharges?page=1&filter=Company")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 261, Col: 88}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" class=\"filterable-hx-rows\" hx-target=\"#employee-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><i class=\"fas fa-building pr-2\"></i> Companies</sl-radio-button> <sl-radio-button data-target-id=\"searchEmployee-CDF\" value=\"employee\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/filterEmployeeCharges?page=1&filter=User")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 271, Col: 85}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" class=\"filterable-hx-rows\" hx-target=\"#employee-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><i class=\"fas fa-user pr-2\"></i>Employees</sl-radio-button></sl-radio-group><div class=\"container w-full m-2 border-black shadow-lg rounded-md overflow-hidden\"><div class=\"bg-black text-white py-2 px-4 flex\"><h2>Grouped charges</h2><div class=\"grow\"></div><div class=\"flex bg-gray-800\"><sl-button class=\"hidden\" size=\"small\" id=\"prevPageEmployee\" type=\"submit\"><i class=\"fas fa-chevron-left\"></i></sl-button> <sl-input class=\"hidden w-12\" id=\"employeePage\" size=\"small\" min=\"0\" step=\"1\" type=\"number\" readonly value=\"0\"></sl-input> <sl-button class=\"hidden\" size=\"small\" id=\"nextPageEmployee\" type=\"submit\"><i class=\"fas fa-chevron-right\"></i></sl-button></div><div class=\"w-2\"></div><sl-button size=\"small\" id=\"Refresh-Employee-charges\" class=\"filterable-hx-rows\" hx-on::before-request=\"document.getElementById(&#39;employeeFilterSelection&#39;).value = &#39;all&#39;;\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs("/panels/invoicing/chargedef/search?page=1&type=employee")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 319, Col: 73}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-target=\"#employee-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><sl-icon name=\"arrow-clockwise\"></sl-icon></sl-button></div><table class=\"min-w-full leading-normal\"><thead><tr><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-4\">T\r</th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-44\"><span>Reference<br>Name</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-44\"><span>Type<br>Concept</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Amount</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Period</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Monthly<br>N (Lack)</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-24\"><span>Prepay<br>P. Month</span></th><th class=\"px-2 py-2 bg-gray-800 text-center text-xs font-semibold text-gray-300 uppercase tracking-wider w-20\"><span>Renew</span></th><th class=\"px-2 py-2 bg-gray-800 text-left text-xs font-semibold text-gray-300 uppercase tracking-wider w-32\"><span>Actions</span></th></tr></thead><tbody id=\"employee-tbody\" class=\"filterable-hx-rows\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/panels/invoicing/chargedef/search?type=%v&companyId=%v", "employee", ""))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/invoicing/charge/chargedefs-view.templ`, Line: 382, Col: 102}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"filterUpdateEvent\" hx-target=\"#employee-tbody\" hx-swap=\"innerHTML\" hx-indicator=\".htmx-indicator\"><td colspan=\"12\" class=\"px-1 py-1 border-b border-gray-200 text-sm text-center h-20\" id=\"spinner\">Filter to show results\r<!-- <sl-spinner class=\"text-2xl m-5\"></sl-spinner> --></td></tbody></table><div class=\"bg-black min-h-4 p-2\"></div></div></div></div><script>\r\n    // JavaScript for Tab Functionality\r\n    document.addEventListener(\"DOMContentLoaded\", function () {\r\n      const tabs = document.querySelectorAll(\".tab-button\");\r\n      const tabPanels = document.querySelectorAll(\".tab-panel\");\r\n\r\n      tabs.forEach((tab) => {\r\n        tab.addEventListener(\"click\", () => {\r\n          tabs.forEach((t) => {\r\n            t.setAttribute(\"aria-selected\", \"false\");\r\n            t.classList.remove(\"bg-black\");\r\n            t.classList.add(\"bg-gray-800\");\r\n          });\r\n          tabPanels.forEach((panel) => panel.classList.add(\"hidden\"));\r\n\r\n          tab.setAttribute(\"aria-selected\", \"true\");\r\n          tab.classList.add(\"bg-black\");\r\n          tab.classList.remove(\"bg-gray-800\");\r\n          const panel = document.querySelector(\r\n            `#${tab.getAttribute(\"aria-controls\")}`,\r\n          );\r\n          panel.classList.remove(\"hidden\");\r\n        });\r\n      });\r\n    });\r\n\r\n    document.addEventListener(\"DOMContentLoaded\", function () {\r\n      const prevButtonCompany = document.querySelector(\"#prevPageCompany\");\r\n      const nextButtonCompany = document.querySelector(\"#nextPageCompany\");\r\n      const companytbody = document.querySelector(\"#company-tbody\");\r\n      let companyPage = 1; // Initialize current page\r\n\r\n      const prevButtonEmployee = document.querySelector(\"#prevPageEmployee\");\r\n      const nextButtonEmployee = document.querySelector(\"#nextPageEmployee\");\r\n      const employeetbody = document.querySelector(\"#employee-tbody\");\r\n      let employeePage = 1; // Initialize current page\r\n\r\n \tfunction extractQueryParams(url) {\r\n        var searchParams = new URLSearchParams(url.search);\r\n        var params = {};\r\n\r\n        searchParams.forEach(function(value, key) {\r\n          params[key] = value;\r\n        });\r\n\r\n        return params;\r\n      }\r\n// Function to update query parameters and convert back to URL\r\n      function updateAndConvertToURL(url, updatedParams) {\r\n        var searchParams = new URLSearchParams(url.search);\r\n\r\n        // Update the existing parameters with the new values\r\n        Object.keys(updatedParams).forEach(function(key) {\r\n          searchParams.set(key, updatedParams[key]);\r\n        });\r\n\r\n        // Set the updated search parameters back to the URL\r\n        url.search = searchParams.toString();\r\n\r\n        return url.pathname + url.search;\r\n      }\r\n\r\n      function updatePaginationParams() {\r\n        // Update button parameters\r\n        // const headers = JSON.parse(companytbody.getAttribute(\"hx-headers\"));\r\n        const qUrlCompany = companytbody.getAttribute(\"hx-get\");\r\n\t\tvar urlCompany = new URL(qUrlCompany, window.location.href)\r\n\t\tconst queriesCompany = extractQueryParams(urlCompany)\r\n\t\tqueriesCompany.page = companyPage\r\n\t\tconst newUrlCompany = updateAndConvertToURL(urlCompany, queriesCompany)\r\n\t\tcompanytbody.setAttribute(\"hx-get\",newUrlCompany)\r\n       // headers.page = companyPage;\r\n\r\n        // companytbody.setAttribute(\"hx-headers\", JSON.stringify(headers));\r\n\r\n        // Hide and disable the prevPageCompany button if the current page is 1\r\n        if (companyPage === 1) {\r\n          prevButtonCompany.style.display = \"none\";\r\n          prevButtonCompany.disabled = true;\r\n        } else {\r\n          prevButtonCompany.style.display = \"\";\r\n          prevButtonCompany.disabled = false;\r\n        }\r\n\r\n        // Update button parameters\r\n\t\tconst qUrlEmployee = employeetbody.getAttribute(\"hx-get\")\r\n\t\tvar urlEmployee = new URL(qUrlEmployee, window.location.href)\r\n\t\tconst queriesEmployee = extractQueryParams(urlEmployee)\r\n\t\tqueriesEmployee.page = employeePage\r\n\t\tconst newUrlEmployee = updateAndConvertToURL(urlEmployee,queriesEmployee)\r\n\t\temployeetbody.setAttribute(\"hx-get\",newUrlEmployee)\r\n        // const headersEmployee = JSON.parse(\r\n        //   employeetbody.getAttribute(\"hx-headers\"),\r\n        // );\r\n        // headersEmployee.page = employeePage;\r\n        // employeetbody.setAttribute(\r\n        //   \"hx-headers\",\r\n        //   JSON.stringify(headersEmployee),\r\n        // );\r\n        // Hide and disable the prevPageCompany button if the current page is 1\r\n        if (employeePage === 1) {\r\n          prevButtonEmployee.style.display = \"none\";\r\n          prevButtonEmployee.disabled = true;\r\n        } else {\r\n          prevButtonEmployee.style.display = \"\";\r\n          prevButtonEmployee.disabled = false;\r\n        }\r\n\r\n        document.getElementById(\"companyPage\").value = companyPage;\r\n        document.getElementById(\"employeePage\").value = employeePage;\r\n      }\r\n\r\n      // Event listeners for pagination buttons\r\n      prevButtonCompany.addEventListener(\"click\", function () {\r\n        if (companyPage > 1) {\r\n          companyPage--;\r\n          updatePaginationParams();\r\n        }\r\n        window.scrollTo({\r\n          top: 0,\r\n          left: 0,\r\n          behavior: \"smooth\",\r\n        });\r\n      });\r\n\r\n      nextButtonCompany.addEventListener(\"click\", function () {\r\n        companyPage++;\r\n        //console.log(companyPage);\r\n        updatePaginationParams();\r\n        window.scrollTo({\r\n          top: 0,\r\n          left: 0,\r\n          behavior: \"smooth\",\r\n        });\r\n      });\r\n\r\n      // Event listeners for pagination buttons\r\n      prevButtonEmployee.addEventListener(\"click\", function () {\r\n        if (employeePage > 1) {\r\n          employeePage--;\r\n          updatePaginationParams();\r\n        }\r\n        window.scrollTo({\r\n          top: 0,\r\n          left: 0,\r\n          behavior: \"smooth\",\r\n        });\r\n      });\r\n\r\n      nextButtonEmployee.addEventListener(\"click\", function () {\r\n        employeePage++;\r\n        //console.log(employeePage);\r\n        updatePaginationParams();\r\n        window.scrollTo({\r\n          top: 0,\r\n          left: 0,\r\n          behavior: \"smooth\",\r\n        });\r\n      });\r\n\r\n      updatePaginationParams(); // Initial update\r\n    });\r\n\r\n  </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base(false, data.Module, data.Page).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
