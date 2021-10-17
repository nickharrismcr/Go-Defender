package states

import (
	"Def/cmp"
	"Def/game"
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

func (s *teststate1) Enter(ai *cmp.AICmp, e *game.Entity) {

	game.NotifyEvent(game.EnterState1Event, e)

}

func (s *teststate1) Update(ai *cmp.AICmp, e *game.Entity) {

	ai.NextStateName = "teststate2"
	game.NotifyEvent(game.UpdateState1Event, e)
}
