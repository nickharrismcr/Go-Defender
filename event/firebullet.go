package event

type FireBullet struct {
	mytype  EventType
	payload interface{}
}

func NewFireBullet(p interface{}) *FireBullet {
	return &FireBullet{
		mytype:  FireBulletEvent,
		payload: p,
	}
}

func (ev *FireBullet) GetType() EventType {
	return ev.mytype
}

func (ev *FireBullet) GetPayload() interface{} {
	return ev.payload
}
