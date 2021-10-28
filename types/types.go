package types

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
	Lander EntityType = iota
	Baiter EntityType = iota
	Bullet EntityType = iota
	Human  EntityType = iota
	Bomb   EntityType = iota
	Bomber EntityType = iota
	Player EntityType = iota
)

const (
	LanderMaterialise StateType = iota
	LanderSearch      StateType = iota
	LanderDrop        StateType = iota
	LanderGrab        StateType = iota
	LanderMutate      StateType = iota
	LanderDie         StateType = iota
	BaiterMaterialise StateType = iota
	BaiterHunt        StateType = iota
	HumanWalking      StateType = iota
	HumanGrabbed      StateType = iota
	HumanDropping     StateType = iota
	HumanRescued      StateType = iota
	HumanDie          StateType = iota
	PlayerPlay        StateType = iota
	PlayerDie         StateType = iota
	BomberMove        StateType = iota
	BomberDie         StateType = iota
)

func (st StateType) String() string {
	switch st {
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
	case BaiterMaterialise:
		return "Baiter-Materialise"
	case BaiterHunt:
		return "Baiter-Search"
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
	}
	return ""
}

type ICmp interface {
	Type() CmpType
}

type IEngine interface {
	GetActiveEntitiesOfClass(EntityType) []EntityID
	GetEntity(EntityID) IEntity
	MountainHeight(float64) float64
	GetCameraX() float64
	SetCameraX(float64)
	TriggerBomb(float64, float64)
	TriggerPS(float64, float64)
}

type IEntity interface {
	GetComponent(c CmpType) ICmp
	SetActive(bool)
	GetEngine() IEngine
	GetID() EntityID
	Active() bool
	Parent() EntityID
	Child() EntityID
	SetParent(EntityID)
	SetChild(EntityID)
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
