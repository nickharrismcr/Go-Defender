package game

import (
	"Def/cmp"
	"Def/logger"
	"Def/types"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UPDATE int = iota
	DRAW
)

type Engine struct {
	entities              map[EntityID]*Entity
	entitiesWithComponent map[types.CmpType]map[EntityID]*Entity
	systems               map[SystemName]ISystem
	update_systems        []ISystem
	draw_systems          []ISystem
	particle_system       *ParticleSystem
	Bullets               []*Entity
}

func NewEngine() *Engine {

	return &Engine{
		entities:              make(map[EntityID]*Entity),
		entitiesWithComponent: make(map[types.CmpType]map[EntityID]*Entity),
		systems:               make(map[SystemName]ISystem),
		particle_system:       NewParticleSystem(),
		Bullets:               []*Entity{},
	}
}

func (eng *Engine) AddSystem(s ISystem, systype int) {
	eng.systems[s.GetName()] = s
	switch systype {
	case UPDATE:
		logger.Debug("Engine added update system %T ", s)
		eng.update_systems = append(eng.update_systems, s)
	case DRAW:
		logger.Debug("Engine added draw system %T ", s)
		eng.draw_systems = append(eng.draw_systems, s)
	}
}

func (eng *Engine) GetSystem(s SystemName) ISystem {
	return eng.systems[s]
}

func (eng *Engine) SetSystemActive(s SystemName, active bool) {
	eng.systems[s].SetActive(active)
}

func (eng *Engine) AddEntity(e *Entity) {
	logger.Debug("Engine added entity %d ", e.Id)
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

}

func (eng *Engine) GetEntity(id EntityID) *Entity {
	return eng.entities[id]
}

func (eng *Engine) GetEntities() map[EntityID]*Entity {
	return eng.entities
}

func (eng *Engine) AddComponent(e *Entity, c types.ICmp) {
	logger.Debug("Engine added component %s to entity %d ", c.Type(), e.Id)
	eng.addToEntitiesWithComponent(e, c)
	for _, s := range eng.systems {
		s.AddEntityIfRequired(e)
	}
	for _, s := range eng.systems {
		s.AddEntityIfRequired(e)
	}
}

func (eng *Engine) RemoveComponent(e *Entity, ct types.CmpType) {
	logger.Debug("Engine removed component %s from entity %d ", ct.String(), e.Id)
	eng.removeFromEntitiesWithComponent(e, ct)
	for _, s := range eng.systems {
		s.RemoveEntityIfRequired(e)
	}
	for _, s := range eng.systems {
		s.RemoveEntityIfRequired(e)
	}
}

func (eng *Engine) addToEntitiesWithComponent(e *Entity, c types.ICmp) {
	_, ok := eng.entitiesWithComponent[c.Type()]
	if !ok {
		eng.entitiesWithComponent[c.Type()] = map[EntityID]*Entity{}
	}
	eng.entitiesWithComponent[c.Type()][e.Id] = e
}

func (eng *Engine) removeFromEntitiesWithComponent(e *Entity, ct types.CmpType) {
	_, ok := eng.entitiesWithComponent[ct]
	if ok {
		delete(eng.entitiesWithComponent[ct], e.Id)
	}
}

func (eng *Engine) GetEntitiesWithComponent(ct types.CmpType) map[EntityID]*Entity {
	list, ok := eng.entitiesWithComponent[ct]
	if ok {
		return list
	}
	return nil
}

func (eng *Engine) Update() {
	for _, s := range eng.update_systems {
		s.Update()
	}
	eng.particle_system.Update()
}

func (eng *Engine) Draw(screen *ebiten.Image) {
	for _, s := range eng.draw_systems {
		s.Draw(screen)
	}
	eng.particle_system.Draw(screen)
}

func (eng *Engine) TriggerPS(x, y float64) {
	eng.particle_system.Trigger(x, y)
}

func (eng *Engine) TriggerBullet(x, y, dx, dy float64) {
	for _, v := range eng.Bullets {
		if !v.Active() {
			v.SetActive(true)
			pc := v.GetComponent(types.Pos).(*cmp.Pos)
			pc.X, pc.Y, pc.DX, pc.DY = x, y, 2*dx, 2*dy
			lc := v.GetComponent(types.Life).(*cmp.Life)
			lc.TicksToLive = 60
			break
		}
	}
}
