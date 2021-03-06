package game

import (
	"Def/gl"
	"Def/types"
	"Def/util"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const MAXSTARS int = 200

var starsImg *ebiten.Image

// background stars
type Stars struct {
	plist  []*star
	engine *Engine
	active bool
}

// individual star. system has a pool of these size = MAXSTARS
type star struct {
	active      bool
	ticksToLive int
	x, y        float64
	color       types.ColorF
	image       *ebiten.Image
	opts        *ebiten.DrawImageOptions
}

// init with pool of MAXSTARS stars
func NewStars(engine *Engine) *Stars {
	starsImg := ebiten.NewImage(4, 4)
	starsImg.Fill(color.White)
	s := &Stars{
		engine: engine,
		active: true,
	}
	for i := 0; i < MAXSTARS; i++ {

		p := &star{
			active:      false,
			ticksToLive: rand.Intn(30),
			x:           rand.Float64() * gl.WorldWidth,
			y:           gl.ScreenTop + rand.Float64()*(gl.ScreenHeight/2),
			color:       gl.Cols[rand.Intn(5)],
			image:       starsImg,
			opts:        &ebiten.DrawImageOptions{},
		}
		s.plist = append(s.plist, p)
	}
	return s
}
func (s *Stars) SetActive(b bool) {
	s.active = b
}

func (s *Stars) Update() {

	if !s.active {
		return
	}
	for _, p := range s.plist {
		p.ticksToLive--
		if p.ticksToLive == 0 {
			p.x = rand.Float64() * gl.WorldWidth
			p.y = gl.ScreenTop + rand.Float64()*(gl.ScreenHeight/2)
			p.ticksToLive = rand.Intn(60) + 30
		}
	}
}

func (s *Stars) Draw(screen *ebiten.Image) {

	if !s.active {
		return
	}
	for _, p := range s.plist {

		screenX := p.x - gl.CameraX()/4
		if util.OffScreen(screenX, p.y) {
			continue
		}
		p.opts.GeoM.Reset()
		p.opts.GeoM.Translate(screenX, p.y)
		p.opts.ColorM.Reset()
		p.opts.ColorM.Scale(p.color.R, p.color.G, p.color.B, 1)
		screen.DrawImage(p.image, p.opts)
	}
}
