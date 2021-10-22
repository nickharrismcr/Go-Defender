package lander

import (
	"Def/cmp"
	"Def/global"
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

	if pc.Y < global.ScreenHeight-global.WorldHeight || pc.Y > global.ScreenHeight {
		pc.DY = -pc.DY
	}
}
