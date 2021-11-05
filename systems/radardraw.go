package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/gl"
	"Def/logger"
	"Def/types"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var sw = float64(gl.ScreenWidth)
var sh = float64(gl.ScreenHeight)
var st = float64(gl.ScreenTop)
var rxs = sw * 0.25
var rxe = sw * 0.75
var rw = rxe - rxs
var ww = float64(gl.WorldWidth)
var rsw = rw * (sw / ww)

var lineImg = ebiten.NewImage(1, 1)
var lineOpts = &ebiten.DrawImageOptions{}

// implements ISystem

type RadarDrawSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewRadarDrawSystem(active bool, engine *game.Engine) *RadarDrawSystem {
	f := game.NewFilter()
	f.Add(types.RadarDraw)
	f.Add(types.Pos)
	lineImg.Fill(color.White)
	return &RadarDrawSystem{
		sysname: game.RadarDrawSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (drawsys *RadarDrawSystem) GetName() game.SystemName {
	return drawsys.sysname
}

func (drawsys *RadarDrawSystem) Update() {}

func (drawsys *RadarDrawSystem) HUD(screen *ebiten.Image) {

	col := gl.LevelCol()

	lineOpts.ColorM.Scale(col.R, col.G, col.B, col.A)
	lineOpts.GeoM.Reset()
	lineOpts.GeoM.Scale(sw, 2)
	lineOpts.GeoM.Translate(0, st)
	screen.DrawImage(lineImg, lineOpts)
	lineOpts.GeoM.Reset()
	lineOpts.GeoM.Scale(2, st)
	lineOpts.GeoM.Translate(rxs, 0)
	screen.DrawImage(lineImg, lineOpts)
	lineOpts.GeoM.Reset()
	lineOpts.GeoM.Scale(2, st)
	lineOpts.GeoM.Translate(rxe, 0)
	screen.DrawImage(lineImg, lineOpts)

	lineOpts.ColorM.Reset()
	lineOpts.ColorM.Scale(1, 1, 1, 1)
	lineOpts.GeoM.Reset()
	lineOpts.GeoM.Scale(2, st)
	lineOpts.GeoM.Translate(sw/2-rsw/2, 0)
	screen.DrawImage(lineImg, lineOpts)
	lineOpts.GeoM.Reset()
	lineOpts.GeoM.Scale(2, st)
	lineOpts.GeoM.Translate(sw/2+rsw/2, 0)
	screen.DrawImage(lineImg, lineOpts)

}

func (drawsys *RadarDrawSystem) Draw(screen *ebiten.Image) {
	if !drawsys.active {
		return
	}

	drawsys.HUD(screen)
	for _, e := range drawsys.targets {
		if e.Active() {
			drawsys.process(e, screen)
		}
	}
}

func (drawsys *RadarDrawSystem) process(e *game.Entity, screen *ebiten.Image) {

	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	if rdc.Hide {
		return
	}

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	op := rdc.Opts
	op.GeoM.Reset()

	var posx = ww/2 + pc.X - gl.CameraX() - sw/2
	if posx > ww {
		posx = posx - ww
	}
	if posx < 0 {
		posx = posx + ww
	}
	screenx := rxs + rw*(posx/ww)

	if rdc.Cycle {
		rdc.CycleIndex += 0.4
		rdc.Color = gl.Cols[int(rdc.CycleIndex)%5]
		op.ColorM.Reset()
	}

	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(screenx, pc.Y*(st/sh))
	c := rdc.Color
	op.ColorM.Reset()
	op.ColorM.Scale(c.R, c.G, c.B, c.A)
	screen.DrawImage(rdc.Image, op)
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
