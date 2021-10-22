package main

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/global"
	"Def/graphics"
	"Def/logger"
	"Def/state"
	"Def/state/baiter"
	"Def/state/lander"
	"Def/systems"
	"Def/types"
	"Def/util"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var landerCount int
var radarImg *ebiten.Image

// game setup

func InitGame(engine *game.Engine) {

	graphics.Load()
	InitEvents(engine)
	InitSystems(engine)
	InitEntities(engine)
	bulletPool(engine)
}

func bulletPool(engine *game.Engine) {

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

		engine.BulletPool = append(engine.BulletPool, ent)
	}

}

func InitEvents(engine *game.Engine) {
	// Events
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
	entityDie := func(e event.IEvent) {
		if ent := e.GetPayload().(*game.Entity); ent != nil {
			logger.Debug("die %d %d ", ent.Class, landerCount)
			if ent.Class == types.Lander {
				landerCount--
				if landerCount == 0 {
					lc := event.NewLanderCleared(ent)
					event.NotifyEvent(lc)
				}
			}
		}
	}
	landerCleared := func(e event.IEvent) {
		if ent := e.GetPayload().(*game.Entity); ent != nil {
			// end of level
		}
	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)
	event.AddEventListener(event.EntityDieEvent, entityDie)
	event.AddEventListener(event.LanderClearedEvent, landerCleared)
}

func InitSystems(engine *game.Engine) {

	engine.AddSystem(systems.NewPosSystem(true), game.UPDATE)
	engine.AddSystem(systems.NewAISystem(true), game.UPDATE)
	engine.AddSystem(systems.NewLifeSystem(true), game.UPDATE)
	engine.AddSystem(systems.NewCollideSystem(true), game.UPDATE)
	engine.AddSystem(systems.NewDrawSystem(true), game.DRAW)
	engine.AddSystem(systems.NewRadarDrawSystem(true), game.DRAW)
}

func InitEntities(engine *game.Engine) {

	radarImg = ebiten.NewImage(10, 10)
	radarImg.Fill(color.White)

	for i := 0; i < 10; i++ {
		ls := lander.NewLanderSearch()
		Add(types.Lander, engine, ls, "lander.png", true, 2)
		landerCount++
	}
	for i := 0; i < 1; i++ {
		ls := baiter.NewBaiterSearch()
		Add(types.Baiter, engine, ls, "baiter.png", false, 5)
	}
}

func Add(class types.EntityType, engine *game.Engine, state state.IState, sprite string, collide bool, speed float64) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, class)
	ent.SetActive(true)
	dx := speed * util.RandChoiceF([]float64{-1, 1})
	dy := speed * util.RandChoiceF([]float64{-1, 1})

	pc := cmp.NewPos(rand.Float64()*global.WorldWidth, global.ScreenHeight-global.WorldHeight*rand.Float64(), dx, dy)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(state)
	testfsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(testfsm, state.GetName())
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap(sprite)
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	var col types.ColorF
	switch class {
	case types.Baiter:
		col = types.ColorF{R: 1, G: 1, B: 0, A: 1}
	case types.Lander:
		col = types.ColorF{R: 0, G: 1, B: 0, A: 1}
	}
	rd := cmp.NewRadarDraw(radarImg, col)
	ent.AddComponent(rd)

	if collide {
		cl := cmp.NewCollide()
		ent.AddComponent(cl)
	}

}
