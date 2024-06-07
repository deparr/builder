package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/deparr/builder/pkg/parser"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		os.Exit(1)
	}
	file, err := os.ReadFile(args[1])
	p := parser.New(string(file))
	res, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, e := range res {
		fmt.Printf("%T ", e)
		fmt.Printf("%v\n", e)
	}

	html := make([]string, len(res)+2)
	html[0] = "<!doctype html>\n<html>\n<body>"
	for i, e := range res {
		html[i+1] = e.ToHtml(nil)
	}
	html[len(html)-1] = "</body>\n</html>"

	outfile := strings.Join(html, "\n")

	out, err := os.Create("out.html")

	_, err = out.Write([]byte(outfile))
}
