package types

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

type EntityGetter interface {
	GetComponent(c CmpType) ICmp
	SetActive(bool)
}
