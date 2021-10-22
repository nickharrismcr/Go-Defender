package event

type EntityDie struct {
	mytype  EventType
	payload interface{}
}

func NewEntityDie(p interface{}) *EntityDie {
	return &EntityDie{
		mytype:  EntityDieEvent,
		payload: p,
	}
}

func (ev *EntityDie) GetType() EventType {
	return ev.mytype
}

func (ev *EntityDie) GetPayload() interface{} {
	return ev.payload
}
