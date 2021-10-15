package game

import "FSM/cmp"

type IState interface {
	GetName() string
	Enter(ai *cmp.AICmp, e *Entity)
	Update(ai *cmp.AICmp, e *Entity)
}
