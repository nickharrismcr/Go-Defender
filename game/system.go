package game

type ISystem interface {
	Active() bool
	SetActive(bool)
	Update(dt float64)
	AddEntityIfRequired(e *Entity)
	RemoveEntityIfRequired(e *Entity)
}
