package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {
	tv := NewTV()
	radio := NewRadio()

	remote := NewRemote(tv)
	remote.ChannelUp()

	remote.Device = radio
	remote.VolumeUp()

	assert.Greater(t, tv.channel, radio.channel)
	assert.Greater(t, radio.volume, tv.volume)
}
