package update_systems

import (
	"Def/cmp"
	"Def/constants"
	"Def/game"
	"Def/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type PosSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[game.EntityID]*game.Entity
}

func NewPosSystem(active bool) *PosSystem {
	f := game.NewFilter()
	f.Add(cmp.PosType)
	return &PosSystem{
		sysname: game.PosSystem,
		active:  active,
		filter:  f,
		targets: make(map[game.EntityID]*game.Entity),
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
		if e.Active() {
			pos.process(e)
		}
	}
}

func (pos *PosSystem) Draw(screen *ebiten.Image) {}

func (pos *PosSystem) process(e *game.Entity) {
	poscmp := e.GetComponent(cmp.PosType).(*cmp.PosCmp)
	poscmp.X += poscmp.DX
	poscmp.Y += poscmp.DY

	if poscmp.X < 0 || poscmp.X > constants.ScreenWidth {
		poscmp.DX = -poscmp.DX
	}
	if poscmp.Y < 0 || poscmp.Y > constants.ScreenHeight {
		poscmp.DY = -poscmp.DY
	}
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
