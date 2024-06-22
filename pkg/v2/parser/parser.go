package parser

import (
	"fmt"
	"slices"
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
			endline := nextNewline(tokens[i:])
			parseParagraphLine(tokens[i : i+endline])
			i += endline
		}
	}
	return renders, nil
}

func parseParagraphLine(ts []token.Token) []render.Renderable {
	type qnode struct {
		dc int
		d  rune
		s  int
	}

	q := make([]qnode, 10)
	top := 0
	ranges := make([][]int, 0, 4)
	for i := 0; i < len(ts); {
		t := ts[i]
		switch t.Type {
		case token.ASTERISK_T:
			fallthrough
		case token.UNDERSCORE_T:
			delimStart := i
			for ; i < len(ts) && ts[i].Type == t.Type; i++ {
			}
			dc := i - delimStart
			d := rune(t.ToLiteral()[0])
			s := i
			if top >= 0 && dc == q[top].dc {
				snode := q[top]
				top -= 1
				if snode.d != d {
					panic("mismatched delim with matching count")
				}

				ranges = append(ranges, []int{snode.s, delimStart})
			} else {
				top += 1
				q[top] = qnode{dc, d, s}
			}
		default:
			i++
		}

	}

	fmt.Println(ranges)
	slices.SortFunc(ranges, func(a, b []int) int {
		if b[0] == a[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	// styles := []render.TextStyle{render.Plain, render.Italic, render.Bold, render.Bold | render.Italic}
	fmt.Println(ranges)

	return nil
}

func next(ts []token.Token, tt token.TokenType) int {
	for i, t := range ts {
		if t.Type == tt || t.Type == token.NEWLINE_T {
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
