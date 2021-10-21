package cmp

import (
	"Def/constants"
	"Def/graphics"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	componentType types.CmpType
	Image         *ebiten.Image
	Opts          *ebiten.DrawImageOptions
	Color         constants.ColorF
	Scale         float64
	SpriteMap     graphics.GFXFrame
	Counter       int
	Frame         int
}

func NewDraw(image *ebiten.Image, smap graphics.GFXFrame, color constants.ColorF) *Draw {

	return &Draw{
		Image:         image,
		Opts:          &ebiten.DrawImageOptions{},
		Color:         color,
		componentType: types.Draw,
		Scale:         1,
		SpriteMap:     smap,
		Counter:       0,
		Frame:         0,
	}
}

func (d *Draw) Type() types.CmpType {
	return d.componentType
}
