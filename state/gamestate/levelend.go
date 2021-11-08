package gamestate

import (
	"Def/cmp"
	"Def/gl"
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
	e.GetEngine().LevelEnd()
	ai.Scratch = 0
}

func (s *GameLevelEnd) Update(ai *cmp.AI, e types.IEntity) {

	ai.Scratch++

	if ai.Scratch > 4*30 {
		gl.NextLevel()
		ai.NextState = types.GameStart
	}
}
