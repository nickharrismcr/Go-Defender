package game

import (
	"Def/global"
	"Def/types"
	"Def/util"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const MAX int = 200

// individual particle. system has a pool of these size = MAX
type particle struct {
	active              bool
	ticksToLive         int
	x, y, dx, dy, scale float64
	color               types.ColorF
	update              *func(p *particle)
	image               *ebiten.Image
	opts                *ebiten.DrawImageOptions
}

// particle system. only particles in activeList are processed.
// a call to Trigger moves inactive particles from the plist pool to the active list.
// when a particles ticksToLive expires it is taken out of the active list.
type ParticleSystem struct {
	plist      []*particle
	activeList []*particle
}

var update1 = func(p *particle) {
	p.x += p.dx
	p.y += p.dy
	p.dx /= 1.03
	p.dy /= 1.03
	p.dy += 0.1
	p.scale /= 1.01
	p.color.A /= 1.01
	if p.ticksToLive < 60 {
		p.color.B /= 1.03
		p.color.G /= 1.03
	}
}

var col types.ColorF

func new1(p *particle, x, y float64, col types.ColorF) {
	p.active = true
	p.ticksToLive = 90
	p.x = x
	p.y = y
	p.scale = 1
	dir := rand.Float64() * math.Pi * 2
	p.dx = math.Sin(dir)
	p.dy = math.Cos(dir)
	p.color = col
	speed := 2 + rand.Float64()*32
	p.dx *= speed
	p.dy *= speed
	p.update = &update1
}

// init with pool of MAX particles
func NewParticleSystem() *ParticleSystem {
	s := &ParticleSystem{}
	for i := 0; i < MAX; i++ {
		img := ebiten.NewImage(10, 10)
		img.Fill(color.White)
		p := &particle{
			active:      false,
			ticksToLive: 0,
			x:           0,
			y:           0,
			scale:       1,
			color:       types.ColorF{R: 1, G: 1, B: 1, A: 1},
			image:       img,
			opts:        &ebiten.DrawImageOptions{},
		}
		s.plist = append(s.plist, p)
	}
	s.activeList = []*particle{}

	return s
}

func (s *ParticleSystem) Trigger(x, y float64) {
	c := 0
	col = types.ColorF{R: 1, G: 1, B: 1, A: 1}
	col.Randomize()
	for _, p := range s.plist {
		if !p.active {
			c++
			if c > 100 {
				return
			}
			new1(p, x, y, col)
			s.activeList = append(s.activeList, p)
		}
	}
}

func (s *ParticleSystem) Update() {

	for i, p := range s.activeList {
		if i >= len(s.activeList) {
			return
		}
		p.ticksToLive--
		if p.ticksToLive == 0 || util.OffScreen(p.x, p.y) {
			p.active = false
			s.activeList = append(s.activeList[:i], s.activeList[i+1:]...)
		}
		(*p.update)(p)
	}
}

func (s *ParticleSystem) Draw(screen *ebiten.Image) {
	for _, p := range s.activeList {

		p.opts.GeoM.Reset()
		p.opts.GeoM.Scale(p.scale, p.scale)
		screenX := p.x - global.CameraX
		p.opts.GeoM.Translate(screenX, p.y)
		p.opts.ColorM.Reset()
		p.opts.ColorM.Scale(p.color.R, p.color.G, p.color.B, 1)

		screen.DrawImage(p.image, p.opts)
	}
}
