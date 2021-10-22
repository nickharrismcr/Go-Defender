package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/global"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

var sw = float64(global.ScreenWidth)
var sh = float64(global.ScreenHeight)
var rxs = sw * 0.25
var rxe = sw * 0.75
var rw = rxe - rxs
var ww = float64(global.WorldWidth)
var rsw = rw * (sw / ww)

// implements ISystem

type RadarDrawSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewRadarDrawSystem(active bool) *RadarDrawSystem {
	f := game.NewFilter()
	f.Add(types.RadarDraw)
	f.Add(types.Pos)
	return &RadarDrawSystem{
		sysname: game.RadarDrawSystem,
		active:  active,
		filter:  f,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (drawsys *RadarDrawSystem) GetName() game.SystemName {
	return drawsys.sysname
}

func (drawsys *RadarDrawSystem) Update() {}

func (drawsys *RadarDrawSystem) Draw(screen *ebiten.Image) {
	if !drawsys.active {
		return
	}
	for _, e := range drawsys.targets {
		if e.Active() {
			drawsys.process(e, screen)
		}
	}
}

func (drawsys *RadarDrawSystem) process(e *game.Entity, screen *ebiten.Image) {

	drawcmp := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	poscmp := e.GetComponent(types.Pos).(*cmp.Pos)
	op := drawcmp.Opts
	op.GeoM.Reset()

	var posx = ww/2 + poscmp.X - global.CameraX - sw/2
	if posx > ww {
		posx = posx - ww
	}
	if posx < 0 {
		posx = posx + ww
	}
	screenx := rxs + rw*(posx/ww)

	op.GeoM.Translate(screenx, poscmp.Y/10)
	c := drawcmp.Color
	op.ColorM.Scale(c.R, c.G, c.B, c.A)
	screen.DrawImage(drawcmp.Image, op)
}

func (drawsys *RadarDrawSystem) Active() bool {
	return drawsys.active
}

func (drawsys *RadarDrawSystem) SetActive(active bool) {
	drawsys.active = active
}

func (drawsys *RadarDrawSystem) AddEntityIfRequired(e *game.Entity) {
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

func (drawsys *RadarDrawSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range drawsys.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", drawsys, e.Id)
			delete(drawsys.targets, e.Id)
			return
		}
	}
}
