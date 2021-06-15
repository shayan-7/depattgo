package creational

type Engine int

const (
	SportEngine Engine = iota
	SUVEngine
)

const (
	SportSeats = 2
	SUVSeats   = 4
)

type IBuilder interface {
	Reset()
	SetSeats()
	SetEngine()
	SetTrip()
	SetTurbo()
	GetCar() *Car
}

type SportCarBuilder struct {
	*Car
}

func NewSportCarBuilder() *SportCarBuilder {
	return &SportCarBuilder{NewCar()}
}

func (scb *SportCarBuilder) Reset() {
	scb.Car = &Car{}
}

func (scb *SportCarBuilder) SetSeats() {
	scb.seats = SportSeats
}

func (scb *SportCarBuilder) SetTrip() {
	scb.trip = false
}

func (scb *SportCarBuilder) SetTurbo() {
	scb.turbo = true
}

func (scb *SportCarBuilder) SetEngine() {
	scb.engine = SportEngine
}

func (scb *SportCarBuilder) GetCar() *Car {
	return scb.Car
}

type SUVCarBuilder struct {
	*Car
}

func NewSUVCarBuilder() *SUVCarBuilder {
	return &SUVCarBuilder{NewCar()}
}

func (scb *SUVCarBuilder) Reset() {
	scb.Car = &Car{}
}

func (scb *SUVCarBuilder) SetSeats() {
	scb.seats = SUVSeats
}

func (scb *SUVCarBuilder) SetTrip() {
	scb.trip = true
}

func (scb *SUVCarBuilder) SetTurbo() {
	scb.turbo = false
}

func (scb *SUVCarBuilder) SetEngine() {
	scb.engine = SUVEngine
}

func (scb *SUVCarBuilder) GetCar() *Car {
	return scb.Car
}

type Car struct {
	seats  int
	engine Engine
	trip   bool
	turbo  bool
}

func NewCar() *Car {
	return &Car{}
}

type Director struct {
	IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{b}
}

func (d *Director) Construct() *Car {
	d.Reset()
	d.SetSeats()
	d.SetEngine()
	d.SetTrip()
	d.SetTurbo()
	return d.GetCar()
}

func (d *Director) SetBuilder(b IBuilder) {
	d.IBuilder = b
}
