package parser

import "fmt"

type Renderable interface {
	ToHtml(d *directive) string
}

type paragraph struct {
	content []Renderable
}

type text struct {
	string
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


type header struct {
	level int
	s     string
}

func (h header) ToHtml(d *directive) string {
	return fmt.Sprintf("<h%d>%s<h%d>", h.level, h.s, h.level)
}

type hr struct{}

func (h hr) ToHtml(d *directive) string {
	return "<hr>"
}

type br struct {}

func (b br) ToHtml(d *directive) string {
	return "<br>"
}


type directive struct {
	kind string
	args []string
}

type code struct{}

type codeBlock struct {
	lang    string
	content string
}
