package systems

import (
	"Def/cmp"
	"Def/game"
	"Def/logger"
	"Def/types"
	"Def/util"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// implements ISystem

type LaserDrawSystem struct {
	sysname game.SystemName
	filter  *game.Filter
	active  bool
	engine  *game.Engine
	targets map[types.EntityID]*game.Entity
	img     *ebiten.Image
	opts    *ebiten.DrawImageOptions
}

func NewLaserDrawSystem(active bool, engine *game.Engine) *LaserDrawSystem {

	f := game.NewFilter()
	f.Add(types.LaserDraw)
	f.Add(types.Pos)
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)

	return &LaserDrawSystem{
		sysname: game.LaserDrawSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]*game.Entity),
		img:     img,
		opts:    &ebiten.DrawImageOptions{},
	}
}

func (lds *LaserDrawSystem) GetName() game.SystemName {
	return lds.sysname
}

func (lds *LaserDrawSystem) Update() {}

func (lds *LaserDrawSystem) Draw(screen *ebiten.Image) {
	if !lds.active {
		return
	}
	for _, e := range lds.targets {
		if e.Active() {
			lds.process(e, screen)
		}
	}
}

func (lds *LaserDrawSystem) process(e *game.Entity, screen *ebiten.Image) {

	ldc := e.GetComponent(types.LaserDraw).(*cmp.LaserDraw)
	lasmov := e.GetComponent(types.LaserMove).(*cmp.LaserMove)
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	c := ldc.Color
	lds.opts.ColorM.Reset()
	lds.opts.ColorM.Scale(c.R, c.G, c.B, c.A)
	lds.opts.GeoM.Reset()
	if pc.DX < 0 {
		lds.opts.GeoM.Scale(-lasmov.Length, 4)
	} else {
		lds.opts.GeoM.Scale(lasmov.Length, 4)
	}

	sx := util.ScreenX(pc.X)
	lds.opts.GeoM.Translate(sx, pc.Y)
	screen.DrawImage(lds.img, lds.opts)

	if ldc.Counter == 0 {
		ldc.Counter = 5
		var s float64 = 0
		for i := 0; i < 9; i += 2 {
			if pc.DX < 0 {
				s -= rand.Float64() * 300
			} else {
				s += rand.Float64() * 300
			}
			ldc.Black[i] = s
			ln := rand.Float64() * 50
			ldc.Black[i+1] = ln
		}
	}

	for i := 0; i < 9; i += 2 {
		lds.opts.ColorM.Reset()
		lds.opts.ColorM.Scale(0, 0, 0, 1)
		lds.opts.GeoM.Reset()
		lds.opts.GeoM.Scale(ldc.Black[i+1], 4)
		lds.opts.GeoM.Translate(sx+ldc.Black[i], pc.Y)
		screen.DrawImage(lds.img, lds.opts)
	}

	ldc.Counter--

	pc.X += pc.DX * 1.5
	if util.OffScreen(sx, pc.Y) {
		e.SetActive(false)
	}
}

func (lds *LaserDrawSystem) Active() bool {
	return lds.active
}

func (lds *LaserDrawSystem) SetActive(active bool) {
	lds.active = active
}

func (lds *LaserDrawSystem) AddEntityIfRequired(e *game.Entity) {
	if _, ok := lds.targets[e.Id]; ok {
		return
	}
	for _, c := range lds.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}
	logger.Debug("System %T added entity %d ", lds, e.Id)
	lds.targets[e.Id] = e
}

func (lds *LaserDrawSystem) RemoveEntityIfRequired(e *game.Entity) {
	for _, c := range lds.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			logger.Debug("System %T removed entity %d ", lds, e.Id)
			delete(lds.targets, e.Id)
			return
		}
	}
}
