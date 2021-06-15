package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractFactory(t *testing.T) {
	height, width := 100, 120
	invalidCategory := Category(-1)

	factory, err := GetFactory(invalidCategory)
	assert.NotNil(t, err)
	assert.Nil(t, factory)

	factory, err = GetFactory(VueCategory)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	button := factory.createButton()
	assert.Equal(t, button.show(), "I'm showing vue button")

	window := factory.createWindow(height, width)
	assert.Equal(
		t,
		window.render(),
		"Rendering Vue window by height: 100 and width: 120",
	)

	factory, err = GetFactory(ReactCategory)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	button = factory.createButton()
	assert.Equal(t, button.show(), "I'm showing react button")

	window = factory.createWindow(height, width)
	assert.Equal(
		t,
		window.render(),
		"Rendering React window by height: 100 and width: 120",
	)
}
