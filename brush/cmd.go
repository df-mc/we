package brush

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/we/internal/msg"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// BindCommand implements the binding of a Brush to an item in the player's inventory.
type BindCommand struct {
	Sub bind
}

func (BindCommand) Allow(src cmd.Source) bool { return pl(src) }

// Run implements the binding of a Brush to the held item by sending a brush selection form.
func (BindCommand) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	held, _ := p.HeldItems()
	if held.Empty() {
		o.Errorf(msg.BindNeedsItem)
		return
	}
	if _, ok := find(held); ok {
		o.Errorf(msg.AlreadyBound)
		return
	}
	p.SendForm(NewSelectionForm())
}

// UnbindCommand implements unbinding of a Brush previously attached to an item in the player's inventory using
// /brush bind.
type UnbindCommand struct {
	Sub unbind
}

func (UnbindCommand) Allow(src cmd.Source) bool { return pl(src) }

// Run implements the unbinding of a Brush bound to the item held.
func (c UnbindCommand) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	held, other := p.HeldItems()
	if _, ok := find(held); !ok {
		o.Errorf(msg.NotBound)
		return
	}
	p.SetHeldItems(Unbind(held), other)
	o.Printf(text.Colourf("<green>%v</green>", msg.BrushUnbound))
}

// UndoCommand implements the undoing of one of the most recent actions performed by a player using a Brush.
type UndoCommand struct {
	Sub undo
}

func (UndoCommand) Allow(src cmd.Source) bool { return pl(src) }

// Run implements the undoing of an action performed with a Brush.
func (c UndoCommand) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)
	h, _ := LookupHandler(p)
	if !h.UndoLatest() {
		o.Errorf(msg.NoUndo)
		return
	}
	o.Printf(text.Colourf("<green>%v</green>", msg.UndoSuccessful, len(h.undo)))
}

// pl checks if the cmd.Source passed is a *player.Player and returns true if so.
func pl(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}

type bind string

func (bind) SubName() string { return "bind" }

type unbind string

func (unbind) SubName() string { return "unbind" }

type undo string

func (undo) SubName() string { return "undo" }
