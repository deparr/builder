package main

import (
	"fmt"

	"github.com/deparr/builder/pkg/parser"
)

func main() {
	file := `# Header
## header 2
---

*italic text*
**bold text**`

	p := parser.New(file)
	res, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, e := range res {
		fmt.Printf("%v\n", e)
	}
}
