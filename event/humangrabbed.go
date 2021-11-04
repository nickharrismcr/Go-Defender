package event

type HumanGrabbed struct {
	mytype  EventType
	payload interface{}
}

func NewHumanGrabbed(p interface{}) *HumanGrabbed {
	return &HumanGrabbed{
		mytype:  HumanGrabbedEvent,
		payload: p,
	}
}

func (ev *HumanGrabbed) GetType() EventType {
	return ev.mytype
}

func (ev *HumanGrabbed) GetPayload() interface{} {
	return ev.payload
}
