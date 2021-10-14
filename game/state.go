package game

import "FSM/cmp"

type IState interface {
	GetName() string
	Enter(ai *cmp.AICmp)
	Update(ai *cmp.AICmp)
}
