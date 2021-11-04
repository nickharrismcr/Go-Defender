package event

type PlayerExplode struct {
	mytype  EventType
	payload interface{}
}

func NewPlayerExplode(p interface{}) *PlayerExplode {
	return &PlayerExplode{
		mytype:  PlayerExplodeEvent,
		payload: p,
	}
}

func (ev *PlayerExplode) GetType() EventType {
	return ev.mytype
}

func (ev *PlayerExplode) GetPayload() interface{} {
	return ev.payload
}
