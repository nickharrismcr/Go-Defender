package event

type HumanSaved struct {
	mytype  EventType
	payload interface{}
}

func NewHumanSaved(p interface{}) *HumanSaved {
	return &HumanSaved{
		mytype:  HumanSavedEvent,
		payload: p,
	}
}

func (ev *HumanSaved) GetType() EventType {
	return ev.mytype
}

func (ev *HumanSaved) GetPayload() interface{} {
	return ev.payload
}
