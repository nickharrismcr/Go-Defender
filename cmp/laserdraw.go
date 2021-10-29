package cmp

import "Def/types"

type LaserDraw struct {
	componentType types.CmpType
	Color         types.ColorF
	Black         [10]float64
	Counter       int
}

func NewLaserDraw() *LaserDraw {

	return &LaserDraw{

		componentType: types.LaserDraw,
		Color:         types.ColorF{},
		Black:         [10]float64{},
	}
}

func (pos *LaserDraw) Type() types.CmpType {
	return pos.componentType
}
