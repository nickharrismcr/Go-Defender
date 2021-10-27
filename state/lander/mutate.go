package lander

import (
	"Def/cmp"
	"Def/event"
	"Def/global"
	"Def/graphics"
	"Def/types"
	"Def/util"
	"math/rand"
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
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	ai.Counter++
	if ai.Counter > 2 {
		ai.Counter = 0

		pc.DY = util.RandChoiceF([]float64{-gs, 0, gs})
		pc.X += util.RandChoiceF([]float64{-20, 0, 20})
		pc.Y += util.RandChoiceF([]float64{-20, 0, 20})
	}
	if !util.OffScreen(pc.X, pc.Y) && rand.Intn(50) == 0 {
		tc := cmp.NewPos(pc.X+400, pc.Y, 1, 1)
		dx, dy := util.ComputeBullet(pc, tc, 60)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

}
