package lander

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on passed components

type LanderWait struct {
	Name types.StateType
}

func NewLanderWait() *LanderWait {
	return &LanderWait{
		Name: types.LanderWait,
	}
}

func (s *LanderWait) GetName() types.StateType {
	return s.Name
}

func (s *LanderWait) Enter(ai *cmp.AI, e types.IEntity) {
	dr := e.GetComponent(types.Draw).(*cmp.Draw)
	dr.Hide = true
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = true
}

func (s *LanderWait) Update(ai *cmp.AI, e types.IEntity) {
	ai.Wait--
	if ai.Wait <= 0 {
		ai.NextState = types.LanderMaterialise
	}
}
