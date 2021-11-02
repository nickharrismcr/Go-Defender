package event

type HumanRescued struct {
	mytype  EventType
	payload interface{}
}

func NewHumanRescued(p interface{}) *HumanRescued {
	return &HumanRescued{
		mytype:  HumanRescuedEvent,
		payload: p,
	}
}

func (ev *HumanRescued) GetType() EventType {
	return ev.mytype
}

func (ev *HumanRescued) GetPayload() interface{} {
	return ev.payload
}
