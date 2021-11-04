package event

type PlayerThrust struct {
	mytype  EventType
	payload interface{}
}

func NewPlayerThrust(p interface{}) *PlayerThrust {
	return &PlayerThrust{
		mytype:  PlayerThrustEvent,
		payload: p,
	}
}

func (ev *PlayerThrust) GetType() EventType {
	return ev.mytype
}

func (ev *PlayerThrust) GetPayload() interface{} {
	return ev.payload
}
