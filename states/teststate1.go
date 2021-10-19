package states

import (
	"Def/cmp"
	"Def/constants"
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

func (s *teststate1) Enter(ai *cmp.AICmp, e cmp.EntityGetter) {

}

func (s *teststate1) Update(ai *cmp.AICmp, e cmp.EntityGetter) {

	pc := e.GetComponent(cmp.PosType).(*cmp.PosCmp)
	pc.X += pc.DX
	pc.Y += pc.DY
	if pc.X < 0 || pc.X > constants.ScreenWidth {
		pc.DX = -pc.DX
	}
	if pc.Y < 0 || pc.Y > constants.ScreenHeight {
		pc.DY = -pc.DY
	}
}
