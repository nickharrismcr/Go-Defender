package cmp

import "Def/types"

type LaserDraw struct {
	componentType types.CmpType
	Color         types.ColorF
}

func NewLaserDraw() *LaserDraw {

	return &LaserDraw{

		componentType: types.LaserDraw,
		Color:         types.ColorF{},
	}
}

func (pos *LaserDraw) Type() types.CmpType {
	return pos.componentType
}
