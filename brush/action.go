package brush

import (
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"math/rand"
)

// Action is a brush action that may be performed on all blocks in a Shape through a call to Perform.
type Action interface {
	// At returns the world.Block and world.Liquid behind it that should be placed at a specific x, y and z in the
	// *world.World passed.
	// At should use the *rand.Rand instance passed to produce random numbers and must only use the at function to
	// read blocks at a specific position in the world.
	// If At returns a nil world.Block, no block will be placed at that position.
	At(x, y, z int, r *rand.Rand, w *world.World, at func(x, y, z int) world.Block) (world.Block, world.Liquid)
	// Form returns a form that has to be submitted by a player in order to provide additional values for the
	// action. Actions that do not need additional data can return nil for this value. The geo.Shape selected by
	// the *player.Player is passed to the function.
	Form(s Shape) form.Form
}
