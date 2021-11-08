package systems

import (
	"Def/cmp"

	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type AnimateSystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewAnimateSystem(active bool, engine types.IEngine) *AnimateSystem {
	f := NewFilter()
	f.Add(types.Draw)
	return &AnimateSystem{
		sysname: types.AnimateSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (ds *AnimateSystem) GetName() types.SystemName {
	return ds.sysname
}

func (ds *AnimateSystem) Draw(screen *ebiten.Image) {}

func (ds *AnimateSystem) Update() {
	if !ds.active {
		return
	}
	for _, e := range ds.targets {
		if e.Active() {
			dc := e.GetComponent(types.Draw).(*cmp.Draw)
			if dc.Hide {
				continue
			}
			ds.process(dc, e)
		}
	}
}

func (ds *AnimateSystem) process(dc *cmp.Draw, e types.IEntity) {

	frames := dc.SpriteMap.Anim_frames

	// animation frame
	dc.Counter++
	if dc.Counter > dc.SpriteMap.Ticks_per_frame {
		dc.Counter = 0
		dc.Frame++
		if dc.Frame == frames {
			dc.Frame = 0
		}
	}

}

func (ds *AnimateSystem) Active() bool {
	return ds.active
}

func (ds *AnimateSystem) SetActive(active bool) {
	ds.active = active
}

func (ds *AnimateSystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := ds.targets[e.GetID()]; ok {
		return
	}
	for _, c := range ds.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", ds, e.GetID())
	ds.targets[e.GetID()] = e
}

func (ds *AnimateSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range ds.filter.Requires() {
		if !e.HasComponent(c) {

			logger.Debug("System %T removed entity %d ", ds, e.GetID())
			delete(ds.targets, e.GetID())
			return
		}
	}
}

func (s *AnimateSystem) RemoveEntity(e types.IEntity) {

	logger.Debug("System %T removed entity %d ", s, e.GetID())
	delete(s.targets, e.GetID())
}
