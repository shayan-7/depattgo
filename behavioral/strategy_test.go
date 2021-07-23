package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategy(t *testing.T) {
	o := Operation(&Addition{})
	c := NewCalculator(100, 60, o)
	assert.Equal(t, 160, c.Calculate())

	o = &Substraction{}
	c.SetOperation(o)
	assert.Equal(t, 40, c.Calculate())
}
