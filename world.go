package worldedit

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// World is a type that contains *world.World from dragonfly, it also contains a *Positions value which contains two vec3 values.
type World struct {
	*world.World
	pos *CubicPositions
}

// NewWorld returns a new *World containing the *world.World passed
func NewWorld(w *world.World) *World {
	noValueVec3 := mgl64.Vec3{}
	return &World{World: w, pos: NewCubicPositions(noValueVec3, noValueVec3)}
}

// Positions returns the positions that the player set on this *world.World.
func (w *World) Positions() *CubicPositions { return w.pos }

// SetPos1 sets the first position of the player to this world's *Positions.
func (w *World) SetPos1(pos mgl64.Vec3) {
	w.Positions().Pos1 = pos
}

// SetPos2 sets the Second position of the player to this world's *Positions.
func (w *World) SetPos2(pos mgl64.Vec3) {
	w.Positions().Pos2 = pos
}

// SetBlocks will set every block between the two vec3 values that the player have set to this world.
// It returns the amount of blocks that were placed.
func (w *World) SetBlocks(b world.Block) (n int) {
	coords := w.Positions().BlocksCoordinatesBetween()
	for _, i := range coords {
		w.World.SetBlock(cube.PosFromVec3(i), b)
		n++
	}
	return n
}
