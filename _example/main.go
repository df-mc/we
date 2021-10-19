package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	worldedit "github.com/df-mc/we"
	"github.com/df-mc/we/_example/commands"
)

func main() {
	c := server.DefaultConfig()

	s := server.New(&c, nil)
	s.Start()
	s.CloseOnProgramEnd()

	cmd.Register(commands.POS1)
	cmd.Register(commands.POS2)
	cmd.Register(commands.SET)

	commands.Session = worldedit.NewSession()

	for {
		if _, err := s.Accept(); err != nil {
			return
		}
	}
}
