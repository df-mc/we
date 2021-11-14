package act

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/we/brush"
	"github.com/df-mc/we/internal/msg"
	"github.com/df-mc/we/palette"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"math/rand"
)

// Fill is an action which fills the entire selection with one or more blocks.
type Fill struct {
	b []world.Block
}

// At always returns a random block set in the action.
func (f Fill) At(_ int, _ int, _ int, r *rand.Rand, _ *world.World, _ func(x, y, z int) world.Block) (world.Block, world.Liquid) {
	return f.b[r.Intn(len(f.b))], nil
}

// Form ...
func (f Fill) Form(shape brush.Shape) form.Form {
	return form.New(fillForm{s: shape, Palette: form.NewInput(msg.BlockPalette, "M", "M")}, msg.FillMenu)
}

// fillForm is the form for the Fill action.
type fillForm struct {
	Palette form.Input
	s       brush.Shape
}

// Submit ...
func (s fillForm) Submit(submitter form.Submitter) {
	p := submitter.(*player.Player)
	ph, _ := palette.LookupHandler(p)
	pal, ok := ph.Palette(s.Palette.Value())
	if !ok || len(pal.Blocks()) == 0 {
		p.Message(text.Colourf("<red>%v</red>", msg.InvalidPalette))
		return
	}
	held, otherHeld := p.HeldItems()
	p.SetHeldItems(brush.New(s.s, Fill{b: pal.Blocks()}).Bind(held), otherHeld)
}
