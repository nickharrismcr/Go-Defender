package types

import "math/rand"

type StateType int
type EntityType int

const (
	Lander EntityType = iota
	Baiter EntityType = iota
	Bullet EntityType = iota
)

const (
	LanderSearch StateType = iota
	BaiterSearch StateType = iota
)

func (st StateType) String() string {
	switch st {
	case LanderSearch:
		return "Lander-Search"
	case BaiterSearch:
		return "Baiter-Search"
	}
	return ""
}

type CmpType int

const (
	AI      CmpType = iota
	Pos     CmpType = iota
	Draw    CmpType = iota
	Collide CmpType = iota
	Life    CmpType = iota
)

func (t CmpType) String() string {
	switch t {
	case AI:
		return "AI"
	case Pos:
		return "Pos"
	case Draw:
		return "Draw"
	case Collide:
		return "Collide"
	case Life:
		return "Life"
	}
	return ""
}

type ICmp interface {
	Type() CmpType
}

type IEngine interface {
	GetActiveEntityOfClass(EntityType) (IEntity, error)
}

type IEntity interface {
	GetComponent(c CmpType) ICmp
	SetActive(bool)
	GetEngine() IEngine
}

type ColorF struct {
	R, G, B, A float64
}

func (c *ColorF) Randomize() {
	c.R = rand.Float64()
	c.G = rand.Float64()
	c.B = rand.Float64()
}
