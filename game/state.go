package game

import "Def/cmp"

type IState interface {
	GetName() string
	Enter(ai *cmp.AICmp, e *Entity)
	Update(ai *cmp.AICmp, e *Entity)
}
