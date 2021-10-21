package game

import (
	"Def/logger"
	"Def/types"
)

type EntityID int

var idCounter EntityID

func init() {
	idCounter = 0
}

type Entity struct {
	Id     EntityID
	Class  types.EntityType
	comps  map[types.CmpType]types.ICmp
	active bool
	engine *Engine
}

func NewEntity(engine *Engine, class types.EntityType) *Entity {
	rv := &Entity{
		Id:     idCounter,
		Class:  class,
		comps:  map[types.CmpType]types.ICmp{},
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

func (e *Entity) AddComponent(c types.ICmp) {
	logger.Debug("Entity %d add component %s", e.Id, c.Type())
	e.comps[c.Type()] = c
	e.engine.AddComponent(e, c)

}

func (e *Entity) RemoveComponent(ct types.CmpType) {
	logger.Debug("Entity %d remove component %s", e.Id, ct.String())
	delete(e.comps, ct)
	e.engine.RemoveComponent(e, ct)
}

func (e *Entity) HasComponent(c types.CmpType) bool {
	_, ok := e.comps[c]
	return ok
}

func (e *Entity) GetComponent(c types.CmpType) types.ICmp {
	rv, ok := e.comps[c]
	if !ok {
		return nil
	}
	return rv
}

func (e *Entity) GetComponents() map[types.CmpType]types.ICmp {
	return e.comps
}
