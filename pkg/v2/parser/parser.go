package parser

import (
	"fmt"

	"github.com/deparr/builder/pkg/v2/token"
	"github.com/deparr/builder/pkg/v2/render"
)


// grammar
//
// <Header> ::= <Octothorpe> <Text> <Endl>
// <Octothorpe> ::= #
// <Endl> ::= \n
// <Text> ::= \text
func Parse(tokens []token.Token) ([]Renderable, error) {
	if len(tokens) < 1 {
		return nil, fmt.Errorf("cannot parse an empty slice")
	}

	i := 0
	for i < len(tokens) && tokens[i].Type == token.NEWLINE_T {
		i++
	}

	// Each iteration should be at beginning of a new line
	for ; i < len(tokens); i++ {
		tok := tokens[i]
		switch tok.Type {
		case token.OCTOTHORPE_T:
			if i + 1 < len(tokens)
		}
	}
	return nil, nil
}
