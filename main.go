package main

//TODO
//  levels
//  lives/bombs HUD

import (
	"Def/game"
	"Def/gl"
	"errors"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

var engine *game.Engine

func (g *Game) Update() error {

	if !engine.Update() {
		return errors.New("end level")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	engine = game.NewEngine()
	InitGame(engine)
	gm := &Game{}
	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	//ebiten.SetFullscreen(true)
	ebiten.SetMaxTPS(gl.MaxTPS)
	if err := ebiten.RunGame(gm); err != nil {
	}

}
