package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposite(t *testing.T) {
	n := NewDirectory(
		"d1",
		NewFile("f1", 100),
		NewDirectory(
			"d2",
			NewFile("f2", 200),
		),
		NewDirectory(
			"d3",
			NewDirectory(
				"d4",
				NewFile("f3", 300),
				NewFile("f4", 400),
			),
		),
	)
	assert.Equal(t, 1040, calculateSize(n))
	assert.Equal(t, "d1/d3/d4/f4", searchFile("f4", n))
}

func calculateSize(n Node) int {
	return n.Size()
}

func searchFile(keyword string, n Node) string {
	return n.Search(keyword)
}
