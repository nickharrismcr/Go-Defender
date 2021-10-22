package game

import (
	"Def/cmp"
	"Def/types"
)

// interface for an FSM state used by the AI ECS
type IState interface {
	GetName() types.StateType
	Enter(ai *cmp.AI, e types.IEntity)
	Update(ai *cmp.AI, e types.IEntity)
}
