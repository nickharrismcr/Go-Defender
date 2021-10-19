package main

import (
	"Def/game"
	"Def/logger"
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ucount int
	dcount int
}

var engine *game.Engine

func (g *Game) Update() error {

	g.ucount++
	engine.Update()
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("escape pressed")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dcount++
	engine.Draw(screen)
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}

func main() {

	engine = game.NewEngine()

	InitGame(engine)

	gm := &Game{
		ucount: 0,
		dcount: 0,
	}

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	ebiten.SetFullscreen(true)
	ebiten.SetMaxTPS(30)
	if err := ebiten.RunGame(gm); err != nil {
		logger.Debug(">>> %d %d ", gm.ucount, gm.dcount)
	}

}
