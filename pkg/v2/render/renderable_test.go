package render

import (
	"fmt"
	"slices"
	"testing"
)

// note: relies on ordering of expected values
func TestResolveStyles(t *testing.T) {
	const boldItalic = Bold | Italic
	const upperStrikeUnder = Upper | Strike | Under

	b := resolveStyles(Bold)
	if len(b) != 1 || b[0] != "bold" {
		fmt.Println("failed bold: ", b)
		t.Fail()
	}

	bi := resolveStyles(boldItalic)
	if len(bi) != 2 || !slices.Equal(bi, []string{"bold","italic"}) {
		fmt.Println("failed boldItalic: ", bi)
		t.Fail()
	}

	usu := resolveStyles(upperStrikeUnder)
	if len(usu) != 3 || !slices.Equal(usu, []string{"underline","strike", "upper"}) {
		fmt.Println("failed upperStrikeUnder: ", usu)
		t.Fail()
	}


}
