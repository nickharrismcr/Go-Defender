package lander

import (
	"Def/cmp"
	"Def/global"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type LanderGrab struct {
	Name types.StateType
}

func NewLanderGrab() *LanderGrab {
	return &LanderGrab{
		Name: types.LanderGrab,
	}
}

func (s *LanderGrab) GetName() types.StateType {
	return s.Name
}

func (s *LanderGrab) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DY = -global.LanderSpeed
	ai.Counter = 0
}

func (s *LanderGrab) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	if pc.Y < global.ScreenTop+50 {
		ai.NextState = types.LanderMutate
		te := e.GetEngine().GetEntity(e.Child())
		tai := te.GetComponent(types.AI).(*cmp.AI)
		tai.NextState = types.HumanDie
	}
}
