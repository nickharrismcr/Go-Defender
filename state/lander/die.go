package lander

import (
	"Def/cmp"
	"Def/event"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type LanderDie struct {
	Name types.StateType
}

func NewLanderDie() *LanderDie {
	return &LanderDie{
		Name: types.LanderDie,
	}
}

func (s *LanderDie) GetName() types.StateType {
	return s.Name
}

func (s *LanderDie) Enter(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 0
	ev := event.NewLanderDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = true
}

func (s *LanderDie) Update(ai *cmp.AI, e types.IEntity) {

	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse += 7
	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}
