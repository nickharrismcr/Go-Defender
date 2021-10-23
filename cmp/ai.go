package cmp

import "Def/types"

var idCounter int

type AI struct {
	componentType types.CmpType
	Id            int
	Counter       int
	FSMId         int
	State         types.StateType
	NextState     types.StateType
	TargetId      types.EntityID
	Scratch       int
}

func init() {
	idCounter = 0
}

func NewAI(FSMId int, initState types.StateType) *AI {
	idCounter++
	return &AI{
		Id:            idCounter,
		FSMId:         FSMId,
		State:         -1,
		NextState:     initState,
		Counter:       0,
		componentType: types.AI,
		Scratch:       0,
	}
}

func (ai *AI) Type() types.CmpType {
	return ai.componentType
}
