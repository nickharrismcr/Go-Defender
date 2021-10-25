package human

import (
	"Def/cmp"
	"Def/global"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type HumanWalking struct {
	Name types.StateType
}

func NewHumanWalking() *HumanWalking {
	return &HumanWalking{
		Name: types.HumanWalking,
	}
}

func (s *HumanWalking) GetName() types.StateType {
	return s.Name
}

func (s *HumanWalking) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = global.HumanSpeed
	pc.DY = 0
	ai.Counter = 0
	pc.Y = global.ScreenHeight - e.GetEngine().MountainHeight(pc.X)
}

func (s *HumanWalking) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	pc.X += pc.DX
	pc.Y = global.ScreenHeight - e.GetEngine().MountainHeight(pc.X)

}
