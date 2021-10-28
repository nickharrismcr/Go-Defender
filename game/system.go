package game

import "github.com/hajimehoshi/ebiten/v2"

type SystemName int

const (
	AISystem        SystemName = iota
	DrawSystem      SystemName = iota
	PosSystem       SystemName = iota
	CollideSystem   SystemName = iota
	LifeSystem      SystemName = iota
	RadarDrawSystem SystemName = iota
	LaserDrawSystem SystemName = iota
	LaserMoveSystem SystemName = iota
)

func (t SystemName) String() string {
	switch t {
	case AISystem:
		return "AI"
	case DrawSystem:
		return "Draw"
	case PosSystem:
		return "Pos"
	case CollideSystem:
		return "Collide"
	case LifeSystem:
		return "Life"
	case RadarDrawSystem:
		return "RadarDrawSystem"
	case LaserMoveSystem:
		return "LaserMoveSystem"
	}
	return ""
}

// interface for ECS systems
type ISystem interface {
	GetName() SystemName
	Active() bool
	SetActive(bool)
	Update()
	Draw(*ebiten.Image)
	AddEntityIfRequired(e *Entity)
	RemoveEntityIfRequired(e *Entity)
}
