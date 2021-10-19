package main

import (
	"Def/game"
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	count int
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
	engine.Draw(screen)
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}

func main() {

	engine = game.NewEngine()

	InitGame(engine)

	app := &Game{
		count: 0,
	}

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}

}
