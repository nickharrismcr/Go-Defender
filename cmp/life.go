package cmp

import "Def/types"

type Life struct {
	componentType types.CmpType
	TicksToLive   int
}

func NewLife(toLive int) *Life {

	return &Life{
		TicksToLive:   toLive,
		componentType: types.Life,
	}
}

func (pos *Life) Type() types.CmpType {
	return pos.componentType
}
