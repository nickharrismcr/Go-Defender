package main

import (
	"Def/constants"
	"Def/game"
	"fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count int
}

var engine *game.Engine

func (g *Game) Update() error {
	g.count++
	engine.Update()
	if g.count%60 == 0 {
		x := rand.Float64() * constants.ScreenWidth
		y := rand.Float64() * constants.ScreenHeight
		engine.TriggerPS(x, y)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	engine.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.CurrentTPS()))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}

func main() {

	engine = game.NewEngine()
	//engine.AddSystem(update_systems.NewPosSystem(true), game.UPDATE)
	//engine.AddSystem(draw_systems.NewDrawSystem(true), game.DRAW)

	game := &Game{
		count: 0,
	}

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
