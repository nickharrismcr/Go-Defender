package systems

import (
	"FSM/cmp"
	"FSM/game"
)

// implements ISystem

type AISystem struct {
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[game.EntityID]*game.Entity
}

func NewAISystem() *AISystem {
	f := game.NewFilter()
	f.Add(cmp.AIType)
	return &AISystem{

		filter:  f,
		targets: make(map[game.EntityID]*game.Entity),
	}
}

func (ai *AISystem) Update(dt float64) {
	for _, e := range ai.targets {
		ai.process(e, dt)
	}
}

func (ai *AISystem) process(e *game.Entity, dt float64) {
	aicmp := e.GetComponent(cmp.AIType).(*cmp.AICmp)
	game.GetFSM(aicmp.FSMId).Update(aicmp)

}

func (ai *AISystem) Active() bool {
	return ai.active
}

func (ai *AISystem) SetActive(active bool) {
	ai.active = active
}

func (ai *AISystem) AddEntityIfRequired(e *game.Entity) {
	for _, c := range ai.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	ai.targets[e.Id] = e
}

func (ai *AISystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range ai.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			delete(ai.targets, e.Id)
			return
		}
	}
}
