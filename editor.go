package worldedit

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

// Editor is a type in which contains a player and a list of *World
type Editor struct {
	*player.Player
	worlds []*World
}

// NewEditor returns a new editor with the player passed
func NewEditor(p *player.Player) *Editor {
	return &Editor{
		Player: p,
		worlds: make([]*World, 0),
	}
}

// AddWorld will add a new world to the editor world list creating a new *World with the *world.World passed
func (e *Editor) AddWorld(w *world.World) { e.worlds = append(e.worlds, NewWorld(w)) }

// World returns a *World by looking in the editor world list for the *world.World passed
func (e *Editor) World(world *world.World) *World {
	for _, w := range e.worlds {
		if w.World == world {
			return w
		}
	}
	world2 := &World{World: world, pos: &Positions{}}
	e.worlds = append(e.worlds, world2)
	return world2
}
