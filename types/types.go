package types

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateType int
type EntityType int
type EntityID int
type CmpType int
type ActionType int

const (
	Up         ActionType = iota
	Down       ActionType = iota
	Thrust     ActionType = iota
	Fire       ActionType = iota
	Reverse    ActionType = iota
	SmartBomb  ActionType = iota
	HyperSpace ActionType = iota
)

const (
	Lander     EntityType = iota
	Baiter     EntityType = iota
	Bullet     EntityType = iota
	Human      EntityType = iota
	Bomb       EntityType = iota
	Bomber     EntityType = iota
	Player     EntityType = iota
	Laser      EntityType = iota
	Pod        EntityType = iota
	Swarmer    EntityType = iota
	Score      EntityType = iota
	Flame      EntityType = iota
	Game       EntityType = iota
	PlayerLife EntityType = iota
	HUDBomb    EntityType = iota
)

func (et EntityType) String() string {
	switch et {
	case Lander:
		return "Lander"
	case Baiter:
		return "Baiter"
	case Bullet:
		return "Bullet"
	case Human:
		return "Human"
	case Bomb:
		return "Bomb"
	case Bomber:
		return "Bomber"
	case Player:
		return "Player"
	case Laser:
		return "Laser"
	case Pod:
		return "Pod"
	case Swarmer:
		return "Swarmer"
	case Score:
		return "Score"
	case Flame:
		return "Flame"
	}
	return ""
}

const (
	LanderWait        StateType = iota
	LanderMaterialise StateType = iota
	LanderSearch      StateType = iota
	LanderDrop        StateType = iota
	LanderGrab        StateType = iota
	LanderMutate      StateType = iota
	LanderDie         StateType = iota
	BaiterMaterialise StateType = iota
	BaiterWait        StateType = iota
	BaiterHunt        StateType = iota
	BaiterDie         StateType = iota
	HumanWalking      StateType = iota
	HumanGrabbed      StateType = iota
	HumanDropping     StateType = iota
	HumanRescued      StateType = iota
	HumanDie          StateType = iota
	PlayerPlay        StateType = iota
	PlayerDie         StateType = iota
	BomberMove        StateType = iota
	BomberDie         StateType = iota
	PodMove           StateType = iota
	PodDie            StateType = iota
	SwarmerMove       StateType = iota
	SwarmerDie        StateType = iota
	GameIntro         StateType = iota
	GameStart         StateType = iota
	GamePlay          StateType = iota
	GameLevelEnd      StateType = iota
	GameOver          StateType = iota
)

func (st StateType) String() string {
	switch st {
	case LanderWait:
		return "Lander-Wait"
	case LanderMaterialise:
		return "Lander-Materialise"
	case LanderSearch:
		return "Lander-Search"
	case LanderDrop:
		return "Lander-Drop"
	case LanderGrab:
		return "Lander-Grab"
	case LanderMutate:
		return "Lander-Mutate"
	case LanderDie:
		return "Lander-Die"
	case BaiterWait:
		return "Baiter-Wait"
	case BaiterMaterialise:
		return "Baiter-Materialise"
	case BaiterHunt:
		return "Baiter-Search"
	case BaiterDie:
		return "Baiter-Die"
	case HumanWalking:
		return "Human-Walking"
	case HumanGrabbed:
		return "Human-Grabbed"
	case HumanDropping:
		return "Human-Dropping"
	case HumanRescued:
		return "Human-Rescued"
	case HumanDie:
		return "Human-Die"
	case PlayerPlay:
		return "Player Play	"
	case PlayerDie:
		return "Player-Die"
	case BomberMove:
		return "Bomber Move"
	case BomberDie:
		return "Bomber-Die"
	case PodMove:
		return "Pod-Move"
	case PodDie:
		return "Pod-Die"
	case SwarmerMove:
		return "Swarmer-Move"
	case SwarmerDie:
		return "Swarmer-Die"
	case GameIntro:
		return "Game-Intro"
	case GameStart:
		return "Game-Start"
	case GamePlay:
		return "Game-Play"
	case GameLevelEnd:
		return "Game-LevelEnd"
	case GameOver:
		return "Game-Over"
	}
	return ""
}

