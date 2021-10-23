package baiter

import (
	"Def/cmp"
	"Def/global"
	"Def/types"
	"Def/util"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type BaiterHunt struct {
	Name types.StateType
}

func NewBaiterSearch() *BaiterHunt {
	return &BaiterHunt{
		Name: types.BaiterHunt,
	}
}

func (s *BaiterHunt) GetName() types.StateType {
	return s.Name
}

func (s *BaiterHunt) Enter(ai *cmp.AI, e types.IEntity) {
	ai.Counter = rand.Intn(20) + 20
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.Y = global.ScreenHeight / 2
	pc.DX = 20
	target, err := e.GetEngine().GetActiveEntityOfClass(types.Lander)
	if err == nil {
		ai.TargetId = target.GetID()
	}
}

func (s *BaiterHunt) Update(ai *cmp.AI, e types.IEntity) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	// hunt or find a target
	te := e.GetEngine().GetEntity(ai.TargetId)
	tpos := te.GetComponent(types.Pos).(*cmp.Pos)

	if tpos.X < pc.X && pc.DX > -20 {
		pc.DX -= 1
	}
	if tpos.X > pc.X && pc.DX < 20 {
		pc.DX += 1
	}

	if !te.Active() {
		target, err := e.GetEngine().GetActiveEntityOfClass(types.Lander)
		if err == nil {
			ai.TargetId = target.GetID()
		}
	} else {
		ai.Counter--
		if ai.Counter < 0 {
			ai.Counter = rand.Intn(20) + 20
			pc.DY = util.RandChoiceF([]float64{-3, 0, 3})

			//dx, dy := util.ComputeBullet(pc, tpos, 60)
			//if math.Abs(dx) < 6 {
			//ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
			//event.NotifyEvent(ev)
			//}
		}
	}

	if pc.Y < global.ScreenTop || pc.Y > global.ScreenHeight-100 {
		pc.DY = -pc.DY
	}

	pc.X += pc.DX
	pc.Y += pc.DY

	//global.CameraX = pc.X - global.ScreenWidth/2

}
