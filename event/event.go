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
	PlayerCollideEvent EventType = iota
	PlayerFireEvent    EventType = iota
	SmartBombEvent     EventType = iota
	SwarmerDieEvent    EventType = iota
	PodDieEvent        EventType = iota
	HumanRescuedEvent  EventType = iota
	HumanLandedEvent   EventType = iota
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
	case PlayerCollideEvent:
		return "PlayerCollide"
	case PlayerFireEvent:
		return "PlayerFire"
	case SmartBombEvent:
		return "SmartBomb"
	case PodDieEvent:
		return "PodDie"
	case SwarmerDieEvent:
		return "SwarmerDie"
	case HumanRescuedEvent:
		return "HumanRescued"
	case HumanLandedEvent:
		return "HumanLanded"
	}

	return ""
}

type QueueEntry struct {
	event      IEvent
	delayTicks int
}

var events map[EventType][]func(IEvent)
var queue []*QueueEntry

func init() {
	events = make(map[EventType][]func(IEvent))
	queue = []*QueueEntry{}
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

// add event to queue, record delay
func NotifyEventDelay(e IEvent, ticks int) {
	logger.Debug("EventDelay notified : %s (%d) ", e.GetType().String(), ticks)
	queue = append(queue, &QueueEntry{event: e, delayTicks: ticks})
}

func UpdateQueue() {

	for i, e := range queue {
		logger.Debug("Event queue updated : %s (%d) ", e.event.GetType().String(), e.delayTicks)
		e.delayTicks--
		if e.delayTicks <= 0 {
			NotifyEvent(e.event)
			queue = queue[:i]
			queue = append(queue, queue[i:]...)
		}
	}
}
