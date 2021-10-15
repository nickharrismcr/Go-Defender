package states

import (
	"FSM/cmp"
	"FSM/testlog"
	"fmt"
)

type TestState2 struct {
	Name    string
	Counter int
}

func NewTestState2() *TestState2 {
	return &TestState2{
		Name: "teststate2",
	}
}

func (s *TestState2) GetName() string {
	return s.Name
}

func (s *TestState2) Enter(ai *cmp.AICmp) {

	testlog.Add(fmt.Sprintf("%d Enter %s", ai.Id, s.Name))
}

func (s *TestState2) Update(ai *cmp.AICmp) {
	ai.NextStateName = "teststate1"
	testlog.Add(fmt.Sprintf("%d Update %s", ai.Id, s.Name))
}
