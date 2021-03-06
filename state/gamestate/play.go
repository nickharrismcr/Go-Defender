package gamestate

import (
	"Def/cmp"

	"Def/gl"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type GamePlay struct {
	Name types.StateType
}

func NewGamePlay() *GamePlay {
	return &GamePlay{
		Name: types.GamePlay,
	}
}

func (s *GamePlay) GetName() types.StateType {
	return s.Name
}

func (s *GamePlay) Enter(ai *cmp.AI, e types.IEntity) {

}

func (s *GamePlay) Update(ai *cmp.AI, e types.IEntity) {

	if gl.LandersKilled == gl.CurrentLevel().LanderCount {
		ai.NextState = types.GameLevelEnd
	}
	if gl.PlayerLives == 0 {
		ai.NextState = types.GameOver
	}

}
