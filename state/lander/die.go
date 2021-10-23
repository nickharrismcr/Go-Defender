package lander

import (
	"Def/cmp"
	"Def/global"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type LanderDie struct {
	Name types.StateType
}

func NewLanderDie() *LanderDie {
	return &LanderDie{
		Name: types.LanderDie,
	}
}

func (s *LanderDie) GetName() types.StateType {
	return s.Name
}

func (s *LanderDie) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = global.LanderSpeed
	pc.DY = 0
	ai.Counter = 0
}

func (s *LanderDie) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	pc.X += pc.DX
	pc.Y += pc.DY

	if ai.Counter > 5 {
		ai.Counter = 0
		mh := e.GetEngine().MountainHeight(pc.X)
		if pc.Y+300 < global.ScreenHeight-mh {
			ai.Scratch++
		} else {
			ai.Scratch--
		}
	}
	if ai.Scratch < 0 {
		ai.Scratch = 0
	}
	if ai.Scratch > 5 {
		ai.Scratch = 5
	}
	switch ai.Scratch {
	case 0:
		pc.DY = -global.LanderSpeed
	case 1, 2, 3, 4:
		pc.DY = 0
	case 5:
		pc.DY = global.LanderSpeed
	}
}
