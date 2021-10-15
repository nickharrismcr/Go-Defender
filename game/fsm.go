package game

import (
	"FSM/cmp"
	"fmt"
)

var fsmList []FSM
var fsmCount int

func init() {
	fsmList = []FSM{}
	fsmCount = 0
}

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

func (f FSM) Update(ai *cmp.AICmp, e *Entity) {

	if ai.NextStateName != ai.StateName {
		if ai.StateName != "" {
			if !f.statetree.ValidTransition(ai.StateName, ai.NextStateName) {
				panic(fmt.Sprintf("invalid transition %s -> %s", ai.StateName, ai.NextStateName))
			}
		}
		next_state, err := f.statetree.State(ai.NextStateName)
		if err != nil {
			panic(fmt.Sprint("No state defined in FSM %s for %s", f.name, ai.NextStateName))
		}
		next_state.Enter(ai, e)
		ai.StateName = ai.NextStateName
	}
	curr_state, err := f.statetree.State(ai.StateName)
	if err != nil {
		panic(fmt.Sprintf("no current state %s in FSM %s ", ai.StateName, f.name))
	}
	curr_state.Update(ai, e)
}
