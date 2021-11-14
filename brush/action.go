package brush

import (
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"math/rand"
)

type Action interface {
	At(x, y, z int, r *rand.Rand, w *world.World, at func(x, y, z int) world.Block) (world.Block, world.Liquid)
	// Form returns a form that has to be submitted by a player in order to provide additional values for the
	// action. Actions that do not need additional data can return nil for this value. The geo.Shape selected by
	// the *player.Player is passed to the function.
	Form(s Shape) form.Form
}
