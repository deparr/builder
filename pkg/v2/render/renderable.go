package render

import (
	"fmt"
	"strings"
)

type TextStyle int

const (
	Plain TextStyle = iota
	Bold  TextStyle = 1 << (iota - 1)
	Italic
	Under
	Strike
	Upper
	Lower
	Code
)

type Renderable interface {
	RenderHtml() string
}

type Document struct {
	Body   []Renderable
	Footer Renderable
}

func (d *Document) AddChild(c Renderable) {
	d.Body = append(d.Body, c)
}

type Node struct {
	Tag     string
	Classes []string
}

type Paragraph struct {
	Body []Renderable
}

func NewParagraph(body []Renderable) Paragraph {
	return Paragraph{body}
}

// NOTE:  this might not want to be nweline joined
func (p Paragraph) RenderHtml() string {
	res := make([]string, len(p.Body)+2)
	res[0] = "<p>"
	for i, t := range p.Body {
		res[i+1] = t.RenderHtml()
	}
	res[len(res)-1] = "</p>"

	return strings.Join(res, "\n")
}

type Header struct {
	Level int
	Text  string
}

func NewHeader(l int, t string) Header {
	if l > 6 {
		l = 6
	}
	if l < 1 {
		l = 1
	}
	t = strings.TrimSpace(t)
	return Header{Level: l, Text: t}
}

func (h Header) RenderHtml() string {
	return fmt.Sprintf("<h%d>%s</h%d>", h.Level, h.Text, h.Level)
}

func style2String(s TextStyle) string {
	switch s {
	case Plain:
		return ""
	case Bold:
		return "bold"
	case Italic:
		return "italic"
	case Under:
		return "underline"
	case Strike:
		return "strike"
	case Upper:
		return "upper"
	case Lower:
		return "lower"
	default:
		return ""
	}
}

func resolveStyles(s TextStyle) []string {
	if s == Plain {
		return []string{}
	}

	styles := make([]string, 0, 2)

	for i := 0; i < 32 && s != Plain; i++ {
		c := s & (1 << i)
		if c != 0 {
			styles = append(styles, style2String(c))
		}
		s &= ^c
	}

	return styles
}

type Span struct {
	Text  string
	Style TextStyle
}

func (s Span) RenderHtml() string {
	if s.Style == Plain {
		return s.Text
	}
	styles := resolveStyles(s.Style)
	classes := strings.Join(styles, " ")
	return fmt.Sprintf("<span class=\"%s\">%s</span>", classes, s.Text)
}

type HRule struct{}

func (h HRule) RenderHtml() string {
	return "<hr>"
}

// open target?
type Anchor struct {
	Text  string
	Style TextStyle
	Url   string
}

func (a Anchor) RenderHtml() string {
	styles := resolveStyles(a.Style)
	classes := strings.Join(styles, " ")
	return fmt.Sprintf("<a class=\"%s\" href=\"%s\">%s</a>", classes, a.Url, a.Text)
}
