package main

//TODO
//  mountain explode
//  sound
//  levels
//  baiter

import (
	"Def/game"
	"Def/gl"
	"Def/logger"
	"errors"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ucount int
	dcount int
}

var engine *game.Engine

func (g *Game) Update() error {

	engine.Update()
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("escape pressed")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dcount++
	engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	engine = game.NewEngine()
	InitGame(engine)

	gm := &Game{
		ucount: 0,
		dcount: 0,
	}

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	//ebiten.SetFullscreen(true)
	ebiten.SetMaxTPS(gl.MaxTPS)
	if err := ebiten.RunGame(gm); err != nil {
		logger.Debug(">>> %d %d ", gm.ucount, gm.dcount)
	}

}
