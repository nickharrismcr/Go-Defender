package constants

import "math/rand"

const (
	ScreenWidth  = 320 * 5
	ScreenHeight = 240 * 5
	MaxTPS       = 30
)

type ColorF struct {
	R, G, B, A float64
}

func (c *ColorF) Randomize() {
	c.R = rand.Float64()
	c.G = rand.Float64()
	c.B = rand.Float64()
}
