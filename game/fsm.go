package game

import (
	"Def/cmp"
	"fmt"
)

var fsmList []FSM
var fsmCount int

func init() {
	fsmList = []FSM{}
	fsmCount = 0
}

// a finite state machine used by the AI ECS.
// each entity with an AI component will have its own state tree, this struct
// runs logic in the current AI state and handles transitions between states.
type FSM struct {
	statetree *StateTree
	name      string
}

func NewFSM(s *StateTree, name string) int {
	fsmList = append(fsmList, FSM{
		statetree: s,
		name:      name,
	})
	rv := fsmCount
	fsmCount++
	return rv
}

func GetFSM(id int) FSM {
	return fsmList[id]
}

func (f FSM) Update(ai *cmp.AI, e *Entity) {

	if ai.NextState != ai.State {
		//if ai.StateName != -1 {
		//if !f.statetree.ValidTransition(ai.StateName, ai.NextStateName) {
		//	panic(fmt.Sprintf("invalid transition %s -> %s", ai.StateName.String(), ai.NextStateName.String()))
		//}
		//}
		next_state, err := f.statetree.State(ai.NextState)
		if err != nil {
			panic(fmt.Sprintf("No state defined in FSM %s for %s", f.name, ai.NextState.String()))
		}
		next_state.Enter(ai, e)
		ai.State = ai.NextState
	}
	curr_state, err := f.statetree.State(ai.State)
	if err != nil {
		panic(fmt.Sprintf("no current state %s in FSM %s ", ai.State.String(), f.name))
	}
	curr_state.Update(ai, e)
}
