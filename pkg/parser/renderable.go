package parser

import (
	"fmt"
	"strings"
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

// TODO: combine text nodes into a single type
type span struct {
	string
}

func (s span) ToHtml(d *directive) string {
	return fmt.Sprintf("<span>%s</span>", s.string)
}

type italic struct {
	string
}

func (i italic) ToHtml(d *directive) string {
	return fmt.Sprintf("<em>%s</em>", i.string)
}

type bold struct {
	string
}

func (b bold) ToHtml(d *directive) string {
	return fmt.Sprintf("<strong>%s</strong>", b.string)
}

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
