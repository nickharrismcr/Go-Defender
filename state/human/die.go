package human

import (
	"Def/cmp"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on cmp

type HumanDie struct {
	Name types.StateType
}

func NewHumanDie() *HumanDie {
	return &HumanDie{
		Name: types.HumanDie,
	}
}

func (s *HumanDie) GetName() types.StateType {
	return s.Name
}

func (s *HumanDie) Enter(ai *cmp.AI, e types.IEntity) {
	e.SetActive(false)
}

func (s *HumanDie) Update(ai *cmp.AI, e types.IEntity) {

}
