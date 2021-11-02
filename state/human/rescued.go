package human

import (
	"Def/cmp"
	"Def/event"
	"Def/gl"
	"Def/types"
)

// NB States should not contain entity state ;) they should act on passed components

type HumanRescued struct {
	Name types.StateType
}

func NewHumanRescued() *HumanRescued {
	return &HumanRescued{
		Name: types.HumanRescued,
	}
}

func (s *HumanRescued) GetName() types.StateType {
	return s.Name
}

func (s *HumanRescued) Enter(ai *cmp.AI, e types.IEntity) {

	ev := event.NewHumanRescued(e)
	event.NotifyEvent(ev)
	e.SetParent(e.GetEngine().GetPlayer().GetID())
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
}

func (s *HumanRescued) Update(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pe := e.GetEngine().GetEntity(e.Parent())

	pec := pe.GetComponent(types.Pos).(*cmp.Pos)
	pc.Y = pec.Y + 50
	pc.X = pec.X
	// TODO why not aligned ?
	if pc.Y > gl.ScreenHeight-e.GetEngine().MountainHeight(pc.X) {
		ai.NextState = types.HumanWalking
		ev := event.NewHumanSaved(e)
		event.NotifyEvent(ev)
	}
}
