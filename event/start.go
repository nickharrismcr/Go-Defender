package event

type Start struct {
	mytype  EventType
	payload interface{}
}

func NewStart(p interface{}) *Start {
	return &Start{
		mytype:  StartEvent,
		payload: p,
	}
}

func (ev *Start) GetType() EventType {
	return ev.mytype
}

func (ev *Start) GetPayload() interface{} {
	return ev.payload
}
