package commands

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	worldedit "github.com/df-mc/we"
)

var SET = cmd.New("set", "", nil, set{})

type set struct {
	Block string
}

func (s set) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if b, ok := world.BlockByName("minecraft:"+s.Block, nil); ok {
			e := worldedit.EditorByPlayer(p)
			err := e.World(p.World()).SetBlock(b)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
