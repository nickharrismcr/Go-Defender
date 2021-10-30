package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/gl"
	"Def/logger"
	"Def/types"
	"Def/util"
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

func (lms *LaserMoveSystem) process(laserEnt *game.Entity) {

	pe := laserEnt.GetEngine().GetEntity(gl.PlayerID)
	ppc := pe.GetComponent(types.Pos).(*cmp.Pos)
	psc := pe.GetComponent(types.Ship).(*cmp.Ship)
	lpc := laserEnt.GetComponent(types.Pos).(*cmp.Pos)
	lmc := laserEnt.GetComponent(types.LaserMove).(*cmp.LaserMove)
	// track player dx
	lpc.X += lpc.DX * (20 + math.Abs(ppc.DX))
	lmc.Length += 70

	var h2 float64 = 4
	y2 := lpc.Y
	x2 := util.ScreenX(lpc.X)
	w2 := lmc.Length
	if psc.Direction == -1 {
		x2 = x2 - lmc.Length
	}

	for _, v := range laserEnt.GetEngine().GetEntitiesWithComponent(types.Shootable) {
		tpc := v.GetComponent(types.Pos).(*cmp.Pos)
		x1 := util.ScreenX(tpc.X)
		y1 := tpc.Y
		if util.OffScreen(x1, tpc.Y) {
			continue
		}
		if v.HasComponent(types.Collide) {
			tcc := v.GetComponent(types.Collide).(*cmp.Collide)
			w1 := tcc.W
			h1 := tcc.H
			if util.Collide(x1, y1, w1, h1, x2, y2, w2, h2) {
				laserEnt.SetActive(false)
				laserEnt.GetEngine().Kill(v)
			}
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
