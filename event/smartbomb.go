package event

type SmartBomb struct {
	mytype  EventType
	payload interface{}
}

func NewSmartBomb(p interface{}) *SmartBomb {
	return &SmartBomb{
		mytype:  SmartBombEvent,
		payload: p,
	}
}

func (ev *SmartBomb) GetType() EventType {
	return ev.mytype
}

func (ev *SmartBomb) GetPayload() interface{} {
	return ev.payload
}
