package parse_test

import (
	"testing"
	"github.com/deparr/templar/pkg/parse"
)



func TestParseMd(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("got %d; want 1", got)
	}
}
