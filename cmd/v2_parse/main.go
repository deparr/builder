package main

import (
	"fmt"

	"github.com/deparr/builder/pkg/v2/lexer"
	"github.com/deparr/builder/pkg/v2/lexer/token"
)

func main() {
	file := `
# Header
---

A paragraph with a single line.

This paragraph has two
lines (in source).

Paragraph with link [github](https://github.com/deparr) 
[](https://github.com/deparr)

only special chars are '*_#-[]()<'

**Bold text**


*Italic Text*

<!--date-->
<!--sh "echo"-->
<!--style class=red-->

####### header 2
#header

**bold text that never gets closed
`

	// file := `<!--class bold,big,italic-->`

	l := lexer.New(file)
	res, err := l.Tokenize()
	fmt.Println(res)
	if err != nil {
		println("err:", err)
		return 
	}
	

	for _, r := range res {
		fmt.Printf("%s ", r)
		if r.Type == token.NEWLINE_T {
			fmt.Println()
		}
	}
}
