package palette

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/we/geo"
)

// Selection is a Palette implementation based on a selection done by a player in a world. It is directly tied to
// that world and cannot return the blocks without it.
type Selection struct {
	geo.Area
	w *world.World
}

// NewSelection creates a new Selection based on the two corners and a world.
func NewSelection(a, b cube.Pos, w *world.World) Selection {
	return Selection{w: w, Area: geo.NewArea(a[0], a[1], a[2], b[0], b[1], b[2])}
}

// Zero checks if the Selection was set and created using NewSelection. It also checks if the we.Area held by the
// Selection is non-zero.
func (p Selection) Zero() bool {
	var zero geo.Area
	return p.w == nil || p.Area == zero
}

// Blocks returns all world.Block present between two corners in the world passed to NewSelection upon creation.
func (p Selection) Blocks() []world.Block {
	if p.Zero() {
		return nil
	}
	m := make([]world.Block, 0, p.Dx()*p.Dy()*p.Dz())
	p.Range(func(x, y, z int) {
		m = append(m, p.w.Block(cube.Pos{x, y, z}))
	})
	return m
}
