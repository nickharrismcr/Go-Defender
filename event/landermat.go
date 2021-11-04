package event

type LanderMaterialise struct {
	mytype  EventType
	payload interface{}
}

func NewLanderMaterialise(p interface{}) *LanderMaterialise {
	return &LanderMaterialise{
		mytype:  LanderMaterialiseEvent,
		payload: p,
	}
}

func (ev *LanderMaterialise) GetType() EventType {
	return ev.mytype
}

func (ev *LanderMaterialise) GetPayload() interface{} {
	return ev.payload
}
