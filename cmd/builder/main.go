package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/deparr/builder/pkg/parser"
)

const (
	// singleFileDef   = false
	// singleFileUsage = "whether to "
	rootDef       = "."
	rootUsage     = "root dir to start search in"
	noRenderDef   = false
	noRenderUsage = "only evaluate directives. do not render to html"
	localeDef     = "en"
	localeUsage   = "locale for html output"
)

type cmdOpts struct {
	singleFile bool
	root       string
	noRender   bool
	locale     string
}

func (c cmdOpts) String() string {
	return fmt.Sprintf(
		"cmdOpts{%t,%s,%t,%s}",
		c.singleFile,
		c.root,
		c.noRender,
		c.locale,
	)
}

func cliOpts() *cmdOpts {
	var (
		singleFile bool
		root       string
		noRender   bool
		locale     string
	)

	flag.StringVar(&root, "root", rootDef, rootUsage)
	flag.StringVar(&root, "r", rootDef, rootUsage+" (shorthand)")
	flag.BoolVar(&noRender, "no-html", noRenderDef, noRenderUsage)
	flag.BoolVar(&noRender, "N", noRenderDef, noRenderUsage+" (shorthand)")
	flag.StringVar(&locale, "locale", localeDef, localeUsage)
	flag.StringVar(&locale, "L", localeDef, localeUsage+" (shorthand)")

	flag.Parse()

	return &cmdOpts{singleFile, root, noRender, locale}
}

func main() {
	opts := cliOpts()
	fmt.Println(opts)

	f, err := os.ReadFile(opts.root)
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

	fmt.Println(body)

}
