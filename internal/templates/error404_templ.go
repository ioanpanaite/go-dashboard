// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Error404(fullPage bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><script src=\"https://cdn.tailwindcss.com\"></script><meta charset=\"UTF-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fullPage {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>Page Not Found</title><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body class=\"flex justify-center items-center h-screen bg-gray-100\"><div class=\"w-1/2 border-2 border-gray-400 p-6 rounded-lg bg-white shadow-lg text-center\"><h1 class=\"text-xl font-bold mb-4\">404 - Page Not Found</h1><p class=\"mb-4\">Sorry, the page you are looking for could not be found.\r</p><a href=\"/\" class=\"text-blue-500 hover:text-blue-700\">Go Home</a></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}