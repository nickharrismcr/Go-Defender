package cmp

type CmpType int

const (
	AIType CmpType = iota
)

func (t CmpType) String() string {
	switch t {
	case 0:
		return "AI"
	}
	return ""
}

type ICmp interface {
	Type() CmpType
}
