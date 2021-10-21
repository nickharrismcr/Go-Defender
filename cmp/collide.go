package cmp

import "Def/types"

type CollideCmp struct {
	componentType types.CmpType
}

func NewCollide() *CollideCmp {

	return &CollideCmp{
		componentType: types.Collide,
	}
}

func (pos *CollideCmp) Type() types.CmpType {
	return pos.componentType
}
