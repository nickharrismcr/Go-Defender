package game

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type GameLevelEnd struct {
	Name types.StateType
}

func NewGameLevelEnd() *GameLevelEnd {
	return &GameLevelEnd{
		Name: types.GameLevelEnd,
	}
}

func (s *GameLevelEnd) GetName() types.StateType {
	return s.Name
}

func (s *GameLevelEnd) Enter(ai *cmp.AI, e types.IEntity) {

}

func (s *GameLevelEnd) Update(ai *cmp.AI, e types.IEntity) {

}
