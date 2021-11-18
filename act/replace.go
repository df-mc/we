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

// Replace is an action that replaces set blocks in the selection with other blocks.
type Replace struct {
	b, old []world.Block
}

// At always returns a random block set in the action if the block at the given x, y and z is in the
// old slice.
func (r Replace) At(x int, y int, z int, ra *rand.Rand, _ *world.World, at func(x, y, z int) world.Block) (world.Block, world.Liquid) {
	old := at(x, y, z)
	for _, s := range r.old {
		if s == old {
			return r.b[ra.Intn(len(r.b))], nil
		}
	}
	return nil, nil
}

// Form ...
func (r Replace) Form(shape brush.Shape) form.Form {
	return form.New(replaceForm{s: shape,
		BlockPalette:    form.NewInput(msg.BlockPalette, "M", "M"),
		ReplacedPalette: form.NewInput(msg.ReplacedPalette, "M", "M"),
	}, msg.ReplaceMenu)
}

// replaceForm is the form for the Replace action.
type replaceForm struct {
	BlockPalette    form.Input
	ReplacedPalette form.Input
	s               brush.Shape
}

// Submit ...
func (s replaceForm) Submit(submitter form.Submitter) {
	p := submitter.(*player.Player)
	ph, _ := palette.LookupHandler(p)
	pal, ok := ph.Palette(s.BlockPalette.Value())
	if !ok || len(pal.Blocks()) == 0 {
		p.Message(text.Colourf("<red>%v</red>", msg.InvalidPalette))
		return
	}
	rPal, ok := ph.Palette(s.ReplacedPalette.Value())
	if !ok || len(rPal.Blocks()) == 0 {
		p.Message(text.Colourf("<red>%v</red>", msg.InvalidPalette))
		return
	}
	held, otherHeld := p.HeldItems()
	p.SetHeldItems(brush.New(s.s, Replace{b: pal.Blocks(), old: rPal.Blocks()}).Bind(held), otherHeld)
}
