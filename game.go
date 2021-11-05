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

type (
	GameOver  struct{}
	Terminate struct{}
	LevelEnd  struct{}
)

func (e *GameOver) Error() string {
	return "Game over"
}
func (e *Terminate) Error() string {
	return "Terminated"
}
func (e *LevelEnd) Error() string {
	return "LevelEnd"
}

type Game struct {
	engine *game.Engine
}

func NewGame() *Game {
	game := &Game{
		engine: game.NewEngine(),
	}
	InitEngine(game.engine)
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
