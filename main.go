package main

//TODO
//  intro
//  lives HUD update bug
//  player explode bug

import (
	"Def/gl"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	ebiten.SetFullscreen(true)
	ebiten.SetMaxTPS(gl.MaxTPS)

	app := NewApp()
	if err := ebiten.RunGame(app); err != nil {
	}

}
