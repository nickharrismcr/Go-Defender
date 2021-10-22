// game/ECS functions
package game

import (
	"Def/state"
	"Def/types"
	"errors"
	"fmt"
)

// contains a tree of FSM state structs and allowed transitions between them
type StateTree struct {
	states      map[types.StateType]state.IState
	transitions map[types.StateType]map[types.StateType]bool
	last_added  types.StateType
}

func NewStateTree() *StateTree {
	return &StateTree{
		states:      make(map[types.StateType]state.IState),
		transitions: make(map[types.StateType]map[types.StateType]bool),
		last_added:  -1,
	}
}

func (st *StateTree) AddState(s state.IState) {
	st.states[s.GetName()] = s
	st.last_added = s.GetName()
}

func (st *StateTree) State(name types.StateType) (state.IState, error) {
	if _, ok := st.states[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Invalid state name %s", name.String()))
	}
	return st.states[name], nil
}

func (st *StateTree) AddTransition(from types.StateType, to types.StateType) {
	_, ok := st.transitions[from]
	if !ok {
		st.transitions[from] = make(map[types.StateType]bool)
	}
	st.transitions[from][to] = true
}

func (st *StateTree) ValidTransition(from types.StateType, to types.StateType) bool {
	if from == to {
		return true
	}
	_, ok := st.transitions[from]
	if ok {
		_, ok := st.transitions[from][to]
		return ok
	}
	return false
}
