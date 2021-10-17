package tests

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"
	"Def/states"
	"Def/update_systems"
	"fmt"
	"testing"
)

func GetEventFunction(s string) func(*game.Entity) {
	return func(e *game.Entity) {
		AddTestLog(fmt.Sprintf("Entity %d : %s ", e.Id, s))
	}
}

func TestEngine(t *testing.T) {

	logger.Debug("-------------------------------------------------------------------------------")

	InitTestLog()

	game.AddEventListener(game.EnterState1Event, GetEventFunction("Enter teststate1"))
	game.AddEventListener(game.EnterState2Event, GetEventFunction("Enter teststate2"))
	game.AddEventListener(game.UpdateState1Event, GetEventFunction("Update teststate1"))
	game.AddEventListener(game.UpdateState2Event, GetEventFunction("Update teststate2"))

	engine := game.NewEngine()
	engine.AddSystem(update_systems.NewAISystem(true), game.UPDATE)

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

	t.Run("Transition counts #1 ", func(t *testing.T) {
		logs := len(GetTestLog())
		if logs != 24 {
			t.Errorf("State transition count %d wrong", logs)
		}
		if TestLogCountPattern("teststate1") != 12 {
			t.Errorf("State teststate1 count %d wrong", logs)
		}
		if TestLogCountPattern("teststate2") != 12 {
			t.Errorf("State teststate2 count %d wrong", logs)
		}
	})

	t.Run("Entities with component #1 ", func(t *testing.T) {
		x := engine.GetEntitiesWithComponent(cmp.AIType)
		if len(x) != 3 {
			t.Errorf("Entities with AI component %d != %d", len(x), 3)
		}
	})

	t.Run("Entities with component #2 ", func(t *testing.T) {
		testEntity.RemoveComponent(cmp.AIType)
		x := engine.GetEntitiesWithComponent(cmp.AIType)
		if len(x) != 2 {
			t.Errorf("Entities with AI component %d != %d", len(x), 3)
		}
	})

	InitTestLog()

	for i := 0; i < 4; i++ {
		engine.Update(0)
	}

	t.Run("Transition counts #2 ", func(t *testing.T) {
		logs := len(GetTestLog())
		if logs != 16 {
			t.Errorf("State transition count %d wrong", logs)
		}

		if TestLogCountPattern("teststate1") != 8 {
			t.Errorf("State teststate1 count %d wrong", logs)
		}
		if TestLogCountPattern("teststate2") != 8 {
			t.Errorf("State teststate2 count %d wrong", logs)
		}
	})

	t.Run("Entities with component #2 ", func(t *testing.T) {
		testEntity.AddComponent(testAICmp)
		x := engine.GetEntitiesWithComponent(cmp.AIType)
		if len(x) != 3 {
			t.Errorf("Entities with AI component %d != %d", len(x), 3)
		}
	})

	t.Run("Transition count #3 ", func(t *testing.T) {
		testEntity.SetActive(false)
		e := engine.GetEntity(0)
		e.SetActive(false)

		InitTestLog()

		for i := 0; i < 4; i++ {
			engine.Update(0)
		}
		logs := len(GetTestLog())
		if logs != 8 {
			t.Errorf("State transition count %d wrong", logs)
		}

	})

	t.Run("Transition count #4 ", func(t *testing.T) {
		InitTestLog()
		engine.SetSystemActive(game.AISystem, false)
		engine.Update(0)
		logs := len(GetTestLog())
		if logs != 0 {
			t.Errorf("State transition count %d wrong", logs)
		}
	})
	logger.Close()
}
