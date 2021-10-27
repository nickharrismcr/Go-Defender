package player

import (
	"Def/cmp"
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

}

func (s *PlayerPlay) Update(ai *cmp.AI, e types.IEntity) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	e.GetEngine().SetCameraX(pc.X - 100)
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		pc.X -= 50
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		pc.X += 50
	}
}
