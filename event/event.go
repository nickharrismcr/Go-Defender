package event

import (
	"Def/logger"
)

type EventType int

type IEvent interface {
	GetType() EventType
	GetPayload() interface{}
}

const (
	// test events
	EnterState1Event  EventType = iota
	EnterState2Event  EventType = iota
	UpdateState1Event EventType = iota
	UpdateState2Event EventType = iota
	ExplodeEvent      EventType = iota
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
	case ExplodeEvent:
		return "Explode"
	}

	return ""
}

var events map[EventType][]func(IEvent)

func init() {
	events = make(map[EventType][]func(IEvent))
}

func AddEventListener(et EventType, fn func(IEvent)) {
	if _, ok := events[et]; !ok {
		events[et] = []func(IEvent){}
	}
	events[et] = append(events[et], fn)
}

func NotifyEvent(e IEvent) {
	logger.Debug("Event notified : %s ", e.GetType().String())
	if evs, ok := events[e.GetType()]; ok {
		for _, v := range evs {
			v(e)
		}
	}
}
