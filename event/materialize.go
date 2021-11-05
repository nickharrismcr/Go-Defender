package event

type Materialise struct {
	mytype  EventType
	payload interface{}
}

func NewMaterialise(p interface{}) *Materialise {
	return &Materialise{
		mytype:  MaterialiseEvent,
		payload: p,
	}
}

func (ev *Materialise) GetType() EventType {
	return ev.mytype
}

func (ev *Materialise) GetPayload() interface{} {
	return ev.payload
}
