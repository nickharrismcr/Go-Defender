package lander

import (
	"Def/cmp"
	"Def/event"
	"Def/global"
	"Def/types"
	"Def/util"
	"math"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type LanderSearch struct {
	Name types.StateType
}

func NewLanderSearch() *LanderSearch {
	return &LanderSearch{
		Name: types.LanderSearch,
	}
}

func (s *LanderSearch) GetName() types.StateType {
	return s.Name
}

func (s *LanderSearch) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = global.LanderSpeed
	pc.DY = 0
	ai.Counter = 0
	for _, id := range e.GetEngine().GetActiveEntitiesOfClass(types.Human) {
		te := e.GetEngine().GetEntity(id)
		if te.Parent() == te.GetID() {
			e.SetChild(te.GetID())
			te.SetParent(e.GetID())
			break
		}
	}
}

func (s *LanderSearch) Update(ai *cmp.AI, e types.IEntity) {

	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	pc.X += pc.DX
	pc.Y += pc.DY

	if ai.Counter > 5 {
		ai.Counter = 0
		mh := e.GetEngine().MountainHeight(pc.X)
		if pc.Y+200 < global.ScreenHeight-mh {
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

	// TODO check player is in range
	if rand.Intn(200) == 0 {
		tc := cmp.NewPos(pc.X+400, pc.Y, 1, 1)
		dx, dy := util.ComputeBullet(pc, tc, 60)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

	if e.Child() != e.GetID() {
		te := e.GetEngine().GetEntity(e.Child())
		tpc := te.GetComponent(types.Pos).(*cmp.Pos)
		if math.Abs(tpc.X-pc.X) < 3 {
			ai.NextState = types.LanderDrop
		}
	}

}
