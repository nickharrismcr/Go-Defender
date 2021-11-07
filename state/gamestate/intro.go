package gamestate

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type GameIntro struct {
	Name types.StateType
}

func NewGameIntro() *GameIntro {
	return &GameIntro{
		Name: types.GameIntro,
	}
}

func (s *GameIntro) GetName() types.StateType {
	return s.Name
}

func (s *GameIntro) Enter(ai *cmp.AI, e types.IEntity) {

}

func (s *GameIntro) Update(ai *cmp.AI, e types.IEntity) {
	ai.NextState = types.GameStart
}
