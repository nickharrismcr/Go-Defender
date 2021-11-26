package main

import (
	"Def/game"
	"errors"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	engine *game.Engine
}

func NewApp() *App {
	app := &App{
		engine: game.NewEngine(),
	}
	app.engine.Init()
	return app
}

func (app *App) Update() error {

	if status := app.engine.Update(); status != game.OK {
		return errors.New(fmt.Sprintf("%d", status))
	}
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {

	//s := fmt.Sprintf("%d", gl.PlayerLives)
	//ebitenutil.DebugPrint(screen, s)

	app.engine.Draw(screen)
}

func (g *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}
