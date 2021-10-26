package event

type BaiterDie struct {
	mytype  EventType
	payload interface{}
}

func NewBaiterDie(p interface{}) *BaiterDie {
	return &BaiterDie{
		mytype:  BaiterDieEvent,
		payload: p,
	}
}

func (ev *BaiterDie) GetType() EventType {
	return ev.mytype
}

func (ev *BaiterDie) GetPayload() interface{} {
	return ev.payload
}
