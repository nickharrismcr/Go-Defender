package event

type HumanDie struct {
	mytype  EventType
	payload interface{}
}

func NewHumanDie(p interface{}) *HumanDie {
	return &HumanDie{
		mytype:  HumanDieEvent,
		payload: p,
	}
}

func (ev *HumanDie) GetType() EventType {
	return ev.mytype
}

func (ev *HumanDie) GetPayload() interface{} {
	return ev.payload
}
