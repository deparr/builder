package main

import "fmt"
import "github.com/deparr/builder/pkg/parse"

func main() {
	file := `
	A paragraph with **bold text**
	`

	r := ParseMd(file)

	fmt.Println(r.ToHtml())
}
