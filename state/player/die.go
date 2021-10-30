package player

import (
	"Def/cmp"
	"Def/graphics"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type PlayerDie struct {
	Name types.StateType
}

func NewPlayerDie() *PlayerDie {
	return &PlayerDie{
		Name: types.PlayerDie,
	}
}

func (s *PlayerDie) GetName() types.StateType {
	return s.Name
}

func (s *PlayerDie) Enter(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.SpriteMap = graphics.GetSpriteMap("shipd.png")
	dc.Frame = 0
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
	ai.Counter = 0
}

func (s *PlayerDie) Update(ai *cmp.AI, e types.IEntity) {
	ai.Counter++
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	if ai.Counter == 60 {
		dc.Hide = true
		pc := e.GetComponent(types.Pos).(*cmp.Pos)
		e.GetEngine().TriggerPS(pc.X, pc.Y)
	}
	if ai.Counter == 180 {
		ai.NextState = types.PlayerPlay
		dc.Hide = false
		dc.SpriteMap = graphics.GetSpriteMap("ship.png")
		dc.Frame = 0

	}
}