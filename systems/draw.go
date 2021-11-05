package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/gl"
	"Def/logger"
	"Def/types"
	"Def/util"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type DrawSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
}

func NewDrawSystem(active bool, engine *game.Engine) *DrawSystem {
	f := game.NewFilter()
	f.Add(types.Draw)
	f.Add(types.Pos)
	return &DrawSystem{
		sysname: game.DrawSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
	}
}

func (ds *DrawSystem) GetName() game.SystemName {
	return ds.sysname
}

func (ds *DrawSystem) Update() {}

func (ds *DrawSystem) Draw(screen *ebiten.Image) {
	if !ds.active {
		return
	}
	for _, e := range ds.targets {
		if e.Active() {
			dc := e.GetComponent(types.Draw).(*cmp.Draw)
			if dc.Hide {
				continue
			}
			ds.process(dc, e, screen)
		}
	}
}

func (ds *DrawSystem) process(dc *cmp.Draw, e *game.Entity, screen *ebiten.Image) {

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	op := dc.Opts
	frames := dc.SpriteMap.Anim_frames
	fw, fh := dc.SpriteMap.Frame.W/frames, dc.SpriteMap.Frame.H

	screenx := util.ScreenX(pc.X) - float64(fw)/2
	if pc.ScreenCoords {
		screenx = pc.X
	}

	if util.OffScreen(screenx, pc.Y) {
		return
	}

	// move and scale
	op.GeoM.Reset()
	op.GeoM.Scale(dc.Scale, dc.Scale)
	if dc.FlipX {
		op.GeoM.Scale(-1, 1)
	}
	y := pc.Y - float64(fh)/2
	op.GeoM.Translate(screenx, y)

	// animation frame
	dc.Counter++
	if dc.Counter > dc.SpriteMap.Ticks_per_frame {
		dc.Counter = 0
		dc.Frame++
		if dc.Frame == frames {
			dc.Frame = 0
		}
	}
	sx, sy := dc.SpriteMap.Frame.X+dc.Frame*fw, dc.SpriteMap.Frame.Y
	si := dc.Image.SubImage(image.Rect(sx, sy, sx+fw, sy+fh)).(*ebiten.Image)

	if dc.Disperse == 0 {
		screen.DrawImage(si, op)
		if dc.Bomber {
			ds.DrawBomber(si, screenx, y, dc, op, screen)
		} else {
			ds.Cycle(dc, 1)
		}
	} else {
		ds.DrawDisperse(screenx, pc.Y, sx, sy, fw, fh, dc, op, screen)

	}

}

func (ds *DrawSystem) DrawBomber(si *ebiten.Image, screenx float64, y float64, dc *cmp.Draw, op *ebiten.DrawImageOptions, screen *ebiten.Image) {

	ds.Cycle(dc, 0.1)
	ds.Cycle(dc, 0.1)
	op.GeoM.Reset()
	op.GeoM.Scale(dc.Scale, dc.Scale)
	op.GeoM.Translate(screenx+5, y+5)
	screen.DrawImage(si, op)
	ds.Cycle(dc, 0.1)
	ds.Cycle(dc, 0.1)
	op.GeoM.Reset()
	op.GeoM.Scale(dc.Scale/2, dc.Scale/2)
	op.GeoM.Translate(screenx+11, y+11)
	screen.DrawImage(si, op)
}

func (ds *DrawSystem) DrawDisperse(x, y float64, sx, sy, fw, fh int, dc *cmp.Draw, op *ebiten.DrawImageOptions, screen *ebiten.Image) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			x := x + (float64(i-4) * dc.Disperse)
			y := y + (float64(j-4) * dc.Disperse)
			op.GeoM.Reset()
			op.GeoM.Scale(2-float64(i)/10, 2-float64(j)/10)
			op.GeoM.Translate(x, y)
			x1 := sx + i*(fw/9)
			x2 := x1 + fw/9
			y1 := sy + j*(fh/9)
			y2 := y1 + fh/9
			ssi := dc.Image.SubImage(image.Rect(x1, y1, x2, y2)).(*ebiten.Image)

			screen.DrawImage(ssi, op)
		}
	}
}

func (ds *DrawSystem) Cycle(drawcmp *cmp.Draw, v float64) {
	if drawcmp.Cycle {
		drawcmp.CycleIndex += v
		c := gl.Cols[int(drawcmp.CycleIndex)%5]
		drawcmp.Opts.ColorM.Reset()
		drawcmp.Opts.ColorM.Scale(c.R, c.G, c.B, c.A)
	}
}

func (ds *DrawSystem) Active() bool {
	return ds.active
}

func (ds *DrawSystem) SetActive(active bool) {
	ds.active = active
}

func (ds *DrawSystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := ds.targets[e.Id]; ok {
		return
	}
	for _, c := range ds.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", ds, e.Id)
	ds.targets[e.Id] = e
}

func (ds *DrawSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range ds.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", ds, e.Id)
			delete(ds.targets, e.Id)
			return
		}
	}
}
