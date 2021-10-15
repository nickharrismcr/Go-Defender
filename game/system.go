package game

type SystemName int

const (
	AISystem SystemName = iota
)

func (t SystemName) String() string {
	switch t {
	case 0:
		return "AI"
	}
	return ""
}

type ISystem interface {
	GetName() SystemName
	Active() bool
	SetActive(bool)
	Update(dt float64)
	AddEntityIfRequired(e *Entity)
	RemoveEntityIfRequired(e *Entity)
}
