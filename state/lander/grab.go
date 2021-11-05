package lander

import (
	"Def/cmp"
	"Def/event"
	"Def/gl"
	"Def/types"
	"Def/util"
	"math/rand"
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
	pc.DY = -gl.LanderSpeed
	ai.Counter = 0
}

func (s *LanderGrab) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	he := e.GetEngine().GetEntity(e.Child())
	if !he.Active() {
		ai.NextState = types.LanderSearch
	}

	// TODO gl bullet rate
	if !util.OffScreen(util.ScreenX(pc.X), pc.Y) && rand.Intn(100) == 0 {
		tc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos)
		bullettime := gl.CurrentLevel().BulletTime
		dx, dy := util.ComputeBullet(pc, tc, bullettime)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

	if pc.Y < gl.ScreenTop+50 {
		ai.NextState = types.LanderMutate
		he := e.GetEngine().GetEntity(e.Child())
		hai := he.GetComponent(types.AI).(*cmp.AI)
		hai.NextState = types.HumanDie
	}
}
