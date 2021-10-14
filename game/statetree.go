package game

import (
	"errors"
	"fmt"
)

type StateTree struct {
	states      map[string]IState
	transitions map[string]map[string]bool
	last_added  string
}

func NewStateTree() *StateTree {
	return &StateTree{
		states:      make(map[string]IState),
		transitions: make(map[string]map[string]bool),
		last_added:  "",
	}
}

func (st *StateTree) AddState(s IState) {
	st.states[s.GetName()] = s
	st.last_added = s.GetName()
}

func (st *StateTree) State(name string) (IState, error) {
	if _, ok := st.states[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Invalid state name %s", name))
	}
	return st.states[name], nil
}

func (st *StateTree) AddTransition(from string, to string) {
	_, ok := st.transitions[from]
	if !ok {
		st.transitions[from] = make(map[string]bool)
	}
	st.transitions[from][to] = true
}

func (st *StateTree) ValidTransition(from string, to string) bool {
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
