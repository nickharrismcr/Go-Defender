package event

type HumanLanded struct {
	mytype  EventType
	payload interface{}
}

func NewHumanLanded(p interface{}) *HumanLanded {
	return &HumanLanded{
		mytype:  HumanLandedEvent,
		payload: p,
	}
}

func (ev *HumanLanded) GetType() EventType {
	return ev.mytype
}

func (ev *HumanLanded) GetPayload() interface{} {
	return ev.payload
}
