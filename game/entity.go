package game

import (
	"FSM/cmp"
	"FSM/logger"
	"fmt"
)

type EntityID int

var idCounter EntityID

func init() {
	idCounter = 0
}

type Entity struct {
	Id     EntityID
	comps  map[cmp.CmpType]cmp.ICmp
	active bool
	engine *Engine
}

func NewEntity(engine *Engine) *Entity {
	rv := &Entity{
		Id:     idCounter,
		comps:  map[cmp.CmpType]cmp.ICmp{},
		engine: engine,
		active: false,
	}
	engine.AddEntity(rv)
	idCounter++
	return rv
}

func (e *Entity) Active() bool {
	return e.active
}

func (e *Entity) SetActive(s bool) {
	logger.Debug("Entity %d set active %t ", e.Id, s)
	e.active = s
}

func (e *Entity) AddComponent(c cmp.ICmp) {
	logger.Debug("Entity %d add component %s", e.Id, c.Type())
	e.comps[c.Type()] = c
	e.engine.AddComponent(e, c)

}

func (e *Entity) RemoveComponent(ct cmp.CmpType) {
	logger.Debug("Entity %d remove component %s", e.Id, ct.String())
	delete(e.comps, ct)
	e.engine.RemoveComponent(e, ct)
}

func (e *Entity) HasComponent(c cmp.CmpType) bool {
	_, ok := e.comps[c]
	return ok
}

func (e *Entity) GetComponent(c cmp.CmpType) cmp.ICmp {
	rv, ok := e.comps[c]
	if !ok {
		panic(fmt.Sprintf("Entity %d has no component %s", e.Id, c.String()))
	}
	return rv
}

func (e *Entity) GetComponents() map[cmp.CmpType]cmp.ICmp {
	return e.comps
}
