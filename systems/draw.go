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

func (ds *DrawSystem) process(e *game.Entity, screen *ebiten.Image) {

	dc := e.GetComponent(types.Draw).(*cmp.Draw)

	if dc.Hide {
		return
	}

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	op := dc.Opts
	frames := dc.SpriteMap.Anim_frames
	fw, fh := dc.SpriteMap.Frame.W/frames, dc.SpriteMap.Frame.H
	screenx := util.ScreenX(pc.X) - float64(fw)/2
	if pc.Screen {
		screenx = pc.X
	}

	if util.OffScreen(screenx, pc.Y) {
		return
	}

	op.GeoM.Reset()
	op.GeoM.Scale(dc.Scale, dc.Scale)
	if dc.FlipX {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(screenx, pc.Y-float64(fh)/2)

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
			ds.Cycle(dc, 0.1)
			ds.Cycle(dc, 0.1)
			op.GeoM.Reset()
			op.GeoM.Scale(dc.Scale, dc.Scale)
			op.GeoM.Translate(screenx+7, pc.Y+7)
			screen.DrawImage(si, op)
			ds.Cycle(dc, 0.1)
			ds.Cycle(dc, 0.1)
			op.GeoM.Reset()
			op.GeoM.Scale(dc.Scale/2, dc.Scale/2)
			op.GeoM.Translate(screenx+11, pc.Y+11)
			screen.DrawImage(si, op)
		} else {
			ds.Cycle(dc, 1)
		}
	} else {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				x := screenx + (float64(i-4) * dc.Disperse)
				y := pc.Y + (float64(j-4) * dc.Disperse)
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

}

func (drawsys *DrawSystem) Cycle(drawcmp *cmp.Draw, v float64) {
	if drawcmp.Cycle {
		drawcmp.CycleIndex += v
		c := gl.Cols[int(drawcmp.CycleIndex)%5]
		drawcmp.Opts.ColorM.Reset()
		drawcmp.Opts.ColorM.Scale(c.R, c.G, c.B, c.A)
	}
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
