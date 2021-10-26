package lander

import (
	"Def/cmp"
	"Def/global"
	"Def/graphics"
	"Def/types"
	"Def/util"
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
	pc.DX = global.LanderSpeed * 2
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.SpriteMap = graphics.GetSpriteMap("mutant.png")
	rc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rc.Cycle = true
	ai.Counter = 0
}

func (s *LanderMutate) Update(ai *cmp.AI, e types.IEntity) {
	gs := float64(global.LanderSpeed)
	ai.Counter++
	if ai.Counter > 2 {
		ai.Counter = 0
		pc := e.GetComponent(types.Pos).(*cmp.Pos)
		pc.DY = util.RandChoiceF([]float64{-gs, 0, gs})
		pc.X += util.RandChoiceF([]float64{-20, 0, 20})
		pc.Y += util.RandChoiceF([]float64{-20, 0, 20})
	}

}
