package cmp

import (
	"Def/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawCmp struct {
	componentType CmpType
	Image         *ebiten.Image
	Opts          *ebiten.DrawImageOptions
	Color         constants.ColorF
	Scale         float64
}

func NewDraw(image *ebiten.Image, color constants.ColorF) *DrawCmp {

	return &DrawCmp{
		Image:         image,
		Opts:          &ebiten.DrawImageOptions{},
		Color:         color,
		componentType: DrawType,
		Scale:         1,
	}
}

func (Draw *DrawCmp) Type() CmpType {
	return Draw.componentType
}
