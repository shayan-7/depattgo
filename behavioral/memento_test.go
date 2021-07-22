package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	o := &Originator{}
	c := &Command{make([]Memento, 0)}

	for _, v := range []string{"A", "B", "C"} {
		o.SetState(v)
		c.Append(o.Save())
	}

	for _, v := range []string{"C", "B", "A"} {
		c.Undo()
		assert.Equal(t, v, o.state)
	}
}
