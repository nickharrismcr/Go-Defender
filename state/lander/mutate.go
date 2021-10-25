package lander

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type LanderMutate struct {
	Name types.StateType
}

func NewLanderMutate() *LanderMutate {
	return &LanderMutate{
		Name: types.LanderMutate,
	}
}

func (s *LanderMutate) GetName() types.StateType {
	return s.Name
}

func (s *LanderMutate) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DY = 0
	pc.DX = 0
}

func (s *LanderMutate) Update(ai *cmp.AI, e types.IEntity) {

}
