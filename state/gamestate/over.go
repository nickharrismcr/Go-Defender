package gamestate

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type GameOver struct {
	Name types.StateType
}

func NewGameOver() *GameOver {
	return &GameOver{
		Name: types.GameOver,
	}
}

func (s *GameOver) GetName() types.StateType {
	return s.Name
}

func (s *GameOver) Enter(ai *cmp.AI, e types.IEntity) {

}

func (s *GameOver) Update(ai *cmp.AI, e types.IEntity) {

}
