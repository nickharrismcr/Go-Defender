package lander

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on passed components

type LanderMaterialise struct {
	Name types.StateType
}

func NewLanderMaterialise() *LanderMaterialise {
	return &LanderMaterialise{
		Name: types.LanderMaterialise,
	}
}

func (s *LanderMaterialise) GetName() types.StateType {
	return s.Name
}

func (s *LanderMaterialise) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Hide = false
	dc.Disperse = 300
}

func (s *LanderMaterialise) Update(ai *cmp.AI, e types.IEntity) {

	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse -= 5
	if dc.Disperse < 10 {
		dc.Disperse = 0
		ai.NextState = types.LanderSearch
	}
}
