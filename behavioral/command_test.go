package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	tv := &TV{}
	onCommand := &OnCommand{tv}
	offCommand := &OffCommand{tv}
	onButton := &Button{onCommand}
	offButton := &Button{offCommand}

	onButton.Press()
	assert.True(t, tv.isRunning)

	offButton.Press()
	assert.False(t, tv.isRunning)
}
