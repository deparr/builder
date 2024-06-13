package lexer

import (
	"fmt"
	"strings"

	"github.com/deparr/builder/pkg/v2/lexer/token"
)

type Lexer struct {
	file string
}

func New(input string) *Lexer {
	return &Lexer{input}
}

const (
	whitespace = " \t\r\n"

	commentPre  = "<!--"
	commentPost = "-->"
	hRule       = "---"
)

var commentPre0 = rune(commentPre[0])
var commentPost0 = rune(commentPost[0])

func (l *Lexer) Tokenize() ([]token.Token, error) {
	lines := strings.Split(l.file, "\n")
	fmt.Println(lines)
	tokens := make([]token.Token, 0, 20)
	endl := token.New(token.NEWLINE_T, "")
	for lineNr, line := range lines {
		line = strings.TrimLeft(line, whitespace)
		fmt.Println("tokeninzing: ", line)
		if len(line) < 1 {
			tokens = append(tokens, endl)
			continue
		}

		lineTok, err := TokenizeLine(line)
		if err != nil {
			fmt.Println("unhandled error in lexer::TokenizeLine:", err)
			continue
		}
		tokens = append(tokens, lineTok...)

		// TODO: rethink this
		if lineNr != len(lines)-1 && line != "" {
			tokens = append(tokens, endl)
		}
	}

	tokens = append(tokens, token.New(token.EOF_T, ""))

	return tokens, nil
}

// NOTE: this creates a slice every line
// make this more efficient
func TokenizeLine(line string) ([]token.Token, error) {
	runes := []rune(line)
	res := make([]token.Token, 0, 5)
	var tok token.Token
	for i := 0; i < len(runes); i++ {
		rune := runes[i]
		switch rune {
		case token.OCTOTHORPE:
			tok = token.New(token.OCTOTHORPE_T, "")
		case token.ASTERISK:
			tok = token.New(token.ASTERISK_T, "")
		case token.UNDERSCORE:
			tok = token.New(token.UNDERSCORE_T, "")
		case token.BANG:
			tok = token.New(token.BANG_T, "")
		case token.TILDE:
			tok = token.New(token.TILDE_T, "")
		case token.BACKTICK:
			tok = token.New(token.BACKTICK_T, "")
		case token.L_BRACK:
			tok = token.New(token.L_BRACK_T, "")
		case token.R_BRACK:
			tok = token.New(token.R_BRACK_T, "")
		case token.L_PAREN:
			tok = token.New(token.L_PAREN_T, "")
		case token.R_PAREN:
			tok = token.New(token.R_PAREN_T, "")

		case token.LESS:
			if strStrAt(runes, commentPre, i) {
				tok = token.New(token.COMMENT_OPEN_T, "")
				i += len(commentPre) - 1
				fmt.Println(i)
			} else {
				tok = token.New(token.LESS_T, "")
			}

		case token.DASH:
			if strStrAt(runes, commentPost, i) {
				tok = token.New(token.COMMENT_CLOSE_T, "")
				i += len(commentPost) - 1
			} else if strStrAt(runes, hRule, i) {
				tok = token.New(token.H_RULE_T, "")
				i += len(hRule) - 1
			} else {
				tok = token.New(token.DASH_T, "")
			}

		// text node
		default:
			start := i
			for ; i < len(runes) && token.IsText(runes[i]); i++ {
				// BUG: temporary solution
				// end text in directive, will end any text that has comment
				// suffix in it. Might not be a big deal though
				if runes[i] == token.DASH && strStrAt(runes, commentPost, i) {
					break
				}
			}

			end := i
			if end >= len(runes) {
				end = len(runes)
			} else {
				// walk i back to end of text
				// this sucks
				i--
			}

			t := string(runes[start:end])
			tok = token.New(token.TEXT_T, t)
		}

		res = append(res, tok)
	}

	return res, nil
}

// Checks if the needle is in the haystack at the specified offset.
//
// Returns true if needle is ""
func strStrAt(hs []rune, nd string, off int) bool {
	hlen := len(hs)
	for i, r := range nd {
		idx := i + off
		if idx >= hlen || r != hs[idx] {
			return false
		}
	}

	return true
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}
