package event

type PlayerFire struct {
	mytype  EventType
	payload interface{}
}

func NewPlayerFire(p interface{}) *PlayerFire {
	return &PlayerFire{
		mytype:  PlayerFireEvent,
		payload: p,
	}
}

func (ev *PlayerFire) GetType() EventType {
	return ev.mytype
}

func (ev *PlayerFire) GetPayload() interface{} {
	return ev.payload
}
