package creational

import (
	"errors"
)

const (
	Road = iota
	Sea
)

type Vehicle int

const (
	Truck Vehicle = iota
	Ship
)

type Creator func() Logistics

var TruckCreator = Creator(func() Logistics {
	return &RoadLogistics{vehicle: Truck, isDone: false}
})

var ShipCreator = Creator(func() Logistics {
	return &SeaLogistics{vehicle: Ship, isDone: false}
})

var creatorMap = map[int]Creator{
	Road: TruckCreator,
	Sea:  ShipCreator,
}

type Logistics interface {
	Deliver()
	IsDone() bool
	GetVehicle() Vehicle
}

// NewTransport is a factory method which creates either a Truck or a Ship
func NewLogistics(cargo int) (Logistics, error) {
	creator, ok := creatorMap[cargo]
	if !ok {
		return nil, errors.New("invalid cargo type")
	}
	return creator(), nil
}

type RoadLogistics struct {
	vehicle Vehicle
	isDone  bool
}

func (r *RoadLogistics) Deliver() {
	r.isDone = true
}

func (r *RoadLogistics) IsDone() bool {
	return r.isDone
}

func (r *RoadLogistics) GetVehicle() Vehicle {
	return r.vehicle
}

type SeaLogistics struct {
	vehicle Vehicle
	isDone  bool
}

func (s *SeaLogistics) Deliver() {
	s.isDone = true
}

func (s *SeaLogistics) IsDone() bool {
	return s.isDone
}

func (s *SeaLogistics) GetVehicle() Vehicle {
	return s.vehicle
}
