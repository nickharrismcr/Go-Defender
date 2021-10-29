package human

import (
	"Def/cmp"
	"Def/global"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type HumanDropping struct {
	Name types.StateType
}

func NewHumanDropping() *HumanDropping {
	return &HumanDropping{
		Name: types.HumanDropping,
	}
}

func (s *HumanDropping) GetName() types.StateType {
	return s.Name
}

func (s *HumanDropping) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	e.SetParent(-1)
	pc.DX = 0
	pc.DY = 0
	ai.Counter = 0
}

func (s *HumanDropping) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DY += 0.1

	if pc.Y > global.ScreenHeight-e.GetEngine().MountainHeight(pc.X) {
		if pc.DY > 10 {
			ai.NextState = types.HumanDie
		} else {
			ai.NextState = types.HumanWalking
		}
	}
}
