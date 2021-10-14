package game

import "FSM/cmp"

const (
	UPDATE int = iota
	DRAW
)

type Engine struct {
	entities              map[EntityID]*Entity
	entitiesWithComponent map[cmp.CmpType]map[EntityID]*Entity
	update_systems        []ISystem
	draw_systems          []ISystem
}

func NewEngine() *Engine {
	return &Engine{
		entities:              make(map[EntityID]*Entity),
		entitiesWithComponent: make(map[cmp.CmpType]map[EntityID]*Entity),
	}
}

func (eng *Engine) AddSystem(s ISystem, systype int) {
	switch systype {
	case UPDATE:
		eng.update_systems = append(eng.update_systems, s)
	case DRAW:
		eng.draw_systems = append(eng.draw_systems, s)
	}
}

func (eng *Engine) AddEntity(e *Entity) {
	eng.entities[e.Id] = e
	for _, c := range e.GetComponents() {
		eng.addToEntitiesWithComponent(e, c)
		for _, s := range eng.update_systems {
			s.AddEntityIfRequired(e)
		}
		for _, s := range eng.draw_systems {
			s.AddEntityIfRequired(e)
		}
	}
	e.Activate()
}

func (eng *Engine) GetEntities() map[EntityID]*Entity {
	return eng.entities
}

func (eng *Engine) AddComponent(e *Entity, c cmp.ICmp) {
	eng.addToEntitiesWithComponent(e, c)
	for _, s := range eng.update_systems {
		s.AddEntityIfRequired(e)
	}
	for _, s := range eng.draw_systems {
		s.AddEntityIfRequired(e)
	}
}

func (eng *Engine) RemoveComponent(e *Entity, ct cmp.CmpType) {
	eng.removeFromEntitiesWithComponent(e, ct)
	for _, s := range eng.update_systems {
		s.RemoveEntityIfRequired(e)
	}
	for _, s := range eng.draw_systems {
		s.RemoveEntityIfRequired(e)
	}
}

func (eng *Engine) addToEntitiesWithComponent(e *Entity, c cmp.ICmp) {
	_, ok := eng.entitiesWithComponent[c.Type()]
	if !ok {
		eng.entitiesWithComponent[c.Type()] = map[EntityID]*Entity{}
	}
	eng.entitiesWithComponent[c.Type()][e.Id] = e
}

func (eng *Engine) removeFromEntitiesWithComponent(e *Entity, ct cmp.CmpType) {
	_, ok := eng.entitiesWithComponent[ct]
	if ok {
		delete(eng.entitiesWithComponent[ct], e.Id)
	}
}

func (eng *Engine) GetEntitiesWithComponent(ct cmp.CmpType) map[EntityID]*Entity {
	list, ok := eng.entitiesWithComponent[ct]
	if ok {
		return list
	}
	return nil
}

func (eng *Engine) Update(dt float64) {
	for _, s := range eng.update_systems {
		s.Update(dt)
	}
	for _, s := range eng.draw_systems {
		s.Update(dt)
	}
}
