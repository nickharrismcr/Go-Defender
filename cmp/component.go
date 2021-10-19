package cmp

type CmpType int

const (
	AIType   CmpType = iota
	PosType  CmpType = iota
	DrawType CmpType = iota
)

func (t CmpType) String() string {
	switch t {
	case AIType:
		return "AI"
	case PosType:
		return "Pos"
	case DrawType:
		return "Draw"
	}
	return ""
}

type ICmp interface {
	Type() CmpType
}

type ComponentGetter interface {
	GetComponent(c CmpType) ICmp
}
