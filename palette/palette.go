package palette

import (
	"github.com/df-mc/dragonfly/server/world"
)

type Palette interface {
	Blocks() []world.Block
}
