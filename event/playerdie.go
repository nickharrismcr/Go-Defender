package event

type PlayerDie struct {
	mytype  EventType
	payload interface{}
}

func NewPlayerDie(p interface{}) *PlayerDie {
	return &PlayerDie{
		mytype:  PlayerDieEvent,
		payload: p,
	}
}

func (ev *PlayerDie) GetType() EventType {
	return ev.mytype
}

func (ev *PlayerDie) GetPayload() interface{} {
	return ev.payload
}
