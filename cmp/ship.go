package cmp

import (
	"Def/global"
	"Def/types"
)

type Ship struct {
	componentType     types.CmpType
	Direction         float64
	ScreenOffset      float64
	ReversePressed    bool
	FirePressed       bool
	SmartBombPressed  bool
	HyperSpacePressed bool
}

func NewShip() *Ship {

	return &Ship{
		ScreenOffset:      global.ScreenWidth * 0.2,
		Direction:         1,
		componentType:     types.Ship,
		ReversePressed:    false,
		FirePressed:       false,
		SmartBombPressed:  false,
		HyperSpacePressed: false,
	}
}

func (pos *Ship) Type() types.CmpType {
	return pos.componentType
}
