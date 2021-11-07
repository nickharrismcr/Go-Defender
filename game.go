package main

//TODO
//  levels
//  lives/bombs HUD

import (
	"Def/game"
	"errors"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine *game.Engine
}

func NewGame() *Game {
	game := &Game{
		engine: game.NewEngine(),
	}
	game.engine.Init()
	return game
}

func (g *Game) Update() error {

	if status := g.engine.Update(); status != game.OK {
		return errors.New(fmt.Sprintf("%d", status))
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}
