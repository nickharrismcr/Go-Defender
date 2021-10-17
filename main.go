package main

import (
	"Def/cmp"
	"Def/constants"
	"Def/game"
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"Def/draw_systems"
	"Def/update_systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count int
}

var engine *game.Engine

func (g *Game) Update() error {

	engine.Update()
	for i := 0; i < 5; i++ {
		g.count++
		engine.GetEntity(game.EntityID(g.count)).SetActive(false)
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

func AddBlob(engine *game.Engine, img *ebiten.Image) {

	e := game.NewEntity(engine)
	e.SetActive(true)
	pos := cmp.NewPos(rand.Float64()*float64(constants.ScreenWidth),
		rand.Float64()*float64(constants.ScreenHeight),
		rand.Float64()*4-4,
		rand.Float64()*4-4)
	e.AddComponent(pos)
	col := constants.ColorF{}
	col.Randomize()
	dr := cmp.NewDraw(img, col)
	dr.Scale = rand.Float64() + 1
	e.AddComponent(dr)
}

func main() {

	engine = game.NewEngine()
	engine.AddSystem(update_systems.NewPosSystem(true), game.UPDATE)
	engine.AddSystem(draw_systems.NewDrawSystem(true), game.DRAW)
	img := ebiten.NewImage(5, 5)
	img.Fill(color.White)

	for i := 0; i < 10000; i++ {
		AddBlob(engine, img)
	}

	game := &Game{}

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
