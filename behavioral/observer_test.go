package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	manager := NewEventManager()
	publisher := NewPublisher(manager)
	firstListener := NewListener(1)
	secondListener := NewListener(2)

	manager.Register(firstListener)
	publisher.Do()
	assert.Equal(t, 1, firstListener.NumOfNotification)
	assert.Equal(t, 0, secondListener.NumOfNotification)

	manager.Register(secondListener)
	publisher.Do()
	assert.Equal(t, 2, firstListener.NumOfNotification)
	assert.Equal(t, 1, secondListener.NumOfNotification)

	manager.UnRegister(firstListener)
	publisher.Do()
	assert.Equal(t, 2, firstListener.NumOfNotification)
	assert.Equal(t, 2, secondListener.NumOfNotification)

	manager.UnRegister(secondListener)
	publisher.Do()
	assert.Equal(t, 2, firstListener.NumOfNotification)
	assert.Equal(t, 2, secondListener.NumOfNotification)
}
