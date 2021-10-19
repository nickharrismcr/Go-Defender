package main

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/logger"
	"Def/states"
	"Def/update_systems"
)

//import "Def/draw_systems"

func InitGame(engine *game.Engine) {

	f := func(e event.IEvent) {

		ct := e.GetPayload().(*game.Entity).GetComponent(cmp.PosType)
		if ct != nil {
			c := ct.(*cmp.PosCmp)
			logger.Debug("%T : %f %f ", c, c.X, c.Y)
			engine.TriggerPS(c.X, c.Y)
		}
	}
	event.AddEventListener(event.ExplodeEvent, f)

	engine.AddSystem(update_systems.NewPosSystem(true), game.UPDATE)
	engine.AddSystem(update_systems.NewAISystem(true), game.UPDATE)
	//engine.AddSystem(draw_systems.NewDrawSystem(true), DRAW)

	ent := game.NewEntity(engine)
	ent.SetActive(true)
	pc := cmp.NewPos(100, 100, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	teststate1 := states.NewTestState1()
	stree.AddState(teststate1)
	testfsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(testfsm, "teststate1")
	ent.AddComponent(ai)

}
