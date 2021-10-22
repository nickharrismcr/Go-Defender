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
	FireBulletEvent    EventType = iota
	ExplodeEvent       EventType = iota
	EntityDieEvent     EventType = iota
	LanderClearedEvent EventType = iota
)

func (ev EventType) String() string {
	switch ev {

	case ExplodeEvent:
		return "Explode"
	case FireBulletEvent:
		return "FireBullet"
	case EntityDieEvent:
		return "EntityDie"
	case LanderClearedEvent:
		return "LanderCleared"
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
