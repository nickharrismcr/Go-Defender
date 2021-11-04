package event

type MutantSound struct {
	mytype  EventType
	payload interface{}
}

func NewMutantSound(p interface{}) *MutantSound {
	return &MutantSound{
		mytype:  MutantSoundEvent,
		payload: p,
	}
}

func (ev *MutantSound) GetType() EventType {
	return ev.mytype
}

func (ev *MutantSound) GetPayload() interface{} {
	return ev.payload
}
