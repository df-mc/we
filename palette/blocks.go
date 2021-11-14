package palette

import (
	"github.com/df-mc/dragonfly/server/world"
	"io"
)

type Blocks struct {
	b []world.Block
}

func NewBlocks(b []world.Block) Blocks {
	return Blocks{b: b}
}

func Read(r io.Reader) (Blocks, error) {
	// TODO: Implement palette reading.
	return Blocks{}, nil
}

func (b Blocks) Write(w io.Writer) error {
	// TODO: Implement palette writing.
	return nil
}

func (b Blocks) Blocks() []world.Block {
	return b.b
}
