package states

import (
	"Def/cmp"
	"Def/constants"
	"Def/event"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type teststate1 struct {
	Name string
}

func NewTestState1() *teststate1 {
	return &teststate1{
		Name: "teststate1",
	}
}

func (s *teststate1) GetName() string {
	return s.Name
}

func (s *teststate1) Enter(ai *cmp.AICmp, e cmp.ComponentGetter) {

	ai.Counter = rand.Intn(60)

}

func (s *teststate1) Update(ai *cmp.AICmp, e cmp.ComponentGetter) {

	ai.Counter--
	if ai.Counter <= 0 {
		x := rand.Float64() * constants.ScreenWidth
		y := rand.Float64() * constants.ScreenHeight
		pc := e.GetComponent(cmp.PosType).(*cmp.PosCmp)
		pc.X = x
		pc.Y = y
		ev := event.NewExplode(e)
		event.NotifyEvent(ev)
		ai.Counter = rand.Intn(60)
	}
}
