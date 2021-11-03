package game

import (
	"Def/cmp"
	"Def/event"
	"Def/gl"
	"Def/logger"
	"Def/types"
	"Def/util"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UPDATE int = iota
	DRAW
)

type Engine struct {
	entities              map[types.EntityID]*Entity
	entitiesWithComponent map[types.CmpType]map[types.EntityID]types.IEntity
	systems               map[SystemName]ISystem
	updateSystems         []ISystem
	drawSystems           []ISystem
	particleSystem        *ParticleSystem
	world                 *World
	stars                 *Stars
	chars                 *Characters
	BulletPool            []*Entity
	BombPool              []*Entity
	LaserPool             []*Entity
	LaserColIdx           int
	flash                 int
}

func NewEngine() *Engine {

	e := &Engine{
		entities:              make(map[types.EntityID]*Entity),
		entitiesWithComponent: make(map[types.CmpType]map[types.EntityID]types.IEntity),
		systems:               make(map[SystemName]ISystem),
		chars:                 nil,
		BulletPool:            []*Entity{},
		BombPool:              []*Entity{},
		LaserPool:             []*Entity{},
		LaserColIdx:           0,
	}
	e.particleSystem = NewParticleSystem(e)
	e.stars = NewStars(e)
	e.world = NewWorld(e)
	return e
}

func (eng *Engine) AddSystem(s ISystem, systype int) {
	eng.systems[s.GetName()] = s
	switch systype {
	case UPDATE:
		logger.Debug("Engine added update system %T ", s)
		eng.updateSystems = append(eng.updateSystems, s)
	case DRAW:
		logger.Debug("Engine added draw system %T ", s)
		eng.drawSystems = append(eng.drawSystems, s)
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
		for _, s := range eng.updateSystems {
			s.AddEntityIfRequired(e)
		}
		for _, s := range eng.drawSystems {
			s.AddEntityIfRequired(e)
		}
	}
}

func (eng *Engine) GetEntity(id types.EntityID) types.IEntity {
	return eng.entities[id]
}

func (eng *Engine) GetActiveEntitiesOfClass(et types.EntityType) []types.EntityID {

	rv := []types.EntityID{}
	for _, v := range eng.entities {
		if v.Class == et && v.Active() {
			rv = append(rv, v.Id)
		}
	}
	return rv
}

func (eng *Engine) GetEntities() map[types.EntityID]*Entity {
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
		eng.entitiesWithComponent[c.Type()] = map[types.EntityID]types.IEntity{}
	}
	eng.entitiesWithComponent[c.Type()][e.Id] = e
}

func (eng *Engine) removeFromEntitiesWithComponent(e *Entity, ct types.CmpType) {
	_, ok := eng.entitiesWithComponent[ct]
	if ok {
		delete(eng.entitiesWithComponent[ct], e.Id)
	}
}

func (eng *Engine) GetEntitiesWithComponent(ct types.CmpType) map[types.EntityID]types.IEntity {
	list, ok := eng.entitiesWithComponent[ct]
	if ok {
		return list
	}
	return nil
}

func (eng *Engine) Update() {
	for _, s := range eng.updateSystems {
		s.Update()
	}
	eng.particleSystem.Update()
	eng.stars.Update()
	eng.chars.Update()
	event.UpdateQueue()
	eng.world.Update()

}

func (eng *Engine) Draw(screen *ebiten.Image) {

	if eng.flash > 0 {
		if eng.flash%2 == 0 {
			col := gl.Cols[(eng.flash/2)%5].Convert()
			screen.Fill(col)
		}
		eng.flash--
	}

	for _, s := range eng.drawSystems {
		s.Draw(screen)
	}
	eng.particleSystem.Draw(screen)
	eng.world.Draw(screen)
	eng.stars.Draw(screen)
	eng.chars.Draw(screen)
}

func (eng *Engine) TriggerPS(x, y float64) {
	eng.particleSystem.Trigger(x, y)
}

