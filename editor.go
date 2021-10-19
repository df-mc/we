package worldedit

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var Editors []*Editor

func EditorByPlayer(p *player.Player) *Editor {
	for _, pl := range Editors {
		if pl.Player == p {
			return pl
		}
	}
	editor := &Editor{Player: p, worlds: []*World{}}
	Editors = append(Editors, editor)
	return editor
}

type Editor struct {
	*player.Player
	worlds []*World
}

func (e *Editor) World(wrld *world.World) *World {
	for _, w := range e.worlds {
		if w.World == wrld {
			return w
		}
	}
	cWorld := &World{World: wrld, pos: &Positions{}}
	e.worlds = append(e.worlds, cWorld)
	return cWorld
}
