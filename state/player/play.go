package player

import (
	"Def/cmp"
	"Def/global"
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
	sc.ScreenOffset = global.ScreenWidth * 0.1
}

func (s *PlayerPlay) Update(ai *cmp.AI, e types.IEntity) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	sc := e.GetComponent(types.Ship).(*cmp.Ship)
	dc := e.GetComponent(types.Draw).(*cmp.Draw)

	if sc.Direction == 1 && sc.ScreenOffset > global.ScreenWidth*0.1 {
		sc.ScreenOffset -= 30
	}
	if sc.Direction == -1 && sc.ScreenOffset < global.ScreenWidth*0.9 {
		sc.ScreenOffset += 30
	}

	e.GetEngine().SetCameraX(pc.X - sc.ScreenOffset)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !sc.ReversePressed {
			sc.Direction = -sc.Direction
			dc.FlipX = !dc.FlipX
			sc.ReversePressed = true
		}
	} else {
		sc.ReversePressed = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		pc.DX += sc.Direction * 2
	} else {
		pc.DX /= 1.05
	}
	if pc.DX > global.PlayerSpeedX {
		pc.DX = global.PlayerSpeedX
	}
	if pc.DX < -global.PlayerSpeedX {
		pc.DX = -global.PlayerSpeedX
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) && ebiten.IsKeyPressed(ebiten.KeyA) {
		pc.DY = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyQ) {
		pc.DY = -global.PlayerSpeedY
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		pc.DY = global.PlayerSpeedY
	} else {
		pc.DY = 0
	}

}
