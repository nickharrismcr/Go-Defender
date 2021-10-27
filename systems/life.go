package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type LifeSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewLifeSystem(active bool, engine *game.Engine) *LifeSystem {
	f := game.NewFilter()
	f.Add(types.Life)
	return &LifeSystem{
		sysname: game.LifeSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (pos *LifeSystem) GetName() game.SystemName {
	return pos.sysname
}

func (pos *LifeSystem) Update() {
	if !pos.active {
		return
	}
	for _, e := range pos.targets {
		if e.Active() {
			pos.process(e)
		}
	}
}

func (pos *LifeSystem) Draw(screen *ebiten.Image) {}

func (pos *LifeSystem) process(e *game.Entity) {
	cmp := e.GetComponent(types.Life).(*cmp.Life)
	cmp.TicksToLive--
	if cmp.TicksToLive < 0 {
		e.SetActive(false)
	}
}

func (pos *LifeSystem) Active() bool {
	return pos.active
}

func (pos *LifeSystem) SetActive(active bool) {
	pos.active = active
}

func (pos *LifeSystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := pos.targets[e.Id]; ok {
		return
	}
	for _, c := range pos.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", pos, e.Id)
	pos.targets[e.Id] = e
}

func (pos *LifeSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range pos.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", pos, e.Id)
			delete(pos.targets, e.Id)
			return
		}
	}
}
