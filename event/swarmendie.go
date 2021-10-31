package event

type SwarmerDie struct {
	mytype  EventType
	payload interface{}
}

func NewSwarmerDie(p interface{}) *SwarmerDie {
	return &SwarmerDie{
		mytype:  SwarmerDieEvent,
		payload: p,
	}
}

func (ev *SwarmerDie) GetType() EventType {
	return ev.mytype
}

func (ev *SwarmerDie) GetPayload() interface{} {
	return ev.payload
}
