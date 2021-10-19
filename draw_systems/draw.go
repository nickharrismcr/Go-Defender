package draw_systems

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type DrawSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[game.EntityID]*game.Entity
}

func NewDrawSystem(active bool) *DrawSystem {
	f := game.NewFilter()
	f.Add(cmp.DrawType)
	f.Add(cmp.PosType)
	return &DrawSystem{
		sysname: game.DrawSystem,
		active:  active,
		filter:  f,
		targets: make(map[game.EntityID]*game.Entity),
	}
}

func (drawsys *DrawSystem) GetName() game.SystemName {
	return drawsys.sysname
}

func (drawsys *DrawSystem) Update() {}

func (drawsys *DrawSystem) Draw(screen *ebiten.Image) {
	if !drawsys.active {
		return
	}
	for _, e := range drawsys.targets {
		if e.Active() {
			drawsys.process(e, screen)
		}
	}
}

func (drawsys *DrawSystem) process(e *game.Entity, screen *ebiten.Image) {

	drawcmp := e.GetComponent(cmp.DrawType).(*cmp.DrawCmp)
	poscmp := e.GetComponent(cmp.PosType).(*cmp.PosCmp)
	drawcmp.Opts.GeoM.Reset()
	sx, sy := drawcmp.Image.Size()
	drawcmp.Opts.GeoM.Translate(-float64(sx)/2, -float64(sy)/2)
	drawcmp.Opts.GeoM.Translate(poscmp.X, poscmp.Y)
	drawcmp.Opts.GeoM.Scale(drawcmp.Scale, drawcmp.Scale)
	drawcmp.Opts.ColorM.Reset()
	drawcmp.Opts.ColorM.Scale(drawcmp.Color.R, drawcmp.Color.G, drawcmp.Color.B, 1)

	screen.DrawImage(drawcmp.Image, drawcmp.Opts)

}

func (drawsys *DrawSystem) Active() bool {
	return drawsys.active
}

func (drawsys *DrawSystem) SetActive(active bool) {
	drawsys.active = active
}

func (drawsys *DrawSystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := drawsys.targets[e.Id]; ok {
		return
	}
	for _, c := range drawsys.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", drawsys, e.Id)
	drawsys.targets[e.Id] = e
}

func (drawsys *DrawSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range drawsys.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", drawsys, e.Id)
			delete(drawsys.targets, e.Id)
			return
		}
	}
}