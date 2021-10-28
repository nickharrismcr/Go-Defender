package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"
	"Def/types"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type LaserMoveSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewLaserMoveSystem(active bool, engine *game.Engine) *LaserMoveSystem {

	f := game.NewFilter()
	f.Add(types.LaserMove)
	f.Add(types.Pos)

	return &LaserMoveSystem{
		sysname: game.LaserMoveSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (lms *LaserMoveSystem) GetName() game.SystemName {
	return lms.sysname
}

func (lms *LaserMoveSystem) Update() {
	if !lms.active {
		return
	}
	for _, e := range lms.targets {
		if e.Active() {
			lms.process(e)
		}
	}
}

func (lms *LaserMoveSystem) process(e *game.Entity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	for _, v := range e.GetEngine().GetEntitiesWithComponent(types.Shootable) {
		tpc := v.GetComponent(types.Pos).(*cmp.Pos)
		if math.Abs(tpc.Y-pc.Y) < 30 && tpc.X > pc.X && tpc.X < pc.X+1000 {
			e.SetActive(false)
			e.GetEngine().Kill(v)
		}
	}
}

func (lms *LaserMoveSystem) Draw(screen *ebiten.Image) {}

func (lms *LaserMoveSystem) Active() bool {
	return lms.active
}

func (lms *LaserMoveSystem) SetActive(active bool) {
	lms.active = active
}

func (lms *LaserMoveSystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := lms.targets[e.Id]; ok {
		return
	}
	for _, c := range lms.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", lms, e.Id)
	lms.targets[e.Id] = e
}

func (lms *LaserMoveSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range lms.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", lms, e.Id)
			delete(lms.targets, e.Id)
			return
		}
	}
}
