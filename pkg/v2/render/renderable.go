package render

import (
	"fmt"
	"strings"
)

type TextStyle int

const (
	Plain  TextStyle = 0
	Bold   TextStyle = 1 << 0
	Italic TextStyle = 1 << 1
	Under  TextStyle = 1 << 2
	Strike TextStyle = 1 << 3
	Upper  TextStyle = 1 << 4
	Lower  TextStyle = 1 << 5
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

type paragraph struct {
	Body []Renderable
}

func (p paragraph) RenderHTml() string {
	res := make([]string, len(p.Body)+2)
	res[0] = "<p>"
	for i, t := range p.Body {
		res[i+1] = t.RenderHtml()
	}
	res[len(res)-1] = "</p>"

	return strings.Join(res, "\n")
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

type span struct {
	Text  string
	Style TextStyle
}

func (s span) RenderHtml() string {
	if s.Style == Plain {
		return s.Text
	}
	styles := resolveStyles(s.Style)
	classes := strings.Join(styles, " ")
	return fmt.Sprintf("<span class=\"%s\">%s</span>", classes, s.Text)
}
