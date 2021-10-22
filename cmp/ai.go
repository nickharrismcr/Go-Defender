package cmp

import "Def/types"

var idCounter int

type AI struct {
	componentType types.CmpType
	Id            int
	Counter       int
	FSMId         int
	StateName     types.StateType
	NextStateName types.StateType
	TargetId      types.EntityID
}

func init() {
	idCounter = 0
}

func NewAI(FSMId int, initState types.StateType) *AI {
	idCounter++
	return &AI{
		Id:            idCounter,
		FSMId:         FSMId,
		StateName:     -1,
		NextStateName: initState,
		Counter:       0,
		componentType: types.AI,
	}
}

func (ai *AI) Type() types.CmpType {
	return ai.componentType
}
