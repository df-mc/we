package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var SET = cmd.New("set", "", nil, set{})

type set struct {
	Block string
}

func (s set) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		// This is only working with some blocks since there is no properties argument right now
		// You can test with the 'cobblestone' block
		if b, ok := world.BlockByName("minecraft:" + s.Block, nil); ok {

			e := Session.EditorByPlayer(p)
			n := e.World(p.World()).SetBlocks(b)
			p.Messagef("%v blocks were placed", n)
		}
	}
}
