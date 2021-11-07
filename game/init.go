package game

import (
	"Def/cmp"
	"Def/event"

	"Def/gl"
	"Def/graphics"
	"Def/logger"
	"Def/sound"
	"Def/state/baiter"
	"Def/state/bomber"
	"Def/state/gamestate"
	"Def/state/human"
	"Def/state/lander"
	"Def/state/player"
	"Def/state/pod"
	"Def/state/swarmer"
	"Def/systems"
	"Def/types"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

var blankImg *ebiten.Image

func (e *Engine) Init() {

	graphics.Load()

	blankImg = ebiten.NewImage(20, 20)
	blankImg.Fill(color.White)

	e.initSystems()
	e.addGame()
	e.addPlayer()

	e.bulletPool()
	e.bombPool()
	e.laserPool()
	e.initEvents()

	gl.ScoreCharId = e.AddString("       0", 100, 40)
}

func (e *Engine) initSystems() {

	e.AddSystem(systems.NewPosSystem(true, e), UPDATE)
	e.AddSystem(systems.NewAISystem(true, e), UPDATE)
	e.AddSystem(systems.NewLifeSystem(true, e), UPDATE)
	e.AddSystem(systems.NewCollideSystem(true, e), UPDATE)
	e.AddSystem(systems.NewDrawSystem(true, e), DRAW)
	e.AddSystem(systems.NewRadarDrawSystem(true, e), DRAW)
	e.AddSystem(systems.NewLaserDrawSystem(true, e), DRAW)
	e.AddSystem(systems.NewLaserMoveSystem(true, e), UPDATE)
}

func (e *Engine) InitEntities() {

	for i := 0; i < gl.CurrentLevel().LanderCount; i++ {
		e.addLander(i)
	}
	for i := 0; i < gl.CurrentLevel().BaiterCount; i++ {
		e.addBaiter(i)
	}
	for i := 0; i < gl.CurrentLevel().HumanCount; i++ {
		e.addHuman(i)
	}
	for i := 0; i < gl.CurrentLevel().BomberCount; i++ {
		e.addBomber(i)
	}
	for i := 0; i < gl.CurrentLevel().PodCount; i++ {
		e.addPod(i)
	}
}

func (e *Engine) addGame() {

	gme := NewEntity(e, types.Player)
	gme.SetActive(true)

	sgraph := systems.NewStateGraph()
	sgraph.AddState(gamestate.NewGameIntro())
	sgraph.AddState(gamestate.NewGameStart())
	sgraph.AddState(gamestate.NewGamePlay())
	sgraph.AddState(gamestate.NewGameLevelEnd())
	sgraph.AddState(gamestate.NewGameOver())

	fsm := systems.NewFSM(sgraph)
	ai := cmp.NewAI(fsm, types.GameIntro)
	gme.AddComponent(ai)

}

func (e *Engine) addPlayer() {

	ssheet := graphics.GetSpriteSheet()
	plEnt := NewEntity(e, types.Player)
	gl.PlayerID = plEnt.Id
	plEnt.SetActive(true)

	x := float64(gl.WorldWidth) / 2
	y := float64(gl.ScreenHeight) / 2

	sgraph := systems.NewStateGraph()
	sgraph.AddState(player.NewPlayerPlay())
	sgraph.AddState(player.NewPlayerDie())

	fsm := systems.NewFSM(sgraph)
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

	fEnt := NewEntity(e, types.Player)
	fEnt.SetActive(true)
	fsmap := graphics.GetSpriteMap("thrust.png")
	fdr := cmp.NewDraw(ssheet, fsmap, types.ColorF{R: 1, G: 1, B: 1})

	fdr.Scale = 0.7
	fEnt.AddComponent(fdr)
	plEnt.SetChild(fEnt.Id)
	fpc := cmp.NewPos(0, 0, 0, 0)
	fEnt.AddComponent(fpc)

}

func (e *Engine) addLander(count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := NewEntity(e, types.Lander)
	ent.SetActive(true)

	x := rand.Float64() * gl.WorldWidth
	if count < 2 {
		x = gl.WorldWidth * 0.8
	}
	pc := cmp.NewPos(x, gl.ScreenTop+500*rand.Float64(), 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(lander.NewLanderWait())
	sgraph.AddState(lander.NewLanderSearch())
	sgraph.AddState(lander.NewLanderMaterialise())
	sgraph.AddState(lander.NewLanderDrop())
	sgraph.AddState(lander.NewLanderGrab())
	sgraph.AddState(lander.NewLanderMutate())
	sgraph.AddState(lander.NewLanderDie())

	fsm := systems.NewFSM(sgraph)
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

func (e *Engine) addBaiter(count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := NewEntity(e, types.Baiter)
	ent.SetActive(true)

	pc := cmp.NewPos(0, 0, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(baiter.NewBaiterWait())
	sgraph.AddState(baiter.NewBaiterMaterialise())
	sgraph.AddState(baiter.NewBaiterHunt())
	sgraph.AddState(baiter.NewBaiterDie())

	fsm := systems.NewFSM(sgraph)
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

func (e *Engine) addHuman(count int) {

	ssheet := graphics.GetSpriteSheet()
	ent := NewEntity(e, types.Human)
	ent.SetActive(true)

	x := rand.Float64() * gl.WorldWidth
	if count < 2 {
		x = rand.Float64()*gl.ScreenWidth + gl.CameraX()
	}
	pc := cmp.NewPos(x, 0, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(human.NewHumanWalking())
	sgraph.AddState(human.NewHumanGrabbed())
	sgraph.AddState(human.NewHumanDropping())
	sgraph.AddState(human.NewHumanRescued())
	sgraph.AddState(human.NewHumanDie())

	fsm := systems.NewFSM(sgraph)
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

func (e *Engine) addBomber(count int) {

	ent := NewEntity(e, types.Bomber)
	ent.SetActive(true)

	x := (rand.Float64() * gl.ScreenWidth) + gl.WorldWidth/3
	y := (rand.Float64() * gl.ScreenHeight / 2) + gl.ScreenTop + 50

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(bomber.NewBomberMove())
	sgraph.AddState(bomber.NewBomberDie())

	fsm := systems.NewFSM(sgraph)
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

func (e *Engine) addPod(count int) {

	ent := NewEntity(e, types.Pod)
	ent.SetActive(true)

	x := (rand.Float64() * gl.ScreenWidth) + gl.WorldWidth/2
	y := (rand.Float64() * gl.ScreenHeight / 2) + gl.ScreenTop + 50

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(pod.NewPodMove())
	sgraph.AddState(pod.NewPodDie())

	fsm := systems.NewFSM(sgraph)
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

func (e *Engine) AddSwarmer(count int, x, y float64) {

	ent := NewEntity(e, types.Swarmer)
	ent.SetActive(true)

	pc := cmp.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(swarmer.NewSwarmerMove())
	sgraph.AddState(swarmer.NewSwarmerDie())

	fsm := systems.NewFSM(sgraph)
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

func (e *Engine) AddScoreSprite(ev event.IEvent) {

	eve := ev.GetPayload().(*Entity)
	evpc := eve.GetComponent(types.Pos).(*cmp.Pos)

	ent := NewEntity(e, types.Score)
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

func (e *Engine) bulletPool() {

	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 40; i++ {
		ent := NewEntity(e, types.Bullet)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bullet.png")
		dr := cmp.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		ent.AddComponent(dr)
		li := cmp.NewLife(240)
		ent.AddComponent(li)
		cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
		ent.AddComponent(cl)
		e.BulletPool = append(e.BulletPool, ent)
	}

}

func (e *Engine) bombPool() {

	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 20; i++ {
		ent := NewEntity(e, types.Bomb)
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
		e.BombPool = append(e.BombPool, ent)
	}

}

func (e *Engine) laserPool() {

	for i := 0; i < 15; i++ {
		ent := NewEntity(e, types.Laser)
		pc := cmp.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		dr := cmp.NewLaserDraw()
		ent.AddComponent(dr)
		li := cmp.NewLife(240)
		ent.AddComponent(li)
		mv := cmp.NewLaserMove()
		ent.AddComponent(mv)

		e.LaserPool = append(e.LaserPool, ent)
	}

}

func (e *Engine) initEvents() {

	start := func(ev event.IEvent) {
		e.GetSystem(types.PosSystem).SetActive(true)
		sound.Play(sound.Background)
		sound.Play(sound.Levelstart)
		e.SetPauseAll(false, -1)
	}

	playerCollide := func(ev event.IEvent) {
		en := ev.GetPayload().(*Entity)
		logger.Info("Collide : %s ", en.GetClass().String())
		if en.GetClass() == types.Human {
			ai := en.GetComponent(types.AI).(*cmp.AI)
			if ai.State == types.HumanDropping {
				ai.NextState = types.HumanRescued
			}
		} else {
			e.Kill(en)
			e.SetPauseAll(true, en.GetID())
			ev := event.NewPlayerDie(e.GetPlayer())
			event.NotifyEvent(ev)
			gl.PlayerLives--
			if gl.PlayerLives == 0 {
				e.Terminate(GAME_OVER)
			}
		}
	}

	explodeTrigger := func(ev event.IEvent) {
		if ct := ev.GetPayload().(*cmp.Pos); ct != nil {
			e.TriggerPS(ct.X, ct.Y)
		}
	}

	bulletTrigger := func(ev event.IEvent) {
		if ct := ev.GetPayload().(*cmp.Pos); ct != nil {
			if math.Abs(ct.DX) < 20 {
				e.TriggerBullet(ct.X, ct.Y, ct.DX, ct.DY)
				sound.Play(sound.Bullet)
			}
		}
	}

	materialise := func(ev event.IEvent) {
		sound.PlayIfNot(sound.Materialise)
	}

	landerDie := func(ev event.IEvent) {
		gl.Score += 150
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		ent := ev.GetPayload().(*Entity)
		gl.LandersKilled++
		if gl.LandersKilled == gl.CurrentLevel().LanderCount {
			lc := event.NewLanderCleared(ent)
			event.NotifyEvent(lc)
		}
		sound.Stop(sound.Laser)
		sound.Play(sound.Landerdie)
	}

	landerCleared := func(ev event.IEvent) {
		runtime.GC()
		gl.NextLevel()
		//TODO pause, level end message, human move / bonus
		//TODO reset all entity states, restart
	}

	mutantSound := func(ev event.IEvent) {
		sound.PlayIfNot(sound.Mutant)
	}

	humanDie := func(ev event.IEvent) {
		gl.HumansKilled++
		if gl.HumansKilled == gl.CurrentLevel().HumanCount {
			e.ExplodeWorld()
			e.SetFlash(30)
			e.MutateAll()
		}
		sound.Play(sound.Humandie)
	}

	bomberDie := func(ev event.IEvent) {
		gl.Score += 250
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		sound.Stop(sound.Laser)
		sound.Play(sound.Bomberdie)
	}

	playerDie := func(ev event.IEvent) {
		pe := e.GetEntities()[gl.PlayerID]
		pai := pe.GetComponent(types.AI).(*cmp.AI)
		pai.NextState = types.PlayerDie

	}

	playerExplode := func(ev event.IEvent) {
		sound.Play(sound.Die)
	}

	playerFire := func(ev event.IEvent) {
		pe := ev.GetPayload().(*Entity)
		pc := pe.GetComponent(types.Pos).(*cmp.Pos)
		sc := pe.GetComponent(types.Ship).(*cmp.Ship)
		x := pc.X + 25
		y := pc.Y + 25
		if sc.Direction < 0 {
			x = pc.X - 100
		}
		e.TriggerLaser(x, y, sc.Direction)
		sound.Play(sound.Laser)

	}

	smartBomb := func(ev event.IEvent) {
		e.SetFlash(1)
		e.SmartBomb()
	}

	podDie := func(ev event.IEvent) {
		gl.Score += 1000
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		ent := ev.GetPayload().(*Entity)
		pc := ent.GetComponent(types.Pos).(*cmp.Pos)
		for i := 0; i < gl.CurrentLevel().SwarmerCount; i++ {
			e.AddSwarmer(i, pc.X, pc.Y)
		}
		sound.Stop(sound.Laser)
		sound.Play(sound.Poddie)
	}

	swarmerDie := func(ev event.IEvent) {
		gl.Score += 150
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
	}

	baiterDie := func(ev event.IEvent) {
		gl.Score += 200
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		sound.Play(sound.Baiterdie)
	}

	HumanDropped := func(ev event.IEvent) {
		sound.Play(sound.Dropping)
	}

	humanGrabbed := func(ev event.IEvent) {
		sound.Play(sound.Grabbed)
	}

	humanRescued := func(ev event.IEvent) {
		e.AddScoreSprite(ev)
		gl.Score += 500
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
	}

	humanSaved := func(ev event.IEvent) {
		e.AddScoreSprite(ev)
		gl.Score += 500
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
		sound.Play(sound.Placehuman)
	}

	humanLanded := func(ev event.IEvent) {
		e.AddScoreSprite(ev)
		gl.Score += 250
		e.ChangeString(gl.ScoreCharId, fmt.Sprintf("%8d", gl.Score))
	}

	thrustOn := func(ev event.IEvent) {
		sound.Play(sound.Thruster)
	}

	thrustOff := func(ev event.IEvent) {
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
