package templates

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed template/page.html
var pageTmpl string
var page *template.Template

// TODO: rethink how this module should be loaded and structured
func Page() *template.Template {
	if page != nil {
		return page
	}

	page, err := template.New("page").Parse(pageTmpl)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	return page
}

