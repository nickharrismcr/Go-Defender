package states

import (
	"Def/cmp"
	"Def/types"
)

// interface for an FSM state used by the AI ECS
type IState interface {
	GetName() types.StateType
	Enter(ai *cmp.AICmp, e types.EntityGetter)
	Update(ai *cmp.AICmp, e types.EntityGetter)
}
