package brush

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"math/rand"
	"time"
)

// Perform performs the world edit action passed in a specific shape, in the world that is passed. Perform
// will only ever edit blocks found within the shape passed.
// Perform returns a function which may be called to revert the modification.
func Perform(pos cube.Pos, s Shape, a Action, w *world.World) (revert func()) {
	d := s.Dim()
	// The shapes measure according to a centre position, so the base of our structure is offset.
	base := pos.Add(cube.Pos{-d[0] / 2, -d[1] / 2, -d[2] / 2})
	st := &structure{base: base, s: s, a: a, d: d, w: w, cx: pos[0], cy: pos[1], cz: pos[2], m: make(map[cube.Pos]world.Block)}
	w.BuildStructure(base, st)
	return st.Revert
}

// structure is a wrapper around a Shape and an Action, which together form an operation on the world.
type structure struct {
	base       cube.Pos
	s          Shape
	a          Action
	d          [3]int
	cx, cy, cz int
	w          *world.World
	m          map[cube.Pos]world.Block
	bAt        func(x, y, z int) world.Block
}

// Dimensions returns the dimensions of the shape.
func (s *structure) Dimensions() [3]int {
	return s.d
}

// At checks if the x, y and z passed are within if the shape that the structure holds. If this is the case,
// At returns the block returned by the type held. If not, nil is returned.
func (s *structure) At(x, y, z int, at func(x, y, z int) world.Block) (world.Block, world.Liquid) {
	s.bAt = at
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if s.s.Inside(s.cx, s.cy, s.cz, x+s.cx-s.d[0]/2, y+s.cy-s.d[1]/2, z+s.cz-s.d[2]/2) {
		if v, liq := s.a.At(x, y, z, r, s.w, s.blockAt); v != nil {
			s.m[cube.Pos{x, y, z}] = at(x, y, z)
			s.bAt = nil
			return v, liq
		}
	}
	s.bAt = nil
	return nil, nil
}

// blockAt returns the block at the position passed in the world of the structure, or a block that was there
// before the structure modified it.
func (s *structure) blockAt(x, y, z int) world.Block {
	if v, ok := s.m[cube.Pos{x, y, z}]; ok {
		return v
	}
	return s.bAt(x, y, z)
}

// Revert reverts the placement of the structure, placing back all blocks that were changed by the initial
// placement.
func (s *structure) Revert() {
	s.w.BuildStructure(s.base, &structureRevert{base: s.base, d: s.d, m: s.m})
}

// structureRevert represents a structure that handles the reverting of a normal structure.
type structureRevert struct {
	base cube.Pos
	d    [3]int
	m    map[cube.Pos]world.Block
}

// Dimensions ...
func (s *structureRevert) Dimensions() [3]int {
	return s.d
}

// At ...
func (s *structureRevert) At(x, y, z int, _ func(x, y, z int) world.Block) (world.Block, world.Liquid) {
	b, _ := s.m[cube.Pos{x + s.base[0], y + s.base[1], z + s.base[2]}]
	return b, nil
}
