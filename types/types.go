package types

import "math/rand"

type StateType int
type EntityType int
type EntityID int

const (
	Lander EntityType = iota
	Baiter EntityType = iota
	Bullet EntityType = iota
)

const (
	LanderMaterialise StateType = iota
	LanderSearch      StateType = iota
	LanderDrop        StateType = iota
	LanderGrab        StateType = iota
	LanderMutate      StateType = iota
	LanderDie         StateType = iota
	BaiterMaterialise StateType = iota
	BaiterHunt        StateType = iota
)

func (st StateType) String() string {
	switch st {
	case LanderMaterialise:
		return "Lander-Materialise"
	case LanderSearch:
		return "Lander-Search"
	case LanderDrop:
		return "Lander-Drop"
	case LanderGrab:
		return "Lander-Grab"
	case LanderMutate:
		return "Lander-Mutate"
	case LanderDie:
		return "Lander-Die"
	case BaiterMaterialise:
		return "Baiter-Materialise"
	case BaiterHunt:
		return "Baiter-Search"
	}
	return ""
}

type CmpType int

const (
	AI        CmpType = iota
	Pos       CmpType = iota
	Draw      CmpType = iota
	Collide   CmpType = iota
	Life      CmpType = iota
	RadarDraw CmpType = iota
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
	case RadarDraw:
		return ""
	}
	return ""
}

type ICmp interface {
	Type() CmpType
}

type IEngine interface {
	GetActiveEntityOfClass(EntityType) (IEntity, error)
	GetEntity(EntityID) IEntity
	MountainHeight(float64) float64
}

type IEntity interface {
	GetComponent(c CmpType) ICmp
	SetActive(bool)
	GetEngine() IEngine
	GetID() EntityID
	Active() bool
}

type ColorF struct {
	R, G, B, A float64
}

func (c *ColorF) Randomize() {
	c.R = rand.Float64()
	c.G = rand.Float64()
	c.B = rand.Float64()
}
