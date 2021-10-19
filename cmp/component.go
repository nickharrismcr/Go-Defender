package cmp

type CmpType int

const (
	AIType      CmpType = iota
	PosType     CmpType = iota
	DrawType    CmpType = iota
	CollideType CmpType = iota
)

func (t CmpType) String() string {
	switch t {
	case AIType:
		return "AI"
	case PosType:
		return "Pos"
	case DrawType:
		return "Draw"
	case CollideType:
		return "Collide"
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