const (
	AI        CmpType = iota
	Pos       CmpType = iota
	Draw      CmpType = iota
	Collide   CmpType = iota
	Life      CmpType = iota
	RadarDraw CmpType = iota
	Ship      CmpType = iota
	Shootable CmpType = iota
	LaserDraw CmpType = iota
	LaserMove CmpType = iota
)

func (t CmpType) String() string {
	switch t {
	case AI:
		return "AI"
	case Pos:
		return "Pos"
	case Draw:
		return "Draw"
	case Collide:
		return "Collide"
	case Life:
		return "Life"
	case RadarDraw:
		return "RadarDraw"
	case Ship:
		return "Ship"
	case Shootable:
		return "Shootable"
	case LaserDraw:
		return "LaserDraw"
	case LaserMove:
		return "Laser"
	}
	return ""
}

type SystemName int

const (
	AISystem        SystemName = iota
	DrawSystem      SystemName = iota
	AnimateSystem   SystemName = iota
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
	case AnimateSystem:
		return "Animate"
	case PosSystem:
		return "Pos"
	case CollideSystem:
		return "Collide"
	case LifeSystem:
		return "Life"
	case RadarDrawSystem:
		return "RadarDraw"
	case LaserMoveSystem:
		return "LaserMove"
	}
	return ""
}

type ICmp interface {
	Type() CmpType
}

type IEngine interface {
	InitEnemies()
	InitHumans()
	GetActiveEntitiesOfClass(EntityType) []EntityID
	GetEntity(EntityID) IEntity
	MountainHeight(float64) float64
	TriggerBomb(float64, float64)
	TriggerPS(float64, float64)
	Kill(IEntity)
	GetEntities() map[EntityID]IEntity
	GetEntitiesWithComponent(CmpType) map[EntityID]IEntity
	GetPlayer() IEntity
	SetFlash(int)
	GetSystem(SystemName) ISystem
	LevelStart()
	LevelEnd()
	GameOver()
	Terminate()
}

// interface for ECS systems
type ISystem interface {
	GetName() SystemName
	Active() bool
	SetActive(bool)
	Update()
	Draw(*ebiten.Image)
	AddEntityIfRequired(IEntity)
	RemoveEntityIfRequired(IEntity)
	RemoveEntity(IEntity)
}

type IEntity interface {
	GetComponent(c CmpType) ICmp
	GetComponents() map[CmpType]ICmp
	RemoveComponent(c CmpType)
	HasComponent(c CmpType) bool
	SetActive(bool)
	GetEngine() IEngine
	GetID() EntityID
	Active() bool
	Parent() EntityID
	Child() EntityID
	SetParent(EntityID)
	SetChild(EntityID)
	GetClass() EntityType
	AddComponent(ICmp)
	Paused() bool
	SetPaused(bool)
}

type ColorF struct {
	R, G, B, A float64
}

func (c ColorF) Add(oth ColorF) ColorF {
	return ColorF{
		R: c.R + oth.R,
		G: c.G + oth.G,
		B: c.B + oth.B,
		A: c.A + oth.A,
	}

}

func (c ColorF) Convert() color.Color {

	return color.RGBA{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
		A: uint8(c.A * 255),
	}
}

func (c ColorF) Subtract(oth ColorF) ColorF {
	return ColorF{
		R: c.R - oth.R,
		G: c.G - oth.G,
		B: c.B - oth.B,
		A: c.A - oth.A,
	}
}

func (c ColorF) Multiply(f float64) ColorF {
	return ColorF{
		R: c.R * f,
		B: c.B * f,
		A: c.A * f,
		G: c.G * f,
	}
}
