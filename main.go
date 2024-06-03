package main

import (
	"fmt"
	"html/template"
	"os"
)

type dat = struct {
	HeadTags []string
	Body []string
	Footer string
	Locale string
}

func main() {
	tmpl, err := template.New("index.html").ParseFiles("./template/index.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error")
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	data := dat {
		HeadTags: []string{"head1", "head2"},
		Body: []string{"body1", "body2"},
		Footer: "",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
