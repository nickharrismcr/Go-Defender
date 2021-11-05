package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/gl"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type PosSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewPosSystem(active bool, engine *game.Engine) *PosSystem {
	f := game.NewFilter()
	f.Add(types.Pos)
	return &PosSystem{
		sysname: game.PosSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (pos *PosSystem) GetName() game.SystemName {
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

func (pos *PosSystem) process(e *game.Entity) {
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

func (pos *PosSystem) AddEntityIfRequired(e *game.Entity) {
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

func (pos *PosSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range pos.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", pos, e.Id)
			delete(pos.targets, e.Id)
			return
		}
	}
}
