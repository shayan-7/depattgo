package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediator(t *testing.T) {
	station := &StationManager{isPlatformFree: true}

	firstTrain := NewPassengerTrain(station)
	secondTrain := NewPassengerTrain(station)
	thirdTrain := NewPassengerTrain(station)

	firstTrain.arrive()
	secondTrain.arrive()
	assert.Equal(t, StateArrived, firstTrain.state)
	assert.Equal(t, StateInQueue, secondTrain.state)
	assert.Equal(t, StateOnWay, thirdTrain.state)

	thirdTrain.arrive()
	firstTrain.depart()
	assert.Equal(t, StateDeparted, firstTrain.state)
	assert.Equal(t, StateArrived, secondTrain.state)
	assert.Equal(t, StateInQueue, thirdTrain.state)

	secondTrain.depart()
	assert.Equal(t, StateDeparted, secondTrain.state)
	assert.Equal(t, StateArrived, thirdTrain.state)

	thirdTrain.depart()
	assert.Equal(t, StateDeparted, thirdTrain.state)
}
