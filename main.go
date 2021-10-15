package main

import (
	"FSM/cmp"
	"FSM/game"
	"FSM/states"
	"FSM/update_systems"
	"fmt"
)

func GetEventFunction(s string) func(*game.Entity) {
	return func(e *game.Entity) {
		fmt.Printf("Entity %d : %s \n", e.Id, s)
	}
}

func main() {

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

	for i := 0; i < 10; i++ {

		testAICmp := cmp.NewAI(testfsm, "teststate1")
		testEntity := game.NewEntity(engine)
		testEntity.AddComponent(testAICmp)
		testEntity.SetActive(true)
	}

	for i := 0; i < 20; i++ {
		engine.Update(0)
	}

}
