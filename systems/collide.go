package systems

import (
	"Def/cmp"
	"Def/event"
	"Def/game"
	"Def/gl"
	"Def/logger"
	"Def/types"
	"Def/util"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

// system for detecting collisions between entities with the Collide component
type CollideSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewCollideSystem(active bool, engine *game.Engine) *CollideSystem {
	f := game.NewFilter()
	f.Add(types.Collide)
	return &CollideSystem{
		sysname: game.CollideSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (cs *CollideSystem) GetName() game.SystemName {
	return cs.sysname
}

func (cs *CollideSystem) Update() {
	if !cs.active {
		return
	}
	pe := cs.engine.GetEntities()[gl.PlayerID]
	for _, e := range cs.targets {
		if e.Active() {
			cs.process(e, pe)
		}
	}
}

func (cs *CollideSystem) Draw(screen *ebiten.Image) {}

func (cs *CollideSystem) process(e *game.Entity, player *game.Entity) {

	ep := e.GetComponent(types.Pos).(*cmp.Pos)
	ec := e.GetComponent(types.Collide).(*cmp.Collide)
	ppos := player.GetComponent(types.Pos).(*cmp.Pos)
	psh := player.GetComponent(types.Ship).(*cmp.Ship)
	if util.Collide(ppos.X, ppos.Y, psh.W, psh.H, ep.X, ep.Y, ec.W, ec.H) {
		ev := event.NewPlayerCollide(e)
		event.NotifyEvent(ev)
	}
}

func (cs *CollideSystem) Active() bool {
	return cs.active
}

func (cs *CollideSystem) SetActive(active bool) {
	cs.active = active
}

func (cs *CollideSystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := cs.targets[e.Id]; ok {
		return
	}
	for _, c := range cs.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", cs, e.Id)
	cs.targets[e.Id] = e
}

func (cs *CollideSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range cs.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", cs, e.Id)
			delete(cs.targets, e.Id)
			return
		}
	}
}
