package main

import (
	"Def/cmp"
	"Def/constants"
	"Def/draw_systems"
	"Def/event"
	"Def/game"
	"Def/logger"
	"Def/states"
	"Def/update_systems"
	"math/rand"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// game setup

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
	engine.AddSystem(update_systems.NewCollideSystem(true), game.UPDATE)
	engine.AddSystem(draw_systems.NewDrawSystem(true), game.DRAW)

	for i := 0; i < 20; i++ {
		ent := game.NewEntity(engine)
		ent.SetActive(true)
		dx := rand.Float64()*8 - 8
		dy := rand.Float64()*8 - 8
		pc := cmp.NewPos(rand.Float64()*constants.ScreenWidth, rand.Float64()*constants.ScreenHeight, dx, dy)
		ent.AddComponent(pc)
		stree := game.NewStateTree()
		teststate1 := states.NewTestState1()
		stree.AddState(teststate1)
		testfsm := game.NewFSM(stree, "fsm1")
		ai := cmp.NewAI(testfsm, "teststate1")
		ent.AddComponent(ai)
		img := ebiten.NewImage(20, 20)
		img.Fill(color.White)
		dr := cmp.NewDraw(img, constants.ColorF{R: 1, G: 1, B: 1})
		cl := cmp.NewCollide()
		ent.AddComponent(dr)
		ent.AddComponent(cl)
	}

}
