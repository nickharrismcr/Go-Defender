package swarmer

import (
	"Def/cmp"
	"Def/event"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type SwarmerDie struct {
	Name types.StateType
}

func NewSwarmerDie() *SwarmerDie {
	return &SwarmerDie{
		Name: types.SwarmerDie,
	}
}

func (s *SwarmerDie) GetName() types.StateType {
	return s.Name
}

func (s *SwarmerDie) Enter(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 0
	ev := event.NewSwarmerDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = true
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
	e.RemoveComponent(types.Collide)
}

func (s *SwarmerDie) Update(ai *cmp.AI, e types.IEntity) {

	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse += 7
	if dc.Disperse > 300 {
		e.SetActive(false)
	}

}
