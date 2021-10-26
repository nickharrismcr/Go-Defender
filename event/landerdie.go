package event

type LanderDie struct {
	mytype  EventType
	payload interface{}
}

func NewLanderDie(p interface{}) *LanderDie {
	return &LanderDie{
		mytype:  LanderDieEvent,
		payload: p,
	}
}

func (ev *LanderDie) GetType() EventType {
	return ev.mytype
}

func (ev *LanderDie) GetPayload() interface{} {
	return ev.payload
}
