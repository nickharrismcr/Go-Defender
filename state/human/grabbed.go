package human

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type HumanGrabbed struct {
	Name types.StateType
}

func NewHumanGrabbed() *HumanGrabbed {
	return &HumanGrabbed{
		Name: types.HumanGrabbed,
	}
}

func (s *HumanGrabbed) GetName() types.StateType {
	return s.Name
}

func (s *HumanGrabbed) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
}

func (s *HumanGrabbed) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pe := e.GetEngine().GetEntity(e.Parent())
	pai := pe.GetComponent(types.AI).(*cmp.AI)
	if pai.State != types.LanderDie {
		pec := pe.GetComponent(types.Pos).(*cmp.Pos)
		pc.Y = pec.Y + 50
	} else {
		ai.NextState = types.HumanDropping
	}
}
