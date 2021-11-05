package cmp

import "Def/types"

type Pos struct {
	componentType types.CmpType
	X, Y, DX, DY  float64
	ScreenCoords  bool // is position screen rather than world,  used for score GFX
}

func NewPos(x, y, dx, dy float64) *Pos {

	return &Pos{
		X:             x,
		Y:             y,
		DX:            dx,
		DY:            dy,
		componentType: types.Pos,
	}
}

func (pos *Pos) Type() types.CmpType {
	return pos.componentType
}
