package cmp

type PosCmp struct {
	componentType CmpType
	X, Y, DX, DY  float64
}

func NewPos(x, y, dx, dy float64) *PosCmp {

	return &PosCmp{
		X:             x,
		Y:             y,
		DX:            dx,
		DY:            dy,
		componentType: PosType,
	}
}

func (pos *PosCmp) Type() CmpType {
	return pos.componentType
}
