package cmp

import "Def/types"

type LaserMove struct {
	componentType types.CmpType
}

func NewLaserMove() *LaserMove {

	return &LaserMove{

		componentType: types.LaserMove,
	}
}

func (pos *LaserMove) Type() types.CmpType {
	return pos.componentType
}
