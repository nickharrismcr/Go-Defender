package event

type BomberDie struct {
	mytype  EventType
	payload interface{}
}

func NewBomberDie(p interface{}) *BomberDie {
	return &BomberDie{
		mytype:  BomberDieEvent,
		payload: p,
	}
}

func (ev *BomberDie) GetType() EventType {
	return ev.mytype
}

func (ev *BomberDie) GetPayload() interface{} {
	return ev.payload
}