func (eng *Engine) TriggerBullet(x, y, dx, dy float64) {
	if dx > 5 {
		return
	}
	for _, v := range eng.BulletPool {
		if !v.Active() {
			v.SetActive(true)
			pc := v.GetComponent(types.Pos).(*cmp.Pos)
			pc.X, pc.Y, pc.DX, pc.DY = x, y, dx, dy
			lc := v.GetComponent(types.Life).(*cmp.Life)
			lc.TicksToLive = 120
			break
		}
	}
}

func (eng *Engine) TriggerBomb(x, y float64) {
	for _, v := range eng.BombPool {
		if !v.Active() {
			v.SetActive(true)
			pc := v.GetComponent(types.Pos).(*cmp.Pos)
			pc.X, pc.Y, pc.DX, pc.DY = x, y, 0, 0
			lc := v.GetComponent(types.Life).(*cmp.Life)
			lc.TicksToLive = 320
			break
		}
	}
}

func (eng *Engine) TriggerLaser(x, y, dx float64) {
	for _, v := range eng.LaserPool {
		if !v.Active() {
			v.SetActive(true)
			pc := v.GetComponent(types.Pos).(*cmp.Pos)
			pc.X, pc.Y, pc.DX, pc.DY = x, y, dx, 0
			lc := v.GetComponent(types.Life).(*cmp.Life)
			lc.TicksToLive = 90
			dc := v.GetComponent(types.LaserDraw).(*cmp.LaserDraw)
			dc.Color = gl.LaserCols[eng.LaserColIdx%15]
			mv := v.GetComponent(types.LaserMove).(*cmp.LaserMove)
			mv.Length = 0
			eng.LaserColIdx++
			break
		}
	}

}

func (eng *Engine) GetPlayer() types.IEntity {
	return eng.entities[gl.PlayerID]
}

func (eng *Engine) MountainHeight(wx float64) float64 {
	return eng.world.At(wx)
}

func (eng *Engine) AddString(s string, x, y float64) int {
	if eng.chars == nil {
		eng.chars = NewCharacters()
	}
	return eng.chars.Add(s, x, y)
}

func (eng *Engine) ChangeString(id int, s string) {
	eng.chars.Change(id, s)
}

func (eng *Engine) ClearChars() {
	eng.chars.Clear()
}

func (eng *Engine) Kill(e types.IEntity) {

	if !e.HasComponent(types.AI) {
		e.SetActive(false)
		return
	}
	ai := e.GetComponent(types.AI).(*cmp.AI)
	switch e.GetClass() {
	case types.Lander:
		ai.NextState = types.LanderDie
	case types.Human:
		ai.NextState = types.HumanDie
	case types.Bomber:
		ai.NextState = types.BomberDie
	case types.Pod:
		ai.NextState = types.PodDie
	case types.Swarmer:
		ai.NextState = types.SwarmerDie
	case types.Baiter:
		ai.NextState = types.BaiterDie
	default:
		e.SetActive(false)
	}
}

func (eng *Engine) SetFlash(c int) {
	eng.flash = 2 * c
}

func (eng *Engine) SmartBomb() {

	for id := range eng.entitiesWithComponent[types.Shootable] {
		e := eng.entities[id]
		if e.Active() && e.Class != types.Human {
			pc := e.GetComponent(types.Pos).(*cmp.Pos)
			if !util.OffScreen(util.ScreenX(pc.X), pc.Y) {
				eng.Kill(e)
			}
		}
	}

}

func (eng *Engine) ExplodeWorld() {
	eng.world.Explode()
}

func (eng *Engine) MutateAll() {

	landers := eng.GetActiveEntitiesOfClass(types.Lander)

	for _, id := range landers {
		e := eng.GetEntity(id)
		ai := e.GetComponent(types.AI).(*cmp.AI)
		ai.NextState = types.LanderMutate
	}

}
