package cmp

var idCounter int

type AICmp struct {
	componentType CmpType
	Id            int
	Counter       int
	FSMId         int
	StateName     string
	NextStateName string
}

func init() {
	idCounter = 0
}

func NewAI(FSMId int, init_state string) *AICmp {
	idCounter++
	return &AICmp{
		Id:            idCounter,
		FSMId:         FSMId,
		StateName:     "",
		NextStateName: init_state,
		Counter:       0,
		componentType: AIType,
	}
}

func (ai *AICmp) Type() CmpType {
	return ai.componentType
}
