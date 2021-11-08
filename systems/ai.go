package systems

import (
	"Def/cmp"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type AISystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewAISystem(active bool, engine types.IEngine) *AISystem {
	f := NewFilter()
	f.Add(types.AI)
	return &AISystem{
		sysname: types.AISystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (ai *AISystem) GetName() types.SystemName {
	return ai.sysname
}

func (ai *AISystem) Update() {
	if !ai.active {
		return
	}
	for _, e := range ai.targets {
		if e.Active() && !e.Paused() {
			ai.process(e)
		}
	}
}

func (ai *AISystem) Draw(screen *ebiten.Image) {}

func (ai *AISystem) process(e types.IEntity) {
	aicmp := e.GetComponent(types.AI).(*cmp.AI)
	GetFSM(aicmp.FSMId).Update(aicmp, e)

}

func (ai *AISystem) Active() bool {
	return ai.active
}

func (ai *AISystem) SetActive(active bool) {
	ai.active = active
}

func (ai *AISystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := ai.targets[e.GetID()]; ok {
		return
	}
	for _, c := range ai.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", ai, e.GetID())
	ai.targets[e.GetID()] = e
}

func (ai *AISystem) RemoveEntity(e types.IEntity) {

	logger.Debug("System %T removed entity %d ", ai, e.GetID())
	delete(ai.targets, e.GetID())
}

func (ai *AISystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range ai.filter.Requires() {
		if !e.HasComponent(c) {
			logger.Debug("System %T removed entity %d ", ai, e.GetID())
			delete(ai.targets, e.GetID())
			return
		}
	}
}
