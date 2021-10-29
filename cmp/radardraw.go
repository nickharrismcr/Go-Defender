package cmp

import (
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type RadarDraw struct {
	componentType types.CmpType
	Image         *ebiten.Image
	Opts          *ebiten.DrawImageOptions
	Color         types.ColorF
	Cycle         bool
	CycleIndex    float32
	Hide          bool
}

func NewRadarDraw(image *ebiten.Image, color types.ColorF) *RadarDraw {

	return &RadarDraw{
		Image:         image,
		Opts:          &ebiten.DrawImageOptions{},
		Color:         color,
		componentType: types.RadarDraw,
		Cycle:         false,
		CycleIndex:    0,
	}
}

func (d *RadarDraw) Type() types.CmpType {
	return d.componentType
}
