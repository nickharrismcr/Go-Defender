package bomber

import (
	"Def/cmp"
	"Def/event"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type BomberDie struct {
	Name types.StateType
}

func NewBomberDie() *BomberDie {
	return &BomberDie{
		Name: types.BomberDie,
	}
}

func (s *BomberDie) GetName() types.StateType {
	return s.Name
}

func (s *BomberDie) Enter(ai *cmp.AI, e types.IEntity) {

	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 0
	ev := event.NewBomberDie(e)
	event.NotifyEvent(ev)
	e.RemoveComponent(types.Collide)
}

func (s *BomberDie) Update(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse += 7
	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}
