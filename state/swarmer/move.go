package swarmer

import (
	"Def/cmp"
	"Def/gl"
	"Def/types"
	"math"
	"math/rand"
)

// NB States should not contain entity state ;) they should act on cmp

type SwarmerMove struct {
	Name types.StateType
}

func NewSwarmerMove() *SwarmerMove {
	return &SwarmerMove{
		Name: types.SwarmerMove,
	}
}

func (s *SwarmerMove) GetName() types.StateType {
	return s.Name
}

func (s *SwarmerMove) Enter(ai *cmp.AI, e types.IEntity) {
	ai.Scratch = rand.Intn(10)
	ai.Val = rand.Float64() + 1
}

func (s *SwarmerMove) Update(ai *cmp.AI, e types.IEntity) {

	ai.Scratch += 1
	as := float64(ai.Scratch)
	gs := float64(gl.SwarmerSpeed) * ai.Val
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	ppc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos)

	pc.DY = 20 * math.Sin(as/7) * rand.Float64()
	pc.DX = 20 * math.Cos(as/7) * rand.Float64()

	if pc.X > ppc.X {
		pc.DX -= gs
	} else {
		pc.DX += gs
	}
	if pc.Y > ppc.Y {
		pc.DY -= gs
	} else {
		pc.DY += gs
	}
}
