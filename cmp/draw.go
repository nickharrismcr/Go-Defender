package cmp

import (
	"Def/constants"
	"Def/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawCmp struct {
	componentType CmpType
	Image         *ebiten.Image
	Opts          *ebiten.DrawImageOptions
	Color         constants.ColorF
	Scale         float64
	SpriteMap     graphics.GFXFrame
	Counter       int
	Frame         int
}

func NewDraw(image *ebiten.Image, smap graphics.GFXFrame, color constants.ColorF) *DrawCmp {

	return &DrawCmp{
		Image:         image,
		Opts:          &ebiten.DrawImageOptions{},
		Color:         color,
		componentType: DrawType,
		Scale:         1,
		SpriteMap:     smap,
		Counter:       0,
		Frame:         0,
	}
}

func (Draw *DrawCmp) Type() CmpType {
	return Draw.componentType
}
