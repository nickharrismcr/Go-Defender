package gamestate

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type GameStart struct {
	Name types.StateType
}

func NewGameStart() *GameStart {
	return &GameStart{
		Name: types.GameStart,
	}
}

func (s *GameStart) GetName() types.StateType {
	return s.Name
}

func (s *GameStart) Enter(ai *cmp.AI, e types.IEntity) {

	e.GetEngine().LevelStart()
}

func (s *GameStart) Update(ai *cmp.AI, e types.IEntity) {
	ai.NextState = types.GamePlay
}
