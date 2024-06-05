package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/deparr/builder/pkg/parser"
)

func main() {
	file := `# Header 1

[link](https://github.com/deparr/builder)
Some text after the link

*italic text*
**bold text**`

	p := parser.New(file)
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
