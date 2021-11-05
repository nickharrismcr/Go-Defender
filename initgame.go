package main

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/gl"
	"Def/graphics"
	"Def/logger"
	"Def/sound"
	"Def/state/baiter"
	"Def/state/bomber"
	"Def/state/human"
	"Def/state/lander"
	"Def/state/player"
	"Def/state/pod"
	"Def/state/swarmer"
	"Def/systems"
	"Def/types"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var blankImg *ebiten.Image

func InitGame(engine *game.Engine) {

	graphics.Load()

	initSystems(engine)
	initEntities(engine)
	bulletPool(engine)
	bombPool(engine)
	laserPool(engine)
	InitEvents(engine)

	gl.ScoreCharId = engine.AddString("       0", 100, 40)
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

	}

	for i := 0; i < gl.BaiterCount; i++ {
		AddBaiter(engine, i)
	}

	for i := 0; i < gl.HumanCount; i++ {

		AddHuman(engine, i)

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
	plEnt := game.NewEntity(engine, types.Player)
	gl.PlayerID = plEnt.Id
	plEnt.SetActive(true)

	x := float64(gl.WorldWidth) / 2
	y := float64(gl.ScreenHeight) / 2

	sgraph := game.NewStateGraph()
	sgraph.AddState(player.NewPlayerPlay())
	sgraph.AddState(player.NewPlayerDie())

	fsm := game.NewFSM(sgraph)
	ai := cmp.NewAI(fsm, types.PlayerPlay)
	plEnt.AddComponent(ai)
	smap := graphics.GetSpriteMap("ship.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	plEnt.AddComponent(dr)
	col := types.ColorF{R: 1, G: 1, B: 1, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	plEnt.AddComponent(rd)
	sc := cmp.NewShip(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	plEnt.AddComponent(sc)
	pc := cmp.NewPos(x, y, 0, 0)
	plEnt.AddComponent(pc)

	//flame

	fEnt := game.NewEntity(engine, types.Player)
	fEnt.SetActive(true)
	fsmap := graphics.GetSpriteMap("thrust.png")
	fdr := cmp.NewDraw(ssheet, fsmap, types.ColorF{R: 1, G: 1, B: 1})

	fdr.Scale = 0.7
	fEnt.AddComponent(fdr)
	plEnt.SetChild(fEnt.Id)
	fpc := cmp.NewPos(0, 0, 0, 0)
	fEnt.AddComponent(fpc)

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
	sgraph := game.NewStateGraph()
	sgraph.AddState(lander.NewLanderWait())
	sgraph.AddState(lander.NewLanderSearch())
	sgraph.AddState(lander.NewLanderMaterialise())
	sgraph.AddState(lander.NewLanderDrop())
	sgraph.AddState(lander.NewLanderGrab())
	sgraph.AddState(lander.NewLanderMutate())
	sgraph.AddState(lander.NewLanderDie())

	fsm := game.NewFSM(sgraph)
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

func AddBaiter(engine *game.Engine, count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := game.NewEntity(engine, types.Baiter)
	ent.SetActive(true)

	pc := cmp.NewPos(0, 0, 0, 0)
	ent.AddComponent(pc)
	sgraph := game.NewStateGraph()
	sgraph.AddState(baiter.NewBaiterWait())
	sgraph.AddState(baiter.NewBaiterMaterialise())
	sgraph.AddState(baiter.NewBaiterHunt())
	sgraph.AddState(baiter.NewBaiterDie())

	fsm := game.NewFSM(sgraph)
	ai := cmp.NewAI(fsm, types.BaiterWait)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("baiter.png")
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)
	col := types.ColorF{R: 0, G: 0.5, B: 0, A: 1}
	rd := cmp.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	sh := cmp.NewShootable()
	ent.AddComponent(sh)
	cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)

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
	sgraph := game.NewStateGraph()
	sgraph.AddState(human.NewHumanWalking())
	sgraph.AddState(human.NewHumanGrabbed())
	sgraph.AddState(human.NewHumanDropping())
	sgraph.AddState(human.NewHumanRescued())
	sgraph.AddState(human.NewHumanDie())

	fsm := game.NewFSM(sgraph)
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
	sgraph := game.NewStateGraph()
	sgraph.AddState(bomber.NewBomberMove())
	sgraph.AddState(bomber.NewBomberDie())

	fsm := game.NewFSM(sgraph)
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
	sgraph := game.NewStateGraph()
	sgraph.AddState(pod.NewPodMove())
	sgraph.AddState(pod.NewPodDie())

	fsm := game.NewFSM(sgraph)
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
	sgraph := game.NewStateGraph()
	sgraph.AddState(swarmer.NewSwarmerMove())
	sgraph.AddState(swarmer.NewSwarmerDie())

	fsm := game.NewFSM(sgraph)
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

func AddScoreSprite(engine *game.Engine, ev event.IEvent) {

	eve := ev.GetPayload().(*game.Entity)
	evpc := eve.GetComponent(types.Pos).(*cmp.Pos)

	ent := game.NewEntity(engine, types.Score)
	ent.SetActive(true)
	pc := cmp.NewPos(evpc.X-gl.CameraX(), evpc.Y, 0, 0)
	pc.ScreenCoords = true
	ent.AddComponent(pc)
	s := "500.png"
	if ev.GetType() == event.HumanLandedEvent {
		s = "250.png"
	}
	smap := graphics.GetSpriteMap(s)
	ssheet := graphics.GetSpriteSheet()
	dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Scale = 1
	ent.AddComponent(dr)
	li := cmp.NewLife(60)
	ent.AddComponent(li)

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
		sound.Play(sound.Background)
		sound.Play(sound.Levelstart)
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
			sound.Play(sound.Bullet)
		}

	}

	materialise := func(e event.IEvent) {
		sound.PlayIfNot(sound.Materialise)
	}

	landerDie := func(e event.IEvent) {
		gl.Score += 150
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		ent := e.GetPayload().(*game.Entity)
		gl.LandersKilled++
		if gl.LandersKilled == gl.LanderCount {
			lc := event.NewLanderCleared(ent)
			event.NotifyEvent(lc)
		}
		sound.Stop(sound.Laser)
		sound.Play(sound.Landerdie)
	}

	landerCleared := func(e event.IEvent) {
		if ent := e.GetPayload().(*game.Entity); ent != nil {
			// end of level
		}
	}

	mutantSound := func(e event.IEvent) {
		sound.PlayIfNot(sound.Mutant)
	}

	humanDie := func(e event.IEvent) {

		gl.HumansKilled++
		if gl.HumansKilled == gl.HumanCount {
			engine.ExplodeWorld()
			engine.SetFlash(30)
			engine.MutateAll()
		}
		sound.Play(sound.Humandie)
	}

	bomberDie := func(e event.IEvent) {
		gl.Score += 250
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		sound.Stop(sound.Laser)
		sound.Play(sound.Bomberdie)
	}

	playerDie := func(e event.IEvent) {
		pe := engine.GetEntities()[gl.PlayerID]
		pai := pe.GetComponent(types.AI).(*cmp.AI)
		pai.NextState = types.PlayerDie
		engine.GetSystem(game.PosSystem).SetActive(false)
	}

	playerExplode := func(e event.IEvent) {
		sound.Play(sound.Die)
	}

	playerFire := func(e event.IEvent) {
		if pe := e.GetPayload().(*game.Entity); pe != nil {
			pc := pe.GetComponent(types.Pos).(*cmp.Pos)
			sc := pe.GetComponent(types.Ship).(*cmp.Ship)
			x := pc.X + 25
			y := pc.Y + 25
			if sc.Direction < 0 {
				x = pc.X - 100
			}
			engine.TriggerLaser(x, y, sc.Direction)
			sound.Play(sound.Laser)
		}
	}

	smartBomb := func(e event.IEvent) {
		engine.SetFlash(1)
		engine.SmartBomb()
	}

	podDie := func(e event.IEvent) {
		gl.Score += 1000
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		ent := e.GetPayload().(*game.Entity)
		pc := ent.GetComponent(types.Pos).(*cmp.Pos)
		for i := 0; i < gl.SwarmerCount; i++ {
			AddSwarmer(engine, i, pc.X, pc.Y)
		}
		sound.Stop(sound.Laser)
		sound.Play(sound.Poddie)
	}

	swarmerDie := func(e event.IEvent) {
		gl.Score += 150
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
	}

	baiterDie := func(e event.IEvent) {
		gl.Score += 200
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		sound.Play(sound.Baiterdie)
	}

	HumanDropped := func(e event.IEvent) {
		sound.Play(sound.Dropping)
	}

	humanGrabbed := func(e event.IEvent) {
		sound.Play(sound.Grabbed)
	}

	humanRescued := func(e event.IEvent) {
		AddScoreSprite(engine, e)
		gl.Score += 500
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
	}

	humanSaved := func(e event.IEvent) {
		AddScoreSprite(engine, e)
		gl.Score += 500
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		sound.Play(sound.Placehuman)
	}

	humanLanded := func(e event.IEvent) {
		AddScoreSprite(engine, e)
		gl.Score += 250
		engine.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
	}

	thrustOn := func(e event.IEvent) {
		sound.Play(sound.Thruster)
	}

	thrustOff := func(e event.IEvent) {
		sound.Stop(sound.Thruster)
	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)
	event.AddEventListener(event.LanderDieEvent, landerDie)
	event.AddEventListener(event.LanderClearedEvent, landerCleared)
	event.AddEventListener(event.HumanDieEvent, humanDie)
	event.AddEventListener(event.BomberDieEvent, bomberDie)
	event.AddEventListener(event.PlayerDieEvent, playerDie)
	event.AddEventListener(event.PlayerExplodeEvent, playerExplode)
	event.AddEventListener(event.StartEvent, start)
	event.AddEventListener(event.PlayerFireEvent, playerFire)
	event.AddEventListener(event.SmartBombEvent, smartBomb)
	event.AddEventListener(event.PlayerCollideEvent, playerCollide)
	event.AddEventListener(event.PodDieEvent, podDie)
	event.AddEventListener(event.BaiterDieEvent, baiterDie)
	event.AddEventListener(event.SwarmerDieEvent, swarmerDie)
	event.AddEventListener(event.HumanRescuedEvent, humanRescued)
	event.AddEventListener(event.HumanSavedEvent, humanSaved)
	event.AddEventListener(event.HumanLandedEvent, humanLanded)
	event.AddEventListener(event.HumanGrabbedEvent, humanGrabbed)
	event.AddEventListener(event.HumanDroppedEvent, HumanDropped)
	event.AddEventListener(event.PlayerThrustEvent, thrustOn)
	event.AddEventListener(event.PlayerStopThrustEvent, thrustOff)
	event.AddEventListener(event.MaterialiseEvent, materialise)
	event.AddEventListener(event.MutantSoundEvent, mutantSound)
}
