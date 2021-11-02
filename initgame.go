package main

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/gl"
	"Def/graphics"
	"Def/logger"
	"Def/state/bomber"
	"Def/state/human"
	"Def/state/lander"
	"Def/state/player"
	"Def/state/pod"
	"Def/state/swarmer"
	"Def/systems"
	"Def/types"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var landerCount int
var humanCount int
var blankImg *ebiten.Image
var scoreId int

// game setup

func InitGame(engine *game.Engine) {

	graphics.Load()

	initSystems(engine)
	initEntities(engine)
	bulletPool(engine)
	bombPool(engine)
	laserPool(engine)
	InitEvents(engine)

	scoreId = engine.AddString("       0", 100, 40)
}

func initSystems(engine *game.Engine) {

	engine.AddSystem(systems.NewPosSystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewAISystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewLifeSystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewCollideSystem(true, engine), game.UPDATE)
	engine.AddSystem(systems.NewDrawSystem(true, engine), game.DRAW)
	engine.AddSystem(systems.NewRadarDrawSystem(true, engine), game.DRAW)
	engine.AddSystem(systems.NewLaserDrawSystem(true, engine), game.DRAW)
	engine.AddSystem(systems.NewLaserMoveSystem(true, engine), game.UPDATE)
}

func initEntities(engine *game.Engine) {

	blankImg = ebiten.NewImage(20, 20)
	blankImg.Fill(color.White)

	AddPlayer(engine)

	for i := 0; i < gl.LanderCount; i++ {

		AddLander(engine, i)
		landerCount++
	}

	for i := 0; i < 1; i++ {
		//TODO baiter
	}

	for i := 0; i < gl.HumanCount; i++ {

		AddHuman(engine, i)
		humanCount++
	}
	for i := 0; i < gl.BomberCount; i++ {

		AddBomber(engine, i)
	}
	for i := 0; i < gl.PodCount; i++ {

		AddPod(engine, i)
	}
}

