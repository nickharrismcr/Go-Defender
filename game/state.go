package game

import "Def/cmp"

type IState interface {
	GetName() string
	Enter(ai *cmp.AICmp, e cmp.ComponentGetter)
	Update(ai *cmp.AICmp, e cmp.ComponentGetter)
}
