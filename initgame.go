package main

import (
	"Def/cmp"
	"Def/constants"
	"Def/event"
	"Def/game"
	"Def/graphics"
	"Def/state"
	"Def/state/baiter"
	"Def/state/lander"
	"Def/systems"
	"Def/types"
	"math/rand"
)

// game setup

func InitGame(engine *game.Engine) {

	graphics.Load()

	explodeTrigger := func(e event.IEvent) {
		if ct := e.GetPayload().(*cmp.Pos); ct != nil {
			engine.TriggerPS(ct.X, ct.Y)
		}
	}
	bulletTrigger := func(e event.IEvent) {
		if ct := e.GetPayload().(*cmp.Pos); ct != nil {
			engine.TriggerBullet(ct.X, ct.Y, ct.DX, ct.DY)
		}
	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)

	engine.AddSystem(systems.NewPosSystem(true), game.UPDATE)
	engine.AddSystem(systems.NewAISystem(true), game.UPDATE)
	engine.AddSystem(systems.NewLifeSystem(true), game.UPDATE)
	engine.AddSystem(systems.NewCollideSystem(true), game.UPDATE)
	engine.AddSystem(systems.NewDrawSystem(true), game.DRAW)

	for i := 0; i < 20; i++ {
		ls := lander.NewLanderSearch()
		Add(types.Lander, engine, ls, "lander.png", true, 2)
	}
	for i := 0; i < 2; i++ {
		ls := baiter.NewBaiterSearch()
		Add(types.Baiter, engine, ls, "baiter.png", false, 5)
	}

	bulletList(engine)
}

func Add(class types.EntityType, engine *game.Engine, state state.IState, sprite string, collide bool, speed float64) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, class)
	ent.SetActive(true)
	dx := rand.Float64()*speed - speed
	dy := rand.Float64()*speed - speed
	pc := cmp.NewPos(rand.Float64()*constants.ScreenWidth, rand.Float64()*constants.ScreenHeight, dx, dy)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(state)
	testfsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(testfsm, state.GetName())
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap(sprite)
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)
	if collide {
		cl := cmp.NewCollide()
		ent.AddComponent(cl)
	}

}

func bulletList(engine *game.Engine) {

	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 40; i++ {
		ent := game.NewEntity(engine, types.Bullet)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bullet.png")
		dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		cl := cmp.NewCollide()
		li := cmp.NewLife(120)
		ent.AddComponent(dr)
		ent.AddComponent(cl)
		ent.AddComponent(li)

		engine.Bullets = append(engine.Bullets, ent)
	}

}
