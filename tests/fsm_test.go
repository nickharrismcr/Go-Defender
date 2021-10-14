package tests

import (
	"FSM/cmp"
	"FSM/game"
	"FSM/states"
	"FSM/testlog"
	"testing"
)

func TestFSM(t *testing.T) {

	testlog.Init()

	stree := game.NewStateTree()
	teststate1 := states.NewTestState1()
	teststate2 := states.NewTestState2()
	stree.AddState(teststate1)
	stree.AddState(teststate2)
	stree.AddTransition("teststate1", "teststate2")
	stree.AddTransition("teststate2", "teststate1")

	var ai_list []*cmp.AICmp

	fsmId := game.NewFSM(stree, "fsm1")
	ai_list = append(ai_list, cmp.NewAI(fsmId, "teststate1"))
	ai_list = append(ai_list, cmp.NewAI(fsmId, "teststate1"))
	ai_list = append(ai_list, cmp.NewAI(fsmId, "teststate1"))

	for i := 0; i < 100; i++ {
		for _, AICmp := range ai_list {
			game.GetFSM(AICmp.FSMId).Update(AICmp)
		}
	}

	logs := len(testlog.Get())
	if logs != 600 {
		t.Fatalf("State transition count %d != %d", logs, 800)
	}

}
