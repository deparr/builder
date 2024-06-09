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

const (
	rootDef         = "."
	rootUsage       = "root dir/file to start search in. setting to a file implies singleFile"
	localeDef       = "en"
	localeUsage     = "locale for html output"
	noRenderDef     = false
	noRenderUsage   = "only evaluate directives. do not render to html"
	singleFileDef   = false
	singleFileUsage = "evalute a single file only"
)

type cmdOpts struct {
	root     string
	locale   string
	noRender bool
}

func (c cmdOpts) String() string {
	return fmt.Sprintf(
		"cmdOpts{%s,%s,%t}",
		c.root,
		c.locale,
		c.noRender,
		// c.singleFile,
	)
}

func cliOpts() *cmdOpts {
	var (
		root     string
		locale   string
		noRender bool
	)

	flag.StringVar(&root, "root", rootDef, rootUsage)
	flag.StringVar(&root, "r", rootDef, rootUsage+" (shorthand)")
	flag.StringVar(&locale, "locale", localeDef, localeUsage)
	flag.StringVar(&locale, "L", localeDef, localeUsage+" (shorthand)")
	flag.BoolVar(&noRender, "no-html", noRenderDef, noRenderUsage)
	flag.BoolVar(&noRender, "N", noRenderDef, noRenderUsage+" (shorthand)")
	// flag.BoolVar(&singleFile, "f", singleFileDef, noRenderUsage)

	flag.Parse()

	return &cmdOpts{root, locale, noRender}
}

type tempDat struct {
	Locale   string
	HeadTags []string
	Body     []string
	Footer   string
}

func main() {
	opts := cliOpts()
	fmt.Println(opts)

	stat, err := os.Stat(opts.root)
	if err != nil {
		fmt.Fprint(os.Stderr, err, ": exiting")
		os.Exit(1)
	}

	if stat.IsDir() {
		// files, err := os.ReadDir(opts.root)
		fmt.Fprint(os.Stderr, "todo: implement recurisve root dir")
		os.Exit(1)
	} else {
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

		bodyHtml := make([]string, len(body))
		for i := range bodyHtml {
			bodyHtml[i] = body[i].ToHtml(nil)
		}

		fmt.Println(bodyHtml)

		dat := tempDat{
			Locale: opts.locale,
			Body:   bodyHtml,
		}

		os.Mkdir("build", 0o766)
		// todo: this sucks
		outFilename := "build/" + strings.Replace(path.Base(opts.root), ".md", ".html", 1)
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
}
