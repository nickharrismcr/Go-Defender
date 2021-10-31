package event

type PlayerCollide struct {
	mytype  EventType
	payload interface{}
}

func NewPlayerCollide(p interface{}) *PlayerCollide {
	return &PlayerCollide{
		mytype:  PlayerCollideEvent,
		payload: p,
	}
}

func (ev *PlayerCollide) GetType() EventType {
	return ev.mytype
}

func (ev *PlayerCollide) GetPayload() interface{} {
	return ev.payload
}
