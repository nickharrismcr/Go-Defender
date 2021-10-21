package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"
	"Def/types"
	"image"

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
	f.Add(types.Draw)
	f.Add(types.Pos)
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

	drawcmp := e.GetComponent(types.Draw).(*cmp.Draw)
	poscmp := e.GetComponent(types.Pos).(*cmp.Pos)
	op := drawcmp.Opts
	op.GeoM.Reset()
	op.GeoM.Translate(poscmp.X, poscmp.Y)
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
	screen.DrawImage(drawcmp.Image.SubImage(image.Rect(sx, sy, sx+fw, sy+fh)).(*ebiten.Image), op)

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
