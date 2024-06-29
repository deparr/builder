package node

// type BlockKind int

// const (
// 	Document BlockKind = iota
// 	BlockQuote
// 	Paragraph
// 	List
// 	Text
// 	ListItem
// 	Emph
// 	Strong
// 	SoftBreak
// 	HardBreak
// 	Heading
// )

// type Block struct {
// 	Kind     BlockKind
// 	Text     string
// 	Open     bool
// 	Children []*Block
// }

// func NewDoc() *Block {
// 	return &Block{
// 		Kind:     Document,
// 		Open:     true,
// 		Children: make([]*Block, 0, 5),
// 	}
// }
//

type Block interface {
	block()
}

type Html int

func (Html) block() {}

// html
const ()

type Literal string

func (Literal) block() {}

type Heading struct {
	Level  int
	Setext bool
}

func (Heading) block() {}

type Code struct {
	Info    string
	Literal string
	Fenced  bool
}

func (Code) block() {}

type Link struct {
	Url   string
	Title string
}

func (Link) block() {}

type NodeType int
type Node struct {
	Content string
	Type    NodeType

	Next       *Node
	Prev       *Node
	Parent     *Node
	FirstChild *Node
	LastChild  *Node

	startLine   int
	startColumn int
	endLine     int
	endColumn   int

	As Block
}

// const (
// 	HtmlBlock NodeType = iota ThematicBreak
// 	CodeBlock
// 	Text
// 	SoftBreak
// 	LineBreak
// 	Code
// 	HtmlInline
// )

func (n *Node) isLeaf() bool {
	switch n.Type {
	default:
		return false
	}
}
