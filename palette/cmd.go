package palette

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/we/internal/msg"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// SetCommand implements the selection of a Selection palette in the world that a player is in. This palette may
// later be saved using SaveCommand.
type SetCommand struct {
	command
	Sub set
}

// Run enables palette selection for the *player.Player that runs the command.
func (c SetCommand) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	h, _ := LookupHandler(p)
	h.selecting = 2
	o.Printf(text.Colourf("<green>%v</green>", msg.StartPaletteSelection))
}

// SaveCommand implements the saving of palettes to disk, so that they may be re-used.
type SaveCommand struct {
	command
	Sub save
	// Name is the name by which the palette currently selected should be saved. The palette will be saved to a
	// database so that it can be reloaded when the server restarts.
	Name string `name:"name"`
}

// Run allows a *player.Player to save the Selection previously created using /palette to disk with a specific name,
// so that it can be re-used.
func (s SaveCommand) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	h, _ := LookupHandler(p)
	if _, ok := h.Palette(s.Name); ok {
		// Don't let players create palettes with names that already exist. We don't want to silently overwrite them.
		o.Errorf(msg.PaletteExists, s.Name, s.Name)
		return
	}
	if h.m.Zero() {
		// Players must first select a palette using /palette.
		o.Errorf(msg.NoPaletteSelected)
		return
	}
	h.palettes.Store(s.Name, NewBlocks(h.m.Blocks()))
	o.Printf(text.Colourf("<green>%v</green>", msg.PaletteSaved), h.m.Min, h.m.Max, s.Name)
}

// DeleteCommand implements the deletion of palettes previously saved using SaveCommand.
type DeleteCommand struct {
	command
	Sub del
	// Name is the name of the palette to delete. Upon deleting, the palette will be removed from the database
	// it is stored in.
	Name paletteName `name:"name"`
}

// Run allows a *player.Player to delete a palette previously saved using /palette save.
func (d DeleteCommand) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)
	name := string(d.Name)

	h, _ := LookupHandler(p)
	if _, ok := h.palettes.Load(name); !ok {
		// Palette didn't exist, no point logging this as if it deleted properly, that only masks bugs.
		o.Errorf(msg.PaletteDoesNotExist, name)
		return
	}
	h.palettes.Delete(name)
	o.Printf(text.Colourf("<green>%v</green>", msg.PaletteDeleted), name)
}

type (
	set         string
	save        string
	del         string
	paletteName string
)

func (set) SubName() string  { return "set" }
func (save) SubName() string { return "save" }
func (del) SubName() string  { return "delete" }

func (p paletteName) Type() string { return "PaletteName" }
func (p paletteName) Options(src cmd.Source) []string {
	h, ok := LookupHandler(src.(*player.Player))
	if !ok {
		return nil
	}

	var m []string
	h.palettes.Range(func(key, value interface{}) bool {
		m = append(m, key.(string))
		return true
	})
	return m
}

// command is a base struct used throughout commands in the we repository. It implements cmd.Allower to prevent
// non-player sources from using the command.
type command struct{}

// Allow returns false if the cmd.Source passed is not of the type *player.Player.
func (command) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}
