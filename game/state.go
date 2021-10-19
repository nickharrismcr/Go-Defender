package game

import "Def/cmp"

// interface for an FSM state used by the AI ECS
type IState interface {
	GetName() string
	Enter(ai *cmp.AICmp, e cmp.EntityGetter)
	Update(ai *cmp.AICmp, e cmp.EntityGetter)
}
