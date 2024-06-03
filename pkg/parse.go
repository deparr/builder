package parse

type Renderable interface {
	ToHtml(d directive) string
}

type paragraph struct {
	content Renderable
	next    Renderable
}

type link struct {
	display string
	url     string
}

type italic struct {
	Renderable
}

type bold struct {
	Renderable
}

type header struct {
	level int
	s     string
}

type hr struct{}

type directive struct {
	kind string
	args []string
}

type code struct{}

type codeBlock struct {
	lang    string
	content string
}

func ParseMd(file string) Renderable {
	return nil
}
