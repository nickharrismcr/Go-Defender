package game

import (
	"Def/cmp"
	"Def/constants"
	"Def/event"
	"Def/types"
	"Def/util"
	"math"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type BaiterSearch struct {
	Name types.StateType
}

func NewBaiterSearch() *BaiterSearch {
	return &BaiterSearch{
		Name: types.BaiterSearch,
	}
}

func (s *BaiterSearch) GetName() types.StateType {
	return s.Name
}

func (s *BaiterSearch) Enter(ai *cmp.AI, e types.IEntity) {
	ai.Counter = rand.Intn(10) + 10
}

func (s *BaiterSearch) Update(ai *cmp.AI, e types.IEntity) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	ai.Counter--
	if ai.Counter < 0 {
		ai.Counter = rand.Intn(10) + 10
		pc.DX = util.RandChoiceF([]float64{-3, 0, 3})
		pc.DY = util.RandChoiceF([]float64{-3, 0, 3})

		dir := rand.Float64() * math.Pi * 2
		dx := 3 * math.Sin(dir)
		dy := 3 * math.Cos(dir)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

	pc.X += pc.DX
	pc.Y += pc.DY
	if pc.X < 0 || pc.X > constants.ScreenWidth {
		pc.DX = -pc.DX
	}
	if pc.Y < 0 || pc.Y > constants.ScreenHeight {
		pc.DY = -pc.DY
	}
}
