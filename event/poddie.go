package event

type PodDie struct {
	mytype  EventType
	payload interface{}
}

func NewPodDie(p interface{}) *PodDie {
	return &PodDie{
		mytype:  PodDieEvent,
		payload: p,
	}
}

func (ev *PodDie) GetType() EventType {
	return ev.mytype
}

func (ev *PodDie) GetPayload() interface{} {
	return ev.payload
}
