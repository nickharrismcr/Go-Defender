package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type AISystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewAISystem(active bool, engine *game.Engine) *AISystem {
	f := game.NewFilter()
	f.Add(types.AI)
	return &AISystem{
		sysname: game.AISystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (ai *AISystem) GetName() game.SystemName {
	return ai.sysname
}

func (ai *AISystem) Update() {
	if !ai.active {
		return
	}
	for _, e := range ai.targets {
		if e.Active() {
			ai.process(e)
		}
	}
}

func (ai *AISystem) Draw(screen *ebiten.Image) {}

func (ai *AISystem) process(e *game.Entity) {
	aicmp := e.GetComponent(types.AI).(*cmp.AI)
	// TODO should fsms be held in AISystem?
	game.GetFSM(aicmp.FSMId).Update(aicmp, e)
}

func (ai *AISystem) Active() bool {
	return ai.active
}

func (ai *AISystem) SetActive(active bool) {
	ai.active = active
}

func (ai *AISystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := ai.targets[e.Id]; ok {
		return
	}
	for _, c := range ai.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", ai, e.Id)
	ai.targets[e.Id] = e
}

func (ai *AISystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range ai.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", ai, e.Id)
			delete(ai.targets, e.Id)
			return
		}
	}
}
