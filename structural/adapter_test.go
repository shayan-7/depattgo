package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	m := NewMac()
	w := NewWindows()
	a := NewWindowsAdapter(w)
	c := NewClient()
	assert.Equal(t, c.GetPortSpeed(a), c.GetPortSpeed(m))
}
