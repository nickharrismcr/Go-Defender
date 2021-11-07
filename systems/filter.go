package systems

import "Def/types"

type Filter struct {
	list []types.CmpType
}

func NewFilter() *Filter {
	return &Filter{
		list: []types.CmpType{},
	}
}

func (f *Filter) Add(ct types.CmpType) {
	f.list = append(f.list, ct)
}

func (f *Filter) Requires() []types.CmpType {
	return f.list
}

func (f *Filter) NeedsComponent(ct types.CmpType) bool {
	for _, c := range f.list {
		if ct == c {
			return true
		}
	}
	return false
}
