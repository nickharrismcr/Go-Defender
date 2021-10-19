package event

type Explode struct {
	mytype  EventType
	payload interface{}
}

func NewExplode(p interface{}) *Explode {
	return &Explode{
		mytype:  ExplodeEvent,
		payload: p,
	}
}

func (ev *Explode) GetType() EventType {
	return ev.mytype
}

func (ev *Explode) GetPayload() interface{} {
	return ev.payload
}
