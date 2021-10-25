package systems

import (
	"Def/game"
	"Def/logger"
	"Def/types"

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

func NewCollideSystem(active bool) *CollideSystem {
	f := game.NewFilter()
	f.Add(types.Collide)
	return &CollideSystem{
		sysname: game.CollideSystem,
		active:  active,
		filter:  f,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (pos *CollideSystem) GetName() game.SystemName {
	return pos.sysname
}

func (pos *CollideSystem) Update() {
	if !pos.active {
		return
	}
	for _, e := range pos.targets {
		if e.Active() {
			pos.process(e)
		}
	}
}

func (pos *CollideSystem) Draw(screen *ebiten.Image) {}

func (pos *CollideSystem) process(e *game.Entity) {
	/* 	for _, te := range pos.targets {
		if te.Active() && te.Id != e.Id && te.Class != e.Class {

			ep := e.GetComponent(types.Pos).(*cmp.Pos)
			tep := te.GetComponent(types.Pos).(*cmp.Pos)
			if math.Abs(ep.X-tep.X) < 25 && math.Abs(ep.Y-tep.Y) < 25 {
				ev := event.NewExplode(ep)
				event.NotifyEvent(ev)
				e.SetActive(false)
				te.SetActive(false)
				ed := event.NewEntityDie(e)
				event.NotifyEvent(ed)
				ed = event.NewEntityDie(te)
				event.NotifyEvent(ed)
			}
		}
	} */

}

func (pos *CollideSystem) Active() bool {
	return pos.active
}

func (pos *CollideSystem) SetActive(active bool) {
	pos.active = active
}

func (pos *CollideSystem) AddEntityIfRequired(e *game.Entity) {
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

func (pos *CollideSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range pos.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", pos, e.Id)
			delete(pos.targets, e.Id)
			return
		}
	}
}
