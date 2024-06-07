package parser

import (
	"fmt"
	"strings"
)

type textTransform int

const (
	textBold = iota
	textItalic
	textStrike
	textUnder
	textPlain
)

type Renderable interface {
	ToHtml(d *directive) string
}

type paragraph struct {
	content []Renderable
}

func (p paragraph) ToHtml(d *directive) string {
	lines := make([]string, 0, 5)
	lines = append(lines, "<p>")
	for _, e := range p.content {
		// TODO: check for directives here
		lines = append(lines, e.ToHtml(d))
	}
	lines = append(lines, "</p>")
	return strings.Join(lines, "\n")
}

// TODO: quote renderable
// TODO: list renderable
//list TODO: table renderable

type text struct {
	textTransform
	string
}

func (t text) ToHtml(d *directive) string {
	var tag string
	switch t.textTransform {
	case textBold:
		tag = "strong"
	case textItalic:
		tag = "em"

	// TODO: I dont think this works like this
	// case textUnder
	// case textStrike
	default:
		tag = "span"
		
	}
	return fmt.Sprintf("<%s>%s</%s>", tag, t.string, tag)
}

// type span struct {
// 	string
// }
//
// func (s span) ToHtml(d *directive) string {
// 	return fmt.Sprintf("<span>%s</span>", s.string)
// }
//
// type italic struct {
// 	string
// }
//
// func (i italic) ToHtml(d *directive) string {
// 	return fmt.Sprintf("<em>%s</em>", i.string)
// }
//
// type bold struct {
// 	string
// }
//
// func (b bold) ToHtml(d *directive) string {
// 	return fmt.Sprintf("<strong>%s</strong>", b.string)
// }

type link struct {
	display string
	url     string
}

func (l link) ToHtml(d *directive) string {
	return fmt.Sprintf("<a href=\"%s\">%s</a>", l.display, l.url)
}

type header struct {
	level int
	s     string
}

func (h header) ToHtml(d *directive) string {
	return fmt.Sprintf("<h%d>%s</h%d>", h.level, h.s, h.level)
}

type hr struct{}

func (h hr) ToHtml(d *directive) string {
	return "<hr>"
}

type br struct{}

func (b br) ToHtml(d *directive) string {
	return "<br>"
}

type directive struct {
	kind string
	args []string
}

type code struct {
	string
}

type codeBlock struct {
	lang    string
	content string
}
