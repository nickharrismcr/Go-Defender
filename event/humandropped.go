package event

type HumanDropped struct {
	mytype  EventType
	payload interface{}
}

func NewHumanDropped(p interface{}) *HumanDropped {
	return &HumanDropped{
		mytype:  HumanDroppedEvent,
		payload: p,
	}
}

func (ev *HumanDropped) GetType() EventType {
	return ev.mytype
}

func (ev *HumanDropped) GetPayload() interface{} {
	return ev.payload
}
