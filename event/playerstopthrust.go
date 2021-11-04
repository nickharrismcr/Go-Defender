package event

type PlayerStopThrust struct {
	mytype  EventType
	payload interface{}
}

func NewPlayerStopThrust(p interface{}) *PlayerStopThrust {
	return &PlayerStopThrust{
		mytype:  PlayerStopThrustEvent,
		payload: p,
	}
}

func (ev *PlayerStopThrust) GetType() EventType {
	return ev.mytype
}

func (ev *PlayerStopThrust) GetPayload() interface{} {
	return ev.payload
}
