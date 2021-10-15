package tests

import (
	"FSM/cmp"
	"FSM/game"
	"FSM/logger"
	"FSM/states"
	"FSM/systems"
	"FSM/testlog"
	"testing"
)

func TestEngine(t *testing.T) {

	testlog.Init()

	engine := game.NewEngine()
	engine.AddSystem(systems.NewAISystem(true), game.UPDATE)

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
	testEntity.SetActive(true)

	testAICmp = cmp.NewAI(testfsm, "teststate1")
	testEntity = game.NewEntity(engine)
	testEntity.AddComponent(testAICmp)
	testEntity.SetActive(true)

	testAICmp = cmp.NewAI(testfsm, "teststate1")
	testEntity = game.NewEntity(engine)
	testEntity.AddComponent(testAICmp)
	testEntity.SetActive(true)

	for i := 0; i < 4; i++ {
		engine.Update(0)
	}

	logs := len(testlog.Get())
	if logs != 24 {
		t.Errorf("State transition count %d wrong", logs)
	}
	if testlog.CountPattern("teststate1") != 12 {
		t.Errorf("State teststate1 count %d wrong", logs)
	}
	if testlog.CountPattern("teststate2") != 12 {
		t.Errorf("State teststate2 count %d wrong", logs)
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

	for i := 0; i < 4; i++ {
		engine.Update(0)
	}

	logs = len(testlog.Get())
	if logs != 16 {
		t.Errorf("State transition count %d wrong", logs)
	}

	if testlog.CountPattern("teststate1") != 8 {
		t.Errorf("State teststate1 count %d wrong", logs)
	}
	if testlog.CountPattern("teststate2") != 8 {
		t.Errorf("State teststate2 count %d wrong", logs)
	}

	testEntity.AddComponent(testAICmp)
	x = engine.GetEntitiesWithComponent(cmp.AIType)
	if len(x) != 3 {
		t.Errorf("Entities with AI component %d != %d", len(x), 3)
	}

	testEntity.SetActive(false)
	e := engine.GetEntity(0)
	e.SetActive(false)

	testlog.Init()

	for i := 0; i < 4; i++ {
		engine.Update(0)
	}
	logs = len(testlog.Get())
	if logs != 8 {
		t.Errorf("State transition count %d wrong", logs)
	}

	testlog.Init()
	engine.SetSystemActive(game.AISystem, false)
	engine.Update(0)
	logs = len(testlog.Get())
	if logs != 0 {
		t.Errorf("State transition count %d wrong", logs)
	}

	logger.Close()
}
