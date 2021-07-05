package structural

import (
	"errors"
)

const (
	CaptainDressType = "captain dress"
	SniperDressType  = "sniper dress"
)

var dressFactorySingleInstance = &DressFactory{make(map[string]*dress)}

type DressFactory struct {
	dressMap map[string]*dress
}

func (df *DressFactory) getDress(typ string) (*dress, error) {
	d, ok := df.dressMap[typ]
	if ok {
		return d, nil
	}

	if typ == CaptainDressType {
		var idress dress = NewCaptainDress()
		df.dressMap[typ] = &idress
		return &idress, nil
	}
	if typ == SniperDressType {
		var idress dress = NewSniperDress(2)
		df.dressMap[typ] = &idress
		return &idress, nil
	}

	return nil, errors.New("wrong dress type passed")
}

func GetDressFactory() *DressFactory {
	return dressFactorySingleInstance
}

type dress interface {
	getColor() string
}

type CaptainDress struct {
	color string
}

func NewCaptainDress() *CaptainDress {
	return &CaptainDress{color: "green"}
}

func (cd *CaptainDress) getColor() string {
	return cd.color
}

type SniperDress struct {
	color string
	grade int
}

func NewSniperDress(grade int) *SniperDress {
	return &SniperDress{color: "blue", grade: grade}
}

func (sd *SniperDress) getColor() string {
	return sd.color
}

type Player struct {
	playerType string
	dress      dress
	power      int
	accuracy   int
}

func NewPlayer(playerType, dressType string) (*Player, error) {
	d, err := GetDressFactory().getDress(dressType)
	if err != nil {
		return nil, err
	}
	p := &Player{
		dress:      *d,
		playerType: playerType,
		power:      100,
		accuracy:   100,
	}
	return p, nil
}
