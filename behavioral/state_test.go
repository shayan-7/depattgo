package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	v := NewVendingMachine(10, 10)
	r := NewTransmitionRequest(120, 2)

	err := v.MakeRequest(r)
	assert.Equal(t, "has item", v.Current.GetValue())
	assert.Nil(t, err)

	err = v.Proceed()
	assert.Equal(t, "item requested", v.Current.GetValue())
	assert.Nil(t, err)

	err = v.Proceed()
	assert.Equal(t, "has money", v.Current.GetValue())
	assert.Nil(t, err)

	err = v.Proceed()
	assert.Nil(t, err)
	assert.Nil(t, v.Request)
	assert.Equal(t, "has item", v.Current.GetValue())
	assert.Equal(t, 100, r.Money)

	err = v.Proceed()
	assert.NotNil(t, err)
	assert.Nil(t, v.Request)

	r.ItemCount = v.ItemCount
	v.MakeRequest(r)
	for i := 1; i < 4; i++ {
		err = v.Proceed()
	}
	assert.Equal(t, 20, r.Money)
	assert.Equal(t, "no item", v.Current.GetValue())
	assert.Nil(t, err)
	assert.Nil(t, v.Request)

	v.ItemCount = 20
	v.MakeRequest(r)
	err = v.Proceed()
	assert.Nil(t, err)
	assert.NotNil(t, v.Request)
	assert.Equal(t, "has item", v.Current.GetValue())
	assert.Equal(t, 20, r.Money)
}
