package global

import "Def/types"

const (
	ScreenWidth  = 1600
	ScreenHeight = 1200
	MaxTPS       = 30
	WorldWidth   = ScreenWidth * 10
	ScreenTop    = 150
	LanderSpeed  = 2
)

var Blue = types.ColorF{R: 0, G: 0, B: 1, A: 1}
var Red = types.ColorF{R: 1, G: 0, B: 0, A: 1}
var Green = types.ColorF{R: 0, G: 1, B: 0, A: 1}
var Yellow = types.ColorF{R: 1, G: 1, B: 0, A: 1}
var Magenta = types.ColorF{R: 1, G: 0, B: 1, A: 1}

var Cols = []types.ColorF{Blue, Green, Yellow, Red, Magenta}
var HudCol int = 0
