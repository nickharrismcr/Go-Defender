package main

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/global"
	"Def/graphics"
	"Def/state/bomber"
	"Def/state/human"
	"Def/state/lander"
	"Def/state/player"
	"Def/systems"
	"Def/types"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var landerCount int
var humanCount int
var blankImg *ebiten.Image
var ScoreId int

// game setup

func InitGame(engine *game.Engine) {

	graphics.Load()
	InitEvents(engine)
	InitSystems(engine)
	InitEntities(engine)
	bulletPool(engine)
	bombPool(engine)

	ScoreId = engine.AddString("       0", 100, 40)
}

func InitEvents(engine *game.Engine) {
	// Events

	start := func(e event.IEvent) {
		engine.GetSystem(game.PosSystem).SetActive(true)
	}

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

	bomberDie := func(e event.IEvent) {

	}

	playerDie := func(e event.IEvent) {
		pe := engine.GetEntities()[global.PlayerID]
		pai := pe.GetComponent(types.AI).(*cmp.AI)
		pai.NextState = types.PlayerDie
		engine.GetSystem(game.PosSystem).SetActive(false)
	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)
	event.AddEventListener(event.LanderDieEvent, landerDie)
	event.AddEventListener(event.LanderClearedEvent, landerCleared)
	event.AddEventListener(event.HumanDieEvent, humanDie)
	event.AddEventListener(event.BomberDieEvent, bomberDie)
	event.AddEventListener(event.PlayerDieEvent, playerDie)
	event.AddEventListener(event.StartEvent, start)

}

func InitSystems(engine *game.Engine) {

	engine.AddSystem(systems.NewPosSystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewAISystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewLifeSystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewCollideSystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewDrawSystem(true, engine), game.DRAW)
	engine.AddSystem(systems.NewRadarDrawSystem(true, engine), game.DRAW)
}

func InitEntities(engine *game.Engine) {

	blankImg = ebiten.NewImage(5, 5)
	blankImg.Fill(color.White)

	AddPlayer(engine)

	for i := 0; i < global.BomberCount; i++ {

		AddLander(engine, i)
		landerCount++
	}

	for i := 0; i < 1; i++ {
		//TODO baiter
	}

	for i := 0; i < global.HumanCount; i++ {

		AddHuman(engine, i)
		humanCount++
	}
	for i := 0; i < global.BomberCount; i++ {

		AddBomber(engine, i)
	}
}

func AddPlayer(engine *game.Engine) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Player)
	global.PlayerID = ent.Id
	ent.SetActive(true)

	x := float64(global.WorldWidth) / 2
	y := float64(global.ScreenHeight) / 2

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sc := cmp.NewShip()
	ent.AddComponent(sc)

	stree := game.NewStateTree()
	stree.AddState(player.NewPlayerPlay())
	stree.AddState(player.NewPlayerDie())

	fsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(fsm, types.PlayerPlay)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("ship.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 1, G: 1, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)

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

	fsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(fsm, types.LanderMaterialise)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("lander.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 0, G: 1, B: 0, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
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

	fsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(fsm, types.HumanWalking)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("human.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 1, G: 0, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)

}

func AddBomber(engine *game.Engine, count int) {

	ent := game.NewEntity(engine, types.Bomber)
	ent.SetActive(true)

	x := (rand.Float64() * global.ScreenWidth) + global.WorldWidth/3
	y := (rand.Float64() * global.ScreenHeight / 2) + global.ScreenTop + 50

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(bomber.NewBomberMove())
	stree.AddState(bomber.NewBomberDie())

	fsm := game.NewFSM(stree, "fsm1")
	ai := cmp.NewAI(fsm, types.BomberMove)
	ent.AddComponent(ai)
	smap := graphics.GFXFrame{
		Frame:           graphics.SourceFrame{X: 0, Y: 0, W: 1, H: 1},
		Anim_frames:     1,
		Ticks_per_frame: 30,
	}
	dr := cmp.NewDraw(blankImg, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Cycle = true
	dr.Bomber = true
	dr.Scale = 20
	ent.AddComponent(dr)

	col := types.ColorF{R: 0.5, G: 0, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := cmp.NewCollide()
	ent.AddComponent(cl)
}

func bulletPool(engine *game.Engine) {

	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 40; i++ {
		ent := game.NewEntity(engine, types.Bullet)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bullet.png")
		dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		ent.AddComponent(dr)
		li := cmp.NewLife(240)
		ent.AddComponent(li)
		cl := cmp.NewCollide()
		ent.AddComponent(cl)
		engine.BulletPool = append(engine.BulletPool, ent)
	}

}

func bombPool(engine *game.Engine) {

	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 20; i++ {
		ent := game.NewEntity(engine, types.Bomb)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bomb.png")
		dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		dr.Cycle = true
		ent.AddComponent(dr)
		cl := cmp.NewCollide()
		ent.AddComponent(cl)
		li := cmp.NewLife(240)
		ent.AddComponent(li)
		engine.BombPool = append(engine.BombPool, ent)
	}

}
