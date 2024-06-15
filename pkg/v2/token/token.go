package token

import (
	"fmt"
)

type TokenType int

const (
	EOF_T  TokenType = 0
	TEXT_T TokenType = iota
	ASTERISK_T
	UNDERSCORE_T
	OCTOTHORPE_T
	BACKTICK_T
	BANG_T
	DASH_T
	TILDE_T
	GREATER_T
	LESS_T

	L_BRACK_T
	R_BRACK_T
	L_PAREN_T
	R_PAREN_T

	COMMENT_OPEN_T
	COMMENT_CLOSE_T
	H_RULE_T

	TAB_T
	SPACE_T
	NEWLINE_T
)

const (
	ASTERISK   = '*'
	UNDERSCORE = '_'
	OCTOTHORPE = '#'
	BACKTICK   = '`'
	BANG       = '!'
	DASH       = '-'
	TILDE      = '~'
	GREATER    = '>'
	LESS       = '<'
	L_BRACK    = '['
	R_BRACK    = ']'
	L_PAREN    = '('
	R_PAREN    = ')'
	TAB        = '\t'
	SPACE      = ' '
	NEWLINE    = '\n'
)

type Position struct {
	Line       int
	ByteOffset int
	Char       int
}

type Token struct {
	Type    TokenType
	Literal string
}

func New(t TokenType, lit string) Token {
	return Token{t, lit}
}

func (t Token) String() string {
	var tt string
	switch t.Type {
	case EOF_T:
		tt = "EOF"
	case TEXT_T:
		tt = "TEXT"
	case ASTERISK_T:
		tt = "ASTERISK"
	case UNDERSCORE_T:
		tt = "UNDERSCORE"
	case OCTOTHORPE_T:
		tt = "OCTOTHORPE"
	case BACKTICK_T:
		tt = "BACKTICK"
	case BANG_T:
		tt = "BANG"
	case DASH_T:
		tt = "DASH"
	case TILDE_T:
		tt = "TILDE"
	case GREATER_T:
		tt = "GREATER"
	case LESS_T:
		tt = "LESS"
	case L_BRACK_T:
		tt = "LBRACK"
	case R_BRACK_T:
		tt = "RBRACK"
	case L_PAREN_T:
		tt = "LPAREN"
	case R_PAREN_T:
		tt = "RPAREN"
	case COMMENT_OPEN_T:
		tt = "COMMENT_OPEN"
	case COMMENT_CLOSE_T:
		tt = "COMMENT_CLOSE"
	case H_RULE_T:
		tt = "H_RULE"

	case TAB_T:
		tt = "TAB"
	case SPACE_T:
		tt = "SPACE"
	case NEWLINE_T:
		tt = "NEWLINE"

	default:
		tt = "UNDEFINED"
	}

	if t.Literal != "" {
		return fmt.Sprintf("%s(%s)", tt, t.Literal)
	}

	return tt
}

func (t Token) ToLiteral() string {
	switch t.Type {
	case TEXT_T:
		return t.Literal
	case ASTERISK_T:
		return "*"
	case UNDERSCORE_T:
		return "_"
	case OCTOTHORPE_T:
		return "#"
	case BACKTICK_T:
		return "`"
	case BANG_T:
		return "!"
	case DASH_T:
		return "-"
	case TILDE_T:
		return "~"
	case GREATER_T:
		return ">"
	case LESS_T:
		return "<"
	case L_BRACK_T:
		return "["
	case R_BRACK_T:
		return "]"
	case L_PAREN_T:
		return "("
	case R_PAREN_T:
		return ")"
	case COMMENT_OPEN_T:
		return "<!--"
	case COMMENT_CLOSE_T:
		return "-->"
	case H_RULE_T:
		return "---"

	case TAB_T:
		return "\t"
	case SPACE_T:
		return " "
	case NEWLINE_T:
		return "\n"

	default:
		return ""
	}
}

func IsSymbol(r rune) bool {
	return r == ASTERISK || r == UNDERSCORE || r == OCTOTHORPE ||
		r == BACKTICK || r == BANG || r == DASH || r == TILDE ||
		r == L_BRACK || r == R_BRACK || r == R_PAREN || r == L_PAREN ||
		r == GREATER || r == LESS
}

func IsBeginSymbol(r rune) bool {
	return r == ASTERISK || r == UNDERSCORE || r == BACKTICK ||
		r == BANG || r == TILDE || r == L_BRACK || r == GREATER
}

func IsText(r rune) bool {
	return r != ASTERISK && r != UNDERSCORE && r != BACKTICK &&
		r != BANG && r != TILDE && r != L_BRACK && r != R_BRACK &&
		r != R_PAREN && r != L_PAREN && r != GREATER && r != LESS
}
