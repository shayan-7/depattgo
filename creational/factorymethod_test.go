package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {
	rl, err := NewLogistics(Road)
	assert.Nil(t, err)
	assert.False(t, rl.IsDone())

	rl.Deliver()
	assert.True(t, rl.IsDone())
	assert.Equal(t, Truck, rl.GetVehicle())

	sl, err := NewLogistics(Sea)
	assert.Nil(t, err)
	assert.False(t, sl.IsDone())

	sl.Deliver()
	assert.True(t, sl.IsDone())
	assert.Equal(t, Ship, sl.GetVehicle())
}
