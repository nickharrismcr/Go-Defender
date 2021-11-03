// game/ECS functions
package game

import (
	"Def/state"
	"Def/types"
	"errors"
	"fmt"
)

// contains a map of FSM state structs and allowed transitions between them
type StateGraph struct {
	states      map[types.StateType]state.IState
	transitions map[types.StateType]map[types.StateType]bool
	lastAdded   types.StateType
}

func NewStateGraph() *StateGraph {
	return &StateGraph{
		states:      make(map[types.StateType]state.IState),
		transitions: make(map[types.StateType]map[types.StateType]bool),
		lastAdded:   -1,
	}
}

func (st *StateGraph) AddState(s state.IState) {
	st.states[s.GetName()] = s
	st.lastAdded = s.GetName()
}

func (st *StateGraph) State(name types.StateType) (state.IState, error) {
	if _, ok := st.states[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Invalid state name %s", name.String()))
	}
	return st.states[name], nil
}

func (st *StateGraph) AddTransition(from types.StateType, to types.StateType) {
	_, ok := st.transitions[from]
	if !ok {
		st.transitions[from] = make(map[types.StateType]bool)
	}
	st.transitions[from][to] = true
}

func (st *StateGraph) ValidTransition(from types.StateType, to types.StateType) bool {
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
