package parser

import (
	"fmt"
	"strings"

	"github.com/deparr/builder/pkg/v2/render"
	"github.com/deparr/builder/pkg/v2/token"
)

// grammar
//
// <Header> ::= <Octothorpe> <Text> <Endl>
// <Octothorpe> ::= #
// <Endl> ::= \n
// <Text> ::= \text
func Parse(tokens []token.Token) ([]render.Renderable, error) {
	if len(tokens) < 1 {
		return nil, fmt.Errorf("cannot parse an empty slice")
	}

	i := 0
	for i < len(tokens) && tokens[i].Type == token.NEWLINE_T {
		i++
	}

	renders := make([]render.Renderable, 0, 10)
	// Each iteration should be at beginning of a new line
	paragraph := make([]render.Renderable, 0, 10)
	for ; i < len(tokens); i++ {
		tok := tokens[i]
		switch tok.Type {
		case token.OCTOTHORPE_T:
			level := 0
			for ; tokens[i].Type == token.OCTOTHORPE_T && level < 6; i++ {
				level++
			}
			end := nextNewline(tokens[i:]) + i
			text := joinAsString(tokens[i:end])

			h := render.NewHeader(level, text)
			renders = append(renders, h)

			i = end

		case token.H_RULE_T:
			renders = append(renders, render.HRule{})

		case token.COMMENT_OPEN_T:
			var j int
			for j = i; j < len(tokens) && tokens[j].Type != token.COMMENT_CLOSE_T; {
				j++
			}
			if j == len(tokens) {
				fmt.Println("UNCLOSED COMMENT AT: (todo token locs)")
				break
			}
			i = j
			fmt.Println("TODO: PARSE DIRECTIVES / COMMENTS")

		case token.EOF_T:
			fallthrough
		case token.NEWLINE_T:
			if len(paragraph) > 0 {
				renders = append(renders, render.Paragraph{Body: paragraph})
				clear(paragraph)
			}

		case token.BACKTICK_T:
			fmt.Println("TODO CODEBLOCK")
			if tokens[i+1].Type != token.BACKTICK_T {
			}

		case token.DASH_T:
			fmt.Println("TODO LISTS")
			fallthrough
		// paragraph
		default:
			fmt.Println("TODO PARAGRAPH")
		}
	}
	return renders, nil
}

func parseParagraphLine(ts []token.Token) []render.Renderable {
	active := render.Plain
	joinable := make([]token.Token, 0, 10)
	styleStart := map[render.TextStyle]int{}
	for i := 0; i < len(ts); {
		t := ts[i]
		switch t.Type {
		case token.ASTERISK_T:
			if active == render.Plain {
			}

		case token.UNDERSCORE_T:
			if active == render.Plain {
			}
		}
	}
}

func next(ts []token.Token, tt token.TokenType) int {
	for i, t := range ts {
		if t.Type == tt {
			return i
		}
	}
	return len(ts)
}

// returns offset of nextnewline
func nextNewline(ts []token.Token) int {
	return next(ts, token.NEWLINE_T)
}

func joinAsString(ts []token.Token) string {
	strs := make([]string, len(ts))
	for i := range len(ts) {
		strs[i] = ts[i].ToLiteral()
	}
	return strings.Join(strs, "")
}
