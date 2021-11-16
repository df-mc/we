package palette

import (
	"github.com/df-mc/dragonfly/server/world"
	"io"
)

// Blocks is a Palette that exists out of a slice of world.Block. It is a static palette in the sense that the
// blocks returned in the Blocks method do not change.
type Blocks struct {
	b []world.Block
}

// NewBlocks creates a Blocks palette that returns the blocks passed in the Blocks method.
func NewBlocks(b []world.Block) Blocks {
	return Blocks{b: b}
}

// Read reads a Blocks palette from an io.Reader.
func Read(r io.Reader) (Blocks, error) {
	// TODO: Implement palette reading.
	return Blocks{}, nil
}

// Write writes a Blocks palette to an io.Writer.
func (b Blocks) Write(w io.Writer) error {
	// TODO: Implement palette writing.
	return nil
}

// Blocks returns all world.Block passed to the NewBlocks function upon creation of the palette.
func (b Blocks) Blocks() []world.Block {
	return b.b
}
