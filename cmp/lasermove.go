package cmp

import "Def/types"

type LaserMove struct {
	componentType types.CmpType
	Color         types.ColorF
}

func NewLaserMove() *LaserMove {

	return &LaserMove{

		componentType: types.LaserMove,
		Color:         types.ColorF{},
	}
}

func (pos *LaserMove) Type() types.CmpType {
	return pos.componentType
}
