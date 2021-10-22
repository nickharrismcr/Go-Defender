package lander

import (
	"Def/cmp"
	"Def/constants"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type LanderSearch struct {
	Name types.StateType
}

func NewLanderSearch() *LanderSearch {
	return &LanderSearch{
		Name: types.LanderSearch,
	}
}

func (s *LanderSearch) GetName() types.StateType {
	return s.Name
}

func (s *LanderSearch) Enter(ai *cmp.AI, e types.IEntity) {

}

func (s *LanderSearch) Update(ai *cmp.AI, e types.IEntity) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	pc.X += pc.DX
	pc.Y += pc.DY
	if pc.X < 0 || pc.X > constants.ScreenWidth {
		pc.DX = -pc.DX
	}
	if pc.Y < 0 || pc.Y > constants.ScreenHeight {
		pc.DY = -pc.DY
	}
}
