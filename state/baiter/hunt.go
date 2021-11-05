package baiter

import (
	"Def/cmp"
	"Def/event"
	"Def/gl"
	"Def/types"
	"Def/util"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type BaiterHunt struct {
	Name types.StateType
}

func NewBaiterHunt() *BaiterHunt {
	return &BaiterHunt{
		Name: types.BaiterHunt,
	}
}

func (s *BaiterHunt) GetName() types.StateType {
	return s.Name
}

func (s *BaiterHunt) Enter(ai *cmp.AI, e types.IEntity) {
	ai.Scratch = 0
}

func (s *BaiterHunt) Update(ai *cmp.AI, e types.IEntity) {

	gs := float64(gl.BaiterSpeed)
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	ple := e.GetEngine().GetPlayer()
	plpos := ple.GetComponent(types.Pos).(*cmp.Pos)

	ai.Scratch++

	if ai.Scratch == 30 {
		ai.Scratch = 0
		xoff := rand.Float64()*200 - 200
		yoff := rand.Float64()*100 - 100
		offpos := &cmp.Pos{X: plpos.X + xoff, Y: plpos.Y + yoff, DX: plpos.DX, DY: plpos.DY}
		pc.DX, pc.DY = util.ComputeBullet(pc, offpos, 1)
		pc.DX = util.Clamp(pc.DX, -gs, gs)
	}

	if !util.OffScreen(util.ScreenX(pc.X), pc.Y) && rand.Intn(50) == 0 {
		plp := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos)
		bullettime := gl.CurrentLevel().BulletTime
		dx, dy := util.ComputeBullet(pc, plp, bullettime/2)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

}
