package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/deparr/builder/pkg/parser"
	"github.com/deparr/builder/pkg/templates"
)

type tempDat struct {
	Locale   string
	HeadTags []string
	Body     []string
	Footer   string
}

func main() {
	const root = "examples/basic/index.md"
	f, err := os.ReadFile(root)
	if err != nil {
		fmt.Fprint(os.Stderr, err, ": exiting")
		os.Exit(1)
	}

	inFile := string(f)
	p := parser.New(inFile)

	body, err := p.Parse()
	if err != nil {
		fmt.Fprint(os.Stderr, err, ": exiting")
		return
	}

	bodyHtml := make([]string, len(body))
	for i := range bodyHtml {
		bodyHtml[i] = body[i].ToHtml(nil)
	}

	fmt.Println(bodyHtml)

	dat := tempDat{
		Locale: "en",
		Body:   bodyHtml,
	}

	os.Mkdir("build", 0o766)
	// todo: this sucks
	outFilename := "build/" + strings.Replace(path.Base(root), ".md", ".html", 1)
	outFile, err := os.Create(outFilename)
	if err != nil {
		fmt.Fprint(os.Stderr, err, ": exiting")
		os.Exit(1)
	}
	defer outFile.Close()

	err = templates.Page().Execute(outFile, dat)
	if err != nil {
		fmt.Fprint(os.Stderr, err, ": exiting")
		return
	}

}
