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
	StartEvent         EventType = iota
	FireBulletEvent    EventType = iota
	ExplodeEvent       EventType = iota
	LanderDieEvent     EventType = iota
	HumanDieEvent      EventType = iota
	BaiterDieEvent     EventType = iota
	BomberDieEvent     EventType = iota
	LanderClearedEvent EventType = iota
	PlayerDieEvent     EventType = iota
	PlayerFireEvent    EventType = iota
	SmartBombEvent     EventType = iota
)

func (ev EventType) String() string {
	switch ev {

	case StartEvent:
		return "Start"
	case ExplodeEvent:
		return "Explode"
	case FireBulletEvent:
		return "FireBullet"
	case LanderDieEvent:
		return "LanderDie"
	case HumanDieEvent:
		return "LanderDie"
	case BaiterDieEvent:
		return "LanderDie"
	case LanderClearedEvent:
		return "LanderCleared"
	case PlayerDieEvent:
		return "PlayerDie"
	case PlayerFireEvent:
		return "PlayerFire"
	case SmartBombEvent:
		return "SmartBomb"
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
