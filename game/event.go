package game

type EventType int

const (
	// test events
	EnterState1Event  EventType = iota
	EnterState2Event  EventType = iota
	UpdateState1Event EventType = iota
	UpdateState2Event EventType = iota
)

func (ev EventType) String() string {
	switch ev {
	case EnterState1Event:
		return "EnterState1"
	case EnterState2Event:
		return "EnterState2"
	case UpdateState1Event:
		return "UpdateState1"
	case UpdateState2Event:
		return "UpdateState2"
	}
	return ""
}

var events map[EventType][]func(*Entity)

func init() {
	events = make(map[EventType][]func(*Entity))
}

func AddEventListener(et EventType, fn func(*Entity)) {
	if _, ok := events[et]; !ok {
		events[et] = []func(*Entity){}
	}
	events[et] = append(events[et], fn)
}

func NotifyEvent(et EventType, e *Entity) {
	if evs, ok := events[et]; ok {
		for _, v := range evs {
			v(e)
		}
	}
}
