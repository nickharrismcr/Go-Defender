package human

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on passed components

type HumanRescued struct {
	Name types.StateType
}

func NewHumanRescued() *HumanRescued {
	return &HumanRescued{
		Name: types.HumanRescued,
	}
}

func (s *HumanRescued) GetName() types.StateType {
	return s.Name
}

func (s *HumanRescued) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 300
}

func (s *HumanRescued) Update(ai *cmp.AI, e types.IEntity) {

	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse -= 5
	if dc.Disperse < 10 {
		dc.Disperse = 0
		ai.NextState = types.HumanWalking
	}
}
