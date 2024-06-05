package parser

import (
	"fmt"
	"log"
)

type Parser struct {
	file    string
	pos     int
	readPos int
	ch      byte
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

// Consumes the next renderable section of the input file.
//
// Only returns errors that are unrecoverable.
func (p *Parser) nextToken() (Renderable, error) {
	var next Renderable
	var err tokenError
	log.Printf(">>Remaining:'%s'\ncur:0x%02x", p.file[p.pos:], p.ch)
	// var pos = p.pos
	// ch should always be on a
	switch p.ch {
	case '#':
		next, err = p.parseHeader()
		if err != OK {
			if err == shouldBeText {
				log.Fatal("todo: handle `headerIsText`")
			}
		}

	case '-':
		if y, end := p.peekStr("--\n"); y {
			next = hr{}
			p.setPos(end)
		} else if p.peekChar() == ' ' {
		}
		log.Println(">>hr OK")
		log.Println("PARSER:", p)

	case '\n':
		p.readChar()
		next, err = br{}, OK

	case '_':
		fallthrough
	case '*':
		if p.peekChar() == p.ch {
			next, err = p.parseBold(p.ch)
		} else {
			next, err = p.parseItalic(p.ch)
		}

		if err != OK {
			log.Fatal("unhandled error parsing bolditalic")
		}

	case 0:
		log.Println("EOF in parser")
		break
	default:
		log.Fatal("default in nextToken(), should be unreachable")
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
	for p.pos < len(p.file) {
		next, err = p.nextToken()
		if err != nil {
			log.Println("Some error in Parse()")
		}
		res = append(res, next)
	}

	return res, nil
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

	log.Println(">>header OK")
	return header{level, head}, OK
}

func (p *Parser) parseBold(delim byte) (bold, tokenError) {
	log.Fatal("bold unimplemented")
	return bold{}, OK
}

func (p *Parser) parseItalic(delim byte) (italic, tokenError) {
	p.readChar()
	start := p.pos
	for p.ch != '\n' && p.ch != 0 && p.ch != delim {
		p.readChar()
	}

	if p.ch == '\n' || p.ch == 0 {
		return italic{}, shouldBeText
	}

	s := p.file[start:p.pos]

	p.readToEOL()

	return italic{s}, OK
}

func (p Parser) String() string {
	return fmt.Sprintf("Parser{ch:0x%02x,pos:%d,readPos:%d}", p.ch, p.pos, p.readPos)
}