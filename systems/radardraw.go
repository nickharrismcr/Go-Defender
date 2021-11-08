package systems

import (
	"Def/cmp"
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
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewRadarDrawSystem(active bool, engine types.IEngine) *RadarDrawSystem {
	f := NewFilter()
	f.Add(types.RadarDraw)
	f.Add(types.Pos)
	lineImg.Fill(color.White)
	return &RadarDrawSystem{
		sysname: types.RadarDrawSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (drawsys *RadarDrawSystem) GetName() types.SystemName {
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

func (drawsys *RadarDrawSystem) process(e types.IEntity, screen *ebiten.Image) {

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

func (drawsys *RadarDrawSystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := drawsys.targets[e.GetID()]; ok {
		return
	}
	for _, c := range drawsys.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", drawsys, e.GetID())
	drawsys.targets[e.GetID()] = e
}

func (drawsys *RadarDrawSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range drawsys.filter.Requires() {
		if !e.HasComponent(c) {
			logger.Debug("System %T removed entity %d ", drawsys, e.GetID())
			delete(drawsys.targets, e.GetID())
			return
		}
	}
}

func (s *RadarDrawSystem) RemoveEntity(e types.IEntity) {

	logger.Debug("System %T removed entity %d ", s, e.GetID())
	delete(s.targets, e.GetID())
}
