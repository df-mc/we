package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/we/example/commands"
)

func main() {
	c := server.DefaultConfig()

	s := server.New(&c, nil)
	s.Start()
	s.CloseOnProgramEnd()

	cmd.Register(commands.POS1)
	cmd.Register(commands.POS2)
	cmd.Register(commands.SET)

	for {
		if _, err := s.Accept();err!=nil{
			return
		}
	}
}
