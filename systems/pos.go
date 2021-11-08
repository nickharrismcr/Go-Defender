package systems

import (
	"Def/cmp"
	"Def/gl"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type PosSystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewPosSystem(active bool, engine types.IEngine) *PosSystem {
	f := NewFilter()
	f.Add(types.Pos)
	return &PosSystem{
		sysname: types.PosSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (pos *PosSystem) GetName() types.SystemName {
	return pos.sysname
}

func (pos *PosSystem) Update() {
	if !pos.active {
		return
	}
	for _, e := range pos.targets {
		if e.Active() && !e.Paused() {
			pos.process(e)
		}
	}
}

func (pos *PosSystem) Draw(screen *ebiten.Image) {}

func (pos *PosSystem) process(e types.IEntity) {
	poscmp := e.GetComponent(types.Pos).(*cmp.Pos)
	// "hidden"
	if poscmp.Y == 9999 {
		return
	}
	if poscmp.X < 0 {
		poscmp.X += gl.WorldWidth
	} else if poscmp.X > gl.WorldWidth {
		poscmp.X -= gl.WorldWidth
	}
	if poscmp.Y < gl.ScreenTop+20 {
		poscmp.Y = gl.ScreenTop + 20
	}
	if poscmp.Y > gl.ScreenHeight-50 {
		poscmp.Y = gl.ScreenHeight - 50
	}
	poscmp.X += poscmp.DX
	poscmp.Y += poscmp.DY

}

func (pos *PosSystem) Active() bool {
	return pos.active
}

func (pos *PosSystem) SetActive(active bool) {
	pos.active = active
}

func (pos *PosSystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := pos.targets[e.GetID()]; ok {
		return
	}
	for _, c := range pos.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", pos, e.GetID())
	pos.targets[e.GetID()] = e
}

func (pos *PosSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range pos.filter.Requires() {
		if !e.HasComponent(c) {

			logger.Debug("System %T removed entity %d ", pos, e.GetID())
			delete(pos.targets, e.GetID())
			return
		}
	}
}

func (s *PosSystem) RemoveEntity(e types.IEntity) {

	logger.Debug("System %T removed entity %d ", s, e.GetID())
	delete(s.targets, e.GetID())
}
