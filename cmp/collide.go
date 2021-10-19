package cmp

type CollideCmp struct {
	componentType CmpType
}

func NewCollide() *CollideCmp {

	return &CollideCmp{
		componentType: CollideType,
	}
}

func (pos *CollideCmp) Type() CmpType {
	return pos.componentType
}
