package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/global"
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

	playEnt := laserEnt.GetEngine().GetEntity(global.PlayerID)
	playPosCmp := playEnt.GetComponent(types.Pos).(*cmp.Pos)
	playShipCmp := playEnt.GetComponent(types.Ship).(*cmp.Ship)

	laspos := laserEnt.GetComponent(types.Pos).(*cmp.Pos)
	// track player dx
	laspos.X += laspos.DX * (20 + math.Abs(playPosCmp.DX))

	for _, v := range laserEnt.GetEngine().GetEntitiesWithComponent(types.Shootable) {
		tpc := v.GetComponent(types.Pos).(*cmp.Pos)
		stpcx := util.ScreenX(tpc.X)
		if util.OffScreen(stpcx, tpc.Y) {
			continue
		}
		if v.HasComponent(types.Collide) {
			pcc := v.GetComponent(types.Collide).(*cmp.Collide)
			if math.Abs(tpc.Y-laspos.Y) < pcc.H/2 &&
				stpcx > util.ScreenX(laspos.X) &&
				stpcx < util.ScreenX(laspos.X+(2000*playShipCmp.Direction)) {
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
