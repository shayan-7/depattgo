package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	sportBuilder := NewSportCarBuilder()
	director := NewDirector(sportBuilder)
	sportCar := director.Construct()
	assert.Equal(t, sportCar.engine, SportEngine)
	assert.Equal(t, sportCar.seats, SportSeats)
	assert.Equal(t, sportCar.turbo, true)
	assert.Equal(t, sportCar.trip, false)

	suvBuilder := NewSUVCarBuilder()
	director.SetBuilder(suvBuilder)
	suvCar := director.Construct()
	assert.Equal(t, suvCar.engine, SUVEngine)
	assert.Equal(t, suvCar.seats, SUVSeats)
	assert.Equal(t, suvCar.turbo, false)
	assert.Equal(t, suvCar.trip, true)
}
