package cmp

import "Def/types"

var idCounter int

type AICmp struct {
	componentType types.CmpType
	Id            int
	Counter       int
	FSMId         int
	StateName     types.StateType
	NextStateName types.StateType
}

func init() {
	idCounter = 0
}

func NewAI(FSMId int, initState types.StateType) *AICmp {
	idCounter++
	return &AICmp{
		Id:            idCounter,
		FSMId:         FSMId,
		StateName:     -1,
		NextStateName: initState,
		Counter:       0,
		componentType: types.AI,
	}
}

func (ai *AICmp) Type() types.CmpType {
	return ai.componentType
}
