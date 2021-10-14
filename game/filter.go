package game

import "FSM/cmp"

type Filter struct {
	list []cmp.CmpType
}

func NewFilter() *Filter {
	return &Filter{
		list: []cmp.CmpType{},
	}
}

func (f *Filter) Add(ct cmp.CmpType) {
	f.list = append(f.list, ct)
}

func (f *Filter) Requires() []cmp.CmpType {
	return f.list
}

func (f *Filter) NeedsComponent(ct cmp.CmpType) bool {
	for _, c := range f.list {
		if ct == c {
			return true
		}
	}
	return false
}
