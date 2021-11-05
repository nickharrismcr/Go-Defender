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
	StartEvent            EventType = iota
	FireBulletEvent       EventType = iota
	ExplodeEvent          EventType = iota
	LanderDieEvent        EventType = iota
	MutantSoundEvent      EventType = iota
	HumanDieEvent         EventType = iota
	BaiterDieEvent        EventType = iota
	BomberDieEvent        EventType = iota
	LanderClearedEvent    EventType = iota
	PlayerDieEvent        EventType = iota
	PlayerCollideEvent    EventType = iota
	PlayerFireEvent       EventType = iota
	PlayerThrustEvent     EventType = iota
	PlayerStopThrustEvent EventType = iota
	SmartBombEvent        EventType = iota
	SwarmerDieEvent       EventType = iota
	PodDieEvent           EventType = iota
	HumanDroppedEvent     EventType = iota
	HumanGrabbedEvent     EventType = iota
	HumanRescuedEvent     EventType = iota
	HumanSavedEvent       EventType = iota
	HumanLandedEvent      EventType = iota
	MaterialiseEvent      EventType = iota
	PlayerExplodeEvent    EventType = iota
)

func (ev EventType) String() string {
	switch ev {

	case StartEvent:
		return "StartEvent"
	case FireBulletEvent:
		return "FireBulletEvent"
	case ExplodeEvent:
		return "ExplodeEvent"
	case LanderDieEvent:
		return "LanderDieEvent"
	case MutantSoundEvent:
		return "MutantSoundEvent"
	case HumanDieEvent:
		return "HumanDieEvent"
	case BaiterDieEvent:
		return "BaiterDieEvent"
	case BomberDieEvent:
		return "BomberDieEvent"
	case LanderClearedEvent:
		return "LanderClearedEvent"
	case PlayerDieEvent:
		return "PlayerDieEvent"
	case PlayerCollideEvent:
		return "PlayerCollideEvent"
	case PlayerFireEvent:
		return "PlayerFireEvent"
	case PlayerThrustEvent:
		return "PlayerThrustEvent"
	case PlayerStopThrustEvent:
		return "PlayerStopThrustEvent"
	case SmartBombEvent:
		return "SmartBombEvent"
	case SwarmerDieEvent:
		return "SwarmerDieEvent"
	case PodDieEvent:
		return "PodDieEvent"
	case HumanDroppedEvent:
		return "HumanDroppedEvent"
	case HumanGrabbedEvent:
		return "HumanGrabbedEvent"
	case HumanRescuedEvent:
		return "HumanRescuedEvent"
	case HumanSavedEvent:
		return "HumanSavedEvent"
	case HumanLandedEvent:
		return "HumanLandedEvent"
	case MaterialiseEvent:
		return "MaterialiseEvent"
	case PlayerExplodeEvent:
		return "PlayerExplodeEvent"
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
