package tests

import (
	"FSM/cmp"
	"FSM/game"
	"FSM/states"
	"FSM/systems"
	"FSM/testlog"
	"testing"
)

func TestEngine(t *testing.T) {

	testlog.Init()

	engine := game.NewEngine()
	engine.AddSystem(systems.NewAISystem(), game.UPDATE)

	stree := game.NewStateTree()
	teststate1 := states.NewTestState1()
	teststate2 := states.NewTestState2()
	stree.AddState(teststate1)
	stree.AddState(teststate2)
	stree.AddTransition("teststate1", "teststate2")
	stree.AddTransition("teststate2", "teststate1")

	testfsm := game.NewFSM(stree, "fsm1")

	testAICmp := cmp.NewAI(testfsm, "teststate1")
	testEntity := game.NewEntity(engine)
	testEntity.AddComponent(testAICmp)

	testAICmp = cmp.NewAI(testfsm, "teststate1")
	testEntity = game.NewEntity(engine)
	testEntity.AddComponent(testAICmp)

	testAICmp = cmp.NewAI(testfsm, "teststate1")
	testEntity = game.NewEntity(engine)
	testEntity.AddComponent(testAICmp)

	for i := 0; i < 100; i++ {
		engine.Update(0)
	}

	logs := len(testlog.Get())
	if logs != 600 {
		t.Fatalf("State transition count %d wrong", logs)
	}

	x := engine.GetEntitiesWithComponent(cmp.AIType)
	if len(x) != 3 {
		t.Errorf("Entities with AI component %d != %d", len(x), 3)
	}

	testEntity.RemoveComponent(cmp.AIType)
	x = engine.GetEntitiesWithComponent(cmp.AIType)
	if len(x) != 2 {
		t.Errorf("Entities with AI component %d != %d", len(x), 3)
	}

	testlog.Init()

	for i := 0; i < 100; i++ {
		engine.Update(0)
	}

	logs = len(testlog.Get())
	if logs != 400 {
		t.Fatalf("State transition count %d wrong", logs)
	}

	testEntity.AddComponent(testAICmp)
	x = engine.GetEntitiesWithComponent(cmp.AIType)
	if len(x) != 3 {
		t.Errorf("Entities with AI component %d != %d", len(x), 3)
	}
}
