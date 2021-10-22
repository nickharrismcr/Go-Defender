package cmp

import "Def/types"

type Collide struct {
	componentType types.CmpType
}

func NewCollide() *Collide {

	return &Collide{
		componentType: types.Collide,
	}
}

func (pos *Collide) Type() types.CmpType {
	return pos.componentType
}
