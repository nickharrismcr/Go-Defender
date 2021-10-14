package states

import (
	"FSM/cmp"
	"FSM/testlog"
	"fmt"
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

func (s *teststate1) Enter(ai *cmp.AICmp) {

	testlog.Add(fmt.Sprintf("%d Enter %s %d", ai.Id, s.Name, ai.Counter))
}

func (s *teststate1) Update(ai *cmp.AICmp) {

	ai.NextStateName = "teststate2"
	testlog.Add(fmt.Sprintf("%d Update %s %d", ai.Id, s.Name, ai.Counter))
}
