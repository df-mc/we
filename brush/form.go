package brush

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/we/internal/msg"
)

// NewSelectionForm returns a new SelectionForm form.
func NewSelectionForm() form.Form {
	return form.New(SelectionForm{
		Shape:  form.NewDropdown(msg.BrushShape, shapeNames, 0),
		Radius: form.NewSlider(msg.BrushRadius, 0, 100, 1, 5),
		Action: form.NewDropdown(msg.BrushAction, actionNames, 0),
	}, msg.BrushSelection)
}

// SelectionForm is a form that is sent when the player binds a brush to an item. It will allow the user to select
// a brush shape and action, and will follow up with another form with specific buttons to alter the behaviour
// of the action.
type SelectionForm struct {
	Shape  form.Dropdown
	Radius form.Slider
	Action form.Dropdown
}

// Submit ...
func (s SelectionForm) Submit(submitter form.Submitter) {
	p := submitter.(*player.Player)

	shape, action := shapeByName(shapeNames[s.Shape.Value()], int(s.Radius.Value())), actionByName(actionNames[s.Action.Value()])
	if f := action.Form(shape); f != nil {
		submitter.SendForm(f)
		return
	}
	// Brush is already finalised so we can bind it right away.
	hand, offHand := p.HeldItems()
	p.SetHeldItems(New(shape, action).Bind(hand), offHand)
}
