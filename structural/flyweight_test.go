package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlyweight(t *testing.T) {
	snipers := generateSnipers()
	captains := generateCaptains()

	p := snipers[0]
	for i := range snipers {
		assert.Equal(t, snipers[i].dress, p.dress)
		p = snipers[i]
	}

	p = captains[0]
	for i := range captains {
		assert.Equal(t, captains[i].dress, p.dress)
		p = captains[i]
	}
}

func generateSnipers() []Player {
	ps := make([]Player, 10)
	for i := range ps {
		p, _ := NewPlayer("sniper", SniperDressType)
		ps[i] = *p
	}
	return ps
}

func generateCaptains() []Player {
	ps := make([]Player, 10)
	for i := range ps {
		p, _ := NewPlayer("captain:%d", CaptainDressType)
		ps[i] = *p
	}
	return ps
}
