package event

type LanderCleared struct {
	mytype  EventType
	payload interface{}
}

func NewLanderCleared(p interface{}) *LanderCleared {
	return &LanderCleared{
		mytype:  LanderClearedEvent,
		payload: p,
	}
}

func (ev *LanderCleared) GetType() EventType {
	return ev.mytype
}

func (ev *LanderCleared) GetPayload() interface{} {
	return ev.payload
}
