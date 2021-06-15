package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	first := NewSingleton()
	assert.Equal(t, 1, first.NO())

	second := NewSingleton()
	assert.Equal(t, first.ID(), second.ID())
}
