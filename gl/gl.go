package gl

import (
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Mute              = true //false
	ScreenWidth       = 1600
	ScreenHeight      = 1200
	MaxTPS            = 30
	WorldWidth        = ScreenWidth * 6
	ScreenTop         = 150
	LanderSpeed       = 5
	HumanSpeed        = 0.1
	BomberSpeed       = 3
	PlayerSpeedX      = 50
	PlayerSpeedY      = 15
	SwarmerSpeed      = 5
	WorldExplodeTicks = 350
	BaiterSpeed       = 80
)

var KeyMap = map[types.ActionType]ebiten.Key{
	types.Up:         ebiten.KeyQ,
	types.Down:       ebiten.KeyA,
	types.Reverse:    ebiten.KeySpace,
	types.Thrust:     ebiten.KeyEnter,
	types.Fire:       ebiten.KeyRightBracket,
	types.SmartBomb:  ebiten.KeyBackspace,
	types.HyperSpace: ebiten.KeyControlLeft,
}

var PlayerID types.EntityID
var PlayerLives = 3
var SmartBombs = 5
var ScoreCharId int

var Score int = 0

var Blue = types.ColorF{R: 0, G: 0, B: 1, A: 1}
var Red = types.ColorF{R: 1, G: 0, B: 0, A: 1}
var Green = types.ColorF{R: 0, G: 1, B: 0, A: 1}
var Yellow = types.ColorF{R: 1, G: 1, B: 0, A: 1}
var Magenta = types.ColorF{R: 1, G: 0, B: 1, A: 1}
var Cyan = types.ColorF{R: 0, G: 1, B: 1, A: 1}

var Cols = []types.ColorF{Blue, Green, Yellow, Red, Magenta}
var LaserCols = []types.ColorF{Green, Green, Green, Yellow, Yellow, Yellow, Red, Red, Red, Blue, Blue, Blue, Cyan, Cyan, Cyan}

var HudCol int = 0

var cameraX float64

func CameraX() float64 {
	return cameraX
}

func SetCameraX(x float64) {
	cameraX = x
}
