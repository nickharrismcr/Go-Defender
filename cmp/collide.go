package cmp

import "Def/types"

type Collide struct {
	componentType types.CmpType
	W, H          float64
}

func NewCollide(w, h int) *Collide {

	return &Collide{
		componentType: types.Collide,
		W:             float64(w),
		H:             float64(h),
	}
}

func (pos *Collide) Type() types.CmpType {
	return pos.componentType
}