func AddPlayer(engine *game.Engine) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Player)
	gl.PlayerID = ent.Id
	ent.SetActive(true)

	x := float64(gl.WorldWidth) / 2
	y := float64(gl.ScreenHeight) / 2

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)

	stree := game.NewStateTree()
	stree.AddState(player.NewPlayerPlay())
	stree.AddState(player.NewPlayerDie())

	fsm := game.NewFSM(stree)
	ai := cmp.NewAI(fsm, types.PlayerPlay)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("ship.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 1, G: 1, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)

	sc := cmp.NewShip(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(sc)

}

func AddLander(engine *game.Engine, count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Lander)
	ent.SetActive(true)

	x := rand.Float64() * gl.WorldWidth
	if count < 2 {
		x = gl.WorldWidth * 0.8
	}
	pc := cmp.NewPos(x, gl.ScreenTop+500*rand.Float64(), 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(lander.NewLanderWait())
	stree.AddState(lander.NewLanderSearch())
	stree.AddState(lander.NewLanderMaterialise())
	stree.AddState(lander.NewLanderDrop())
	stree.AddState(lander.NewLanderGrab())
	stree.AddState(lander.NewLanderMutate())
	stree.AddState(lander.NewLanderDie())

	fsm := game.NewFSM(stree)
	ai := cmp.NewAI(fsm, types.LanderWait)
	ai.Wait = 60 + (count%3)*200
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("lander.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 0, G: 1, B: 0, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)

}

func AddHuman(engine *game.Engine, count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Human)
	ent.SetActive(true)

	x := rand.Float64() * gl.WorldWidth
	if count < 2 {
		x = rand.Float64()*gl.ScreenWidth + gl.CameraX()
	}
	pc := cmp.NewPos(x, 0, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(human.NewHumanWalking())
	stree.AddState(human.NewHumanGrabbed())
	stree.AddState(human.NewHumanDropping())
	stree.AddState(human.NewHumanRescued())
	stree.AddState(human.NewHumanDie())

	fsm := game.NewFSM(stree)
	ai := cmp.NewAI(fsm, types.HumanWalking)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("human.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)

	col := types.ColorF{R: 1, G: 0, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	sh := cmp.NewShootable()
	ent.AddComponent(sh)
	cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)

}

func AddBomber(engine *game.Engine, count int) {

	ent := game.NewEntity(engine, types.Bomber)
	ent.SetActive(true)

	x := (rand.Float64() * gl.ScreenWidth) + gl.WorldWidth/3
	y := (rand.Float64() * gl.ScreenHeight / 2) + gl.ScreenTop + 50

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(bomber.NewBomberMove())
	stree.AddState(bomber.NewBomberDie())

	fsm := game.NewFSM(stree)
	ai := cmp.NewAI(fsm, types.BomberMove)
	ent.AddComponent(ai)
	smap := graphics.GFXFrame{
		Frame:           graphics.SourceFrame{X: 0, Y: 0, W: 20, H: 20},
		Anim_frames:     1,
		Ticks_per_frame: 30,
	}
	dr := cmp.NewDraw(blankImg, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Cycle = true
	dr.Bomber = true
	dr.Scale = 1
	ent.AddComponent(dr)

	col := types.ColorF{R: 0.5, G: 0, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
	sh := cmp.NewShootable()
	ent.AddComponent(sh)
}

func AddPod(engine *game.Engine, count int) {

	ent := game.NewEntity(engine, types.Pod)
	ent.SetActive(true)

	x := (rand.Float64() * gl.ScreenWidth) + gl.WorldWidth/2
	y := (rand.Float64() * gl.ScreenHeight / 2) + gl.ScreenTop + 50

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(pod.NewPodMove())
	stree.AddState(pod.NewPodDie())

	fsm := game.NewFSM(stree)
	ai := cmp.NewAI(fsm, types.PodMove)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("pod.png")
	ssheet := graphics.GetSpriteSheet()
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Scale = 1
	ent.AddComponent(dr)

	col := types.ColorF{R: 0.5, G: 0, B: 0.5, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
	sh := cmp.NewShootable()
	ent.AddComponent(sh)
}

func AddSwarmer(engine *game.Engine, count int, x, y float64) {

	ent := game.NewEntity(engine, types.Swarmer)
	ent.SetActive(true)

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	stree := game.NewStateTree()
	stree.AddState(swarmer.NewSwarmerMove())
	stree.AddState(swarmer.NewSwarmerDie())

	fsm := game.NewFSM(stree)
	ai := cmp.NewAI(fsm, types.SwarmerMove)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("swarmer.png")
	ssheet := graphics.GetSpriteSheet()
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Scale = 1
	ent.AddComponent(dr)

	col := types.ColorF{R: 0.7, G: 0, B: 0, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
	sh := cmp.NewShootable()
	ent.AddComponent(sh)
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
		cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
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
		cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
		ent.AddComponent(cl)
		li := cmp.NewLife(240)
		ent.AddComponent(li)
		engine.BombPool = append(engine.BombPool, ent)
	}

}

func laserPool(engine *game.Engine) {

	for i := 0; i < 15; i++ {
		ent := game.NewEntity(engine, types.Laser)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		dr := cmp.NewLaserDraw()
		ent.AddComponent(dr)
		li := cmp.NewLife(240)
		ent.AddComponent(li)
		mv := cmp.NewLaserMove()
		ent.AddComponent(mv)

		engine.LaserPool = append(engine.LaserPool, ent)
	}

}

func InitEvents(engine *game.Engine) {
	// Events

	start := func(e event.IEvent) {
		engine.GetSystem(game.PosSystem).SetActive(true)
	}

	playerCollide := func(ev event.IEvent) {

		e := ev.GetPayload().(*game.Entity)
		logger.Info("Collide : %s ", e.Class.String())
		if e.Class == types.Human {
			ai := e.GetComponent(types.AI).(*cmp.AI)
			if ai.State == types.HumanDropping {
				ai.NextState = types.HumanRescued
			}
		} else {
			engine.Kill(e)
			//ev := event.NewPlayerDie(engine.GetPlayer())
			//event.NotifyEvent(ev)
		}
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
		pe := engine.GetEntities()[gl.PlayerID]
		pai := pe.GetComponent(types.AI).(*cmp.AI)
		pai.NextState = types.PlayerDie
		engine.GetSystem(game.PosSystem).SetActive(false)
	}

	playerFire := func(e event.IEvent) {
		if pe := e.GetPayload().(*game.Entity); pe != nil {
			pc := pe.GetComponent(types.Pos).(*cmp.Pos)
			sc := pe.GetComponent(types.Ship).(*cmp.Ship)
			x := pc.X + 100
			y := pc.Y + 25
			if sc.Direction < 0 {
				x = pc.X - 100
			}
			engine.TriggerLaser(x, y, sc.Direction)
		}
	}

	smartBomb := func(e event.IEvent) {
		engine.SetFlash()
		engine.SmartBomb()
	}

	podDie := func(e event.IEvent) {
		ent := e.GetPayload().(*game.Entity)
		pc := ent.GetComponent(types.Pos).(*cmp.Pos)
		for i := 0; i < gl.SwarmerCount; i++ {
			AddSwarmer(engine, i, pc.X, pc.Y)
		}
	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)
	event.AddEventListener(event.LanderDieEvent, landerDie)
	event.AddEventListener(event.LanderClearedEvent, landerCleared)
	event.AddEventListener(event.HumanDieEvent, humanDie)
	event.AddEventListener(event.BomberDieEvent, bomberDie)
	event.AddEventListener(event.PlayerDieEvent, playerDie)
	event.AddEventListener(event.StartEvent, start)
	event.AddEventListener(event.PlayerFireEvent, playerFire)
	event.AddEventListener(event.SmartBombEvent, smartBomb)
	event.AddEventListener(event.PlayerCollideEvent, playerCollide)
	event.AddEventListener(event.PodDieEvent, podDie)

}
