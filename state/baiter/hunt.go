package baiter

import (
	"Def/cmp"
	"Def/global"
	"Def/types"
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

	if pc.Y < global.ScreenTop || pc.Y > global.ScreenHeight-100 {
		pc.DY = -pc.DY
	}

	pc.X += pc.DX
	pc.Y += pc.DY

}
