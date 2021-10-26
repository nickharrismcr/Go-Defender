package main

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/global"
	"Def/graphics"
	"Def/state/human"
	"Def/state/lander"
	"Def/systems"
	"Def/types"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var landerCount int
var humanCount int
var radarImg *ebiten.Image
var ScoreId int

// game setup

func InitGame(engine *game.Engine) {

	graphics.Load()
	InitEvents(engine)
	InitSystems(engine)
	InitEntities(engine)
	bulletPool(engine)

	ScoreId = engine.AddString("       0", 100, 40)
}

func bulletPool(engine *game.Engine) {

	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 40; i++ {
		ent := game.NewEntity(engine, types.Bullet)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bullet.png")
		dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		//cl := cmp.NewCollide()
		li := cmp.NewLife(120)
		ent.AddComponent(dr)
		//ent.AddComponent(cl)
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
	landerDie := func(e event.IEvent) {
		ent := e.GetPayload().(*game.Entity)
		landerCount--
		if landerCount == 0 {
			lc := event.NewLanderCleared(ent)
			event.NotifyEvent(lc)
		}
	}
	landerCleared := func(e event.IEvent) {
		if ent := e.GetPayload().(*game.Entity); ent != nil {
			// end of level
		}
	}
	humanDie := func(e event.IEvent) {

	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)
	event.AddEventListener(event.LanderDieEvent, landerDie)
	event.AddEventListener(event.LanderClearedEvent, landerCleared)
	event.AddEventListener(event.HumanDieEvent, humanDie)
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

	radarImg = ebiten.NewImage(5, 5)
	radarImg.Fill(color.White)

	for i := 0; i < 20; i++ {

		AddLander(engine, i)
		landerCount++
	}
	//for i := 0; i < 1; i++ {
	//	ls := baiter.NewBaiterSearch()
	//	Add(types.Baiter, engine, ls, "baiter.png", false, 5)
	//}
	for i := 0; i < 20; i++ {

		AddHuman(engine, i)
		humanCount++
	}
}

func AddLander(engine *game.Engine, count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Lander)
	ent.SetActive(true)

	x := rand.Float64() * global.WorldWidth
	if count < 2 {
		x = rand.Float64()*global.ScreenWidth + engine.CameraX
	}
	pc := cmp.NewPos(x, global.ScreenTop+500*rand.Float64(), 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(lander.NewLanderSearch())
	stree.AddState(lander.NewLanderMaterialise())
	stree.AddState(lander.NewLanderDrop())
	stree.AddState(lander.NewLanderGrab())
	stree.AddState(lander.NewLanderMutate())
	stree.AddState(lander.NewLanderDie())

	testfsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(testfsm, types.LanderMaterialise)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("lander.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 0, G: 1, B: 0, A: 1}
	rd := cmp.NewRadarDraw(radarImg, col)
	ent.AddComponent(rd)
	cl := cmp.NewCollide()
	ent.AddComponent(cl)
}

func AddHuman(engine *game.Engine, count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Human)
	ent.SetActive(true)

	x := rand.Float64() * global.WorldWidth
	if count < 2 {
		x = rand.Float64()*global.ScreenWidth + engine.CameraX
	}
	pc := cmp.NewPos(x, 0, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(human.NewHumanWalking())
	stree.AddState(human.NewHumanGrabbed())
	stree.AddState(human.NewHumanDropping())
	stree.AddState(human.NewHumanRescued())

	stree.AddState(human.NewHumanDie())

	testfsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(testfsm, types.HumanWalking)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("human.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 1, G: 0, B: 1, A: 1}
	rd := cmp.NewRadarDraw(radarImg, col)
	ent.AddComponent(rd)
	cl := cmp.NewCollide()
	ent.AddComponent(cl)
}
