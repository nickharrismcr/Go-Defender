package baiter

import (
	"Def/cmp"
	"Def/gl"
	"Def/types"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type BaiterWait struct {
	Name types.StateType
}

func NewBaiterWait() *BaiterWait {
	return &BaiterWait{
		Name: types.BaiterWait,
	}
}

func (s *BaiterWait) GetName() types.StateType {
	return s.Name
}

func (s *BaiterWait) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.Y = 9999
}

func (s *BaiterWait) Update(ai *cmp.AI, e types.IEntity) {

	if gl.CurrentLevel().LanderCount-gl.LandersKilled < 3 {
		ai.NextState = types.BaiterMaterialise
		pc := e.GetComponent(types.Pos).(*cmp.Pos)
		pc.Y = gl.ScreenHeight / 2
		pc.X = gl.CameraX() + rand.Float64()*gl.ScreenWidth
		pc.DX = e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos).DX
	}
}
