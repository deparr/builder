package parser

import (
	"fmt"
	"log"
)

type position struct {
	pos     int
	readPos int
	ch      byte
}

type Parser struct {
	file string
	position
}

type tokenError int

const (
	OK = iota
	malformedHeader
	shouldBeText
)

func New(input string) *Parser {
	p := &Parser{
		file: input,
	}
	p.readChar()
	return p
}

var parserReset position

// Consumes the next renderable section of the input file.
//
// Only returns errors that are unrecoverable.
func (p *Parser) nextToken() (Renderable, error) {
	var next Renderable
	var err tokenError
	// log.Printf(">>Remaining:'%s'\ncur:0x%02x", p.file[p.pos:], p.ch)
	// var pos = p.pos
	parserReset = p.position
	switch p.ch {
	case '#':
		next, err = p.parseHeader()

	case '-':
		if y, end := p.peekStr("--\n"); y {
			next = hr{}
			p.setPos(end)
		} else if p.peekChar() == ' ' {
		}

	case '\n':
		p.readChar()
		next, err = br{}, OK

	case 0:
		log.Println("EOF in parser")

	default:
		next, err = p.parseParagraph()
	}

	if err != OK {
		if err == shouldBeText {
			p.position = parserReset
			next, err = p.parseParagraph()
		}
	}
	// todo: hmm
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	return next, nil
}

func (p *Parser) Parse() ([]Renderable, error) {
	log.Printf(">>parsing << END\n%s\nEND\nlen:%d\n", p.file, len(p.file))
	res := make([]Renderable, 0, 10)
	var next Renderable
	var err error
	for p.hasInput() {
		p.skipWhitespace()
		next, err = p.nextToken()
		if err != nil {
			log.Println("Some error in Parse():", err)
		}
		res = append(res, next)
	}

	return res, nil
}

func (p *Parser) hasInput() bool {
	return p.pos < len(p.file) || p.ch != 0
}

func (p *Parser) skipWhitespace() {
	for p.ch == ' ' || p.ch == '\t' || p.ch == '\n' || p.ch == '\r' {
		p.readChar()
	}
}

func (p *Parser) incPos(i int) {
	p.pos += i
	p.readPos = p.pos + 1
	p.ch = p.file[p.pos]
}

func (p *Parser) setPos(pos int) {
	p.pos = pos
	p.readPos = pos + 1
	p.ch = p.file[pos]
}

func (p *Parser) readChar() {
	if p.readPos >= len(p.file) {
		p.ch = 0
	} else {
		p.ch = p.file[p.readPos]
	}

	p.pos = p.readPos
	p.readPos += 1
}

func (p *Parser) readToEOL() {
	for p.ch != '\n' && p.ch != 0 {
		p.readChar()
	}
}

func (p *Parser) takeUntil(target byte) (string, tokenError) {
	start := p.pos
	for p.ch != 0 && p.ch != '\n' && p.ch != target {
		p.readChar()
	}
	if p.ch == 0 || (target != '\n' && p.ch == '\n') {
		return "", shouldBeText
	}
	s := p.file[start:p.pos]
	return s, OK
}

func (p *Parser) peekChar() byte {
	if p.readPos >= len(p.file) {
		return 0
	} else {
		return p.file[p.readPos]
	}
}

func (p *Parser) peekStr(str string) (bool, int) {
	end := p.readPos + len(str)
	return p.file[p.readPos:end] == str, end
}

func (p *Parser) parseHeader() (header, tokenError) {
	markStartPos := p.pos
	for p.ch == '#' {
		p.readChar()
	}

	if p.ch != ' ' || p.pos-markStartPos > 6 {
		return header{}, shouldBeText
	}
	level := p.pos - markStartPos

	// consume space
	p.readChar()

	headStart := p.pos
	for p.ch != '\n' && p.ch != 0 {
		p.readChar()
	}

	head := p.file[headStart:p.pos]

	// consume newline
	//	and silently ignore eof
	p.readChar()

	return header{level, head}, OK
}

func (p *Parser) parseParagraph() (paragraph, tokenError) {
	content := make([]Renderable, 0, 5)
	for p.ch != '\n' && p.ch != 0 {
		var next Renderable
		var err tokenError
		switch p.ch {
		case '_':
			fallthrough
		case '*':
			if p.peekChar() == p.ch {
				next, err = p.parseBold(p.ch)
			} else {
				next, err = p.parseItalic(p.ch)
			}

			if err != OK {
				log.Fatal("todo: unhandled error parsing bolditalic")
			}

		case '[':
			next, err = p.parseLink()

		default:
			// BUG: this is wrong, assumes fmts are always whole lines
			s, err := p.takeUntil('\n')
			if err != OK {
			}

			next = text{textPlain, s}
		}
		content = append(content, next)

		if p.ch == '\n' {
			if p.peekChar() == '\n' {
				break
			}
			p.readChar()
			log.Println("WARN: need to handle double space at end of line to insert br")
		}
	}

	return paragraph{content}, OK
}

func (p *Parser) parseBold(delim byte) (text, tokenError) {
	// consume delim
	p.readChar()
	p.readChar()

	start := p.pos
	for p.ch != '\n' && p.ch != 0 {
		p.readChar()
		if p.ch == delim && p.peekChar() == delim {
			break
		}
	}

	if p.ch == '\n' || p.ch == 0 {
		return text{}, shouldBeText
	}

	s := p.file[start:p.pos]
	p.readChar()
	p.readChar()

	return text{textBold, s}, OK
}

func (p *Parser) parseItalic(delim byte) (text, tokenError) {
	p.readChar()
	start := p.pos
	for p.ch != '\n' && p.ch != 0 && p.ch != delim {
		p.readChar()
	}

	if p.ch == '\n' || p.ch == 0 {
		return text{}, shouldBeText
	}

	s := p.file[start:p.pos]

	// consume delim
	p.readChar()

	return text{textItalic, s}, OK
}

func (p *Parser) parseLink() (link, tokenError) {
	// start := p.pos
	p.readChar()
	disp, err := p.takeUntil(']')

	if err != OK {
		log.Fatal("Non ok in parseLink")
	}

	p.readChar()

	if p.ch != '(' {
		return link{}, shouldBeText
	}

	p.readChar()
	url, err := p.takeUntil(')')
	if err != OK {
		log.Fatal("Non ok in parseLink")
	}
	p.readChar()

	if len(url) == 0 {
		return link{}, shouldBeText
	}
	if len(disp) == 0 {
		disp = url
	}

	return link{disp, url}, OK
}

func (p Parser) String() string {
	return fmt.Sprintf("Parser{ch:0x%02x,pos:%d,readPos:%d}", p.ch, p.pos, p.readPos)
}
