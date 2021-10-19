package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

var POS1 = cmd.New("pos1", "", nil, pos1{})
var POS2 = cmd.New("pos2", "", nil, pos2{})

type pos1 struct{}

func (pos1) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		Session.EditorByPlayer(p).World(p.World()).SetPos1(p.Position())
	}
}

type pos2 struct{}

func (pos2) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		Session.EditorByPlayer(p).World(p.World()).SetPos2(p.Position())
	}
}
