package pod

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type PodMove struct {
	Name types.StateType
}

func NewPodMove() *PodMove {
	return &PodMove{
		Name: types.PodMove,
	}
}

func (s *PodMove) GetName() types.StateType {
	return s.Name
}

func (s *PodMove) Enter(ai *cmp.AI, e types.IEntity) {

}

func (s *PodMove) Update(ai *cmp.AI, e types.IEntity) {

}
