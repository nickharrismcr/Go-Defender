package states

import (
	"FSM/cmp"
	"FSM/game"
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

func (s *TestState2) Enter(ai *cmp.AICmp, e *game.Entity) {

	game.NotifyEvent(game.EnterState2Event, e)
}

func (s *TestState2) Update(ai *cmp.AICmp, e *game.Entity) {
	ai.NextStateName = "teststate1"
	game.NotifyEvent(game.UpdateState2Event, e)
}
