package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/global"
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

func NewDrawSystem(active bool) *DrawSystem {
	f := game.NewFilter()
	f.Add(types.Draw)
	f.Add(types.Pos)
	return &DrawSystem{
		sysname: game.DrawSystem,
		active:  active,
		filter:  f,
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

func (drawsys *DrawSystem) process(e *game.Entity, screen *ebiten.Image) {

	drawcmp := e.GetComponent(types.Draw).(*cmp.Draw)
	poscmp := e.GetComponent(types.Pos).(*cmp.Pos)

	op := drawcmp.Opts
	op.GeoM.Reset()
	px := poscmp.X

	camx := e.GetEngine().GetCameraX()
	ww := float64(global.WorldWidth)
	sw := float64(global.ScreenWidth)
	translate := px - camx

	if px < (sw - (ww - camx)) {
		translate += ww
	}

	if util.OffScreen(translate, poscmp.Y) {
		return
	}

	op.GeoM.Translate(translate, poscmp.Y)
	frames := drawcmp.SpriteMap.Anim_frames
	drawcmp.Counter++
	if drawcmp.Counter > drawcmp.SpriteMap.Ticks_per_frame {
		drawcmp.Counter = 0
		drawcmp.Frame++
		if drawcmp.Frame == frames {
			drawcmp.Frame = 0
		}
	}

	fw, fh := drawcmp.SpriteMap.Frame.W/frames, drawcmp.SpriteMap.Frame.H
	sx, sy := drawcmp.SpriteMap.Frame.X+drawcmp.Frame*fw, drawcmp.SpriteMap.Frame.Y

	si := drawcmp.Image.SubImage(image.Rect(sx, sy, sx+fw, sy+fh)).(*ebiten.Image)
	if drawcmp.Disperse == 0 {
		screen.DrawImage(si, op)
	} else {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				x := translate + (float64(i-2) * drawcmp.Disperse)
				y := poscmp.Y + (float64(j-2) * drawcmp.Disperse)
				op.GeoM.Reset()
				op.GeoM.Translate(x, y)
				x1 := sx + i*(fw/9)
				x2 := x1 + fw/9
				y1 := sy + j*(fh/9)
				y2 := y1 + fh/9
				ssi := drawcmp.Image.SubImage(image.Rect(x1, y1, x2, y2)).(*ebiten.Image)
				screen.DrawImage(ssi, op)
			}
		}
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
