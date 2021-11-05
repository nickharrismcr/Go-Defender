package player

import (
	"Def/cmp"
	"Def/event"
	"Def/gl"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// NB States should not contain entity state ;) they should act on cmp

type PlayerPlay struct {
	Name types.StateType
}

func NewPlayerPlay() *PlayerPlay {
	return &PlayerPlay{
		Name: types.PlayerPlay,
	}
}

func (s *PlayerPlay) GetName() types.StateType {
	return s.Name
}

func (s *PlayerPlay) Enter(ai *cmp.AI, e types.IEntity) {
	sc := e.GetComponent(types.Ship).(*cmp.Ship)
	sc.ScreenOffset = gl.ScreenWidth * 0.1
	ev := event.NewStart(e)
	event.NotifyEvent(ev)
}

func (s *PlayerPlay) Update(ai *cmp.AI, e types.IEntity) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	sc := e.GetComponent(types.Ship).(*cmp.Ship)
	dc := e.GetComponent(types.Draw).(*cmp.Draw)

	if sc.Direction == 1 && sc.ScreenOffset > gl.ScreenWidth*0.1 {
		sc.ScreenOffset -= 30
	}
	if sc.Direction == -1 && sc.ScreenOffset < gl.ScreenWidth*0.9 {
		sc.ScreenOffset += 30
	}

	camx := (pc.X - sc.ScreenOffset)
	if camx < 0 {
		camx += gl.WorldWidth
	}
	gl.SetCameraX(camx)

	fle := e.GetEngine().GetEntity(e.Child())
	flpc := fle.GetComponent(types.Pos).(*cmp.Pos)
	fldc := fle.GetComponent(types.Draw).(*cmp.Draw)

	flpc.X = pc.X - 40
	if sc.Direction < 0 {
		flpc.X = pc.X + 5
	}
	flpc.Y = pc.Y + 10

	fdc := fle.GetComponent(types.Draw).(*cmp.Draw)

	if ebiten.IsKeyPressed(gl.KeyMap[types.Reverse]) {
		if !sc.ReversePressed {
			sc.Direction = -sc.Direction
			dc.FlipX = !dc.FlipX
			fdc.FlipX = !fdc.FlipX
			sc.ReversePressed = true
			pc.DX /= 2
		}
	} else {
		sc.ReversePressed = false
	}
	if ebiten.IsKeyPressed(gl.KeyMap[types.Thrust]) {
		if !sc.ThrustPressed {
			sc.ThrustPressed = true
			ev := event.NewPlayerThrust(e)
			event.NotifyEvent(ev)
			fldc.Hide = false
		}
		pc.DX += sc.Direction * 2
	} else {
		fldc.Hide = true
		sc.ThrustPressed = false
		ev := event.NewPlayerStopThrust(e)
		event.NotifyEvent(ev)
		pc.DX /= 1.05
	}
	if pc.DX > gl.PlayerSpeedX {
		pc.DX = gl.PlayerSpeedX
	}
	if pc.DX < -gl.PlayerSpeedX {
		pc.DX = -gl.PlayerSpeedX
	}

	if ebiten.IsKeyPressed(gl.KeyMap[types.Up]) && ebiten.IsKeyPressed(gl.KeyMap[types.Down]) {
		pc.DY = 0
	} else if ebiten.IsKeyPressed(gl.KeyMap[types.Up]) {
		if pc.DY > -gl.PlayerSpeedY {
			pc.DY -= 2
		}
	} else if ebiten.IsKeyPressed(gl.KeyMap[types.Down]) {
		if pc.DY < gl.PlayerSpeedY {
			pc.DY += 2
		}
	} else {
		pc.DY = 0
	}

	if ebiten.IsKeyPressed(gl.KeyMap[types.Fire]) {
		if !sc.FirePressed {
			sc.FirePressed = true
			ev := event.NewPlayerFire(e)
			event.NotifyEvent(ev)
		}
	} else {
		sc.FirePressed = false
	}

	if ebiten.IsKeyPressed(gl.KeyMap[types.SmartBomb]) {
		if !sc.SmartBombPressed {
			sc.SmartBombPressed = true
			ev := event.NewSmartBomb(pc)
			event.NotifyEvent(ev)
			event.NotifyEventDelay(ev, 5)
			event.NotifyEventDelay(ev, 10)
		}
	} else {
		sc.SmartBombPressed = false
	}

	if ebiten.IsKeyPressed(gl.KeyMap[types.HyperSpace]) {
		if !sc.HyperSpacePressed {
			sc.HyperSpacePressed = true
		}
	} else {
		sc.HyperSpacePressed = false
	}
}
