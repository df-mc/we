package palette

import (
	"github.com/df-mc/dragonfly/server/world"
)

// Palette is a source for a selection of world.Block to be used in a world edit action.
type Palette interface {
	// Blocks returns the list of world.Block that should be used as palette for a world edit action. Blocks can
	// return the same world.Block multiple times to change the occurrence of one block vs another block.
	Blocks() []world.Block
}
