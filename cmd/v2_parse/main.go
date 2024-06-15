package main

import (
	"fmt"

	"github.com/deparr/builder/pkg/v2/lexer"
	"github.com/deparr/builder/pkg/v2/parser"
)

func main() {
// 	file := `
// # Header
// ---
//
// A paragraph with a single line.
//
// This paragraph has two
// lines (in source).
//
// Paragraph with link [github](https://github.com/deparr) 
// [](https://github.com/deparr)
//
// only special chars are '*_#-[]()<'
//
// **Bold text**
//
//
// *Italic Text*
//
// <!--date-->
// <!--sh "echo"-->
// <!--style class=red-->
//
// ####### header 2
// #header
//
// **bold text that never gets closed
// `

	file := `# header
### ### a real header * with symbols` + "```rust\n```"

	l := lexer.New(file)
	tokens, err := l.Tokenize()
	fmt.Println(tokens)
	if err != nil {
		println("err:", err)
		return
	}

	parsed, err := parser.Parse(tokens)
	fmt.Println(parsed)
	for _, r := range parsed {
		fmt.Println(r.RenderHtml())
	}
}
