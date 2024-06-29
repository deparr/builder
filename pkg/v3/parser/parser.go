package parser

import (
	"strings"

	"github.com/deparr/builder/pkg/v3/node"
)

type blockStack struct {
	arr []*node.Block
	i   int
}

func newBlockStack(init ...*node.Block) blockStack {
	return blockStack{
		arr: init,
		i:   len(init) - 1,
	}
}

func (s *blockStack) top() *node.Block {
	return s.arr[s.i]
}

func (s *blockStack) pop() *node.Block {
	if s.i < 0 {
		return nil
	}

	v := s.arr[s.i]
	s.i--
	return v
}

func (s *blockStack) push(b *node.Block) {
	s.arr = append(s.arr, b)
	s.i++
}

// Phase 1 of parse
func ParseBlocks(lines []string) *node.Block {
	doc := node.NewDoc()
	childQueue := newBlockStack(doc)
	for i, l := range lines {
	}
	return doc
}

func Parse(file string) node.Block {
	lines := strings.Split(file, "\n")
	doc := ParseBlocks(lines)
	return doc
}
