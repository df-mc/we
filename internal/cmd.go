package internal

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

// Command is a base struct used throughout commands in the we repository. It implements cmd.Allower to prevent
// non-player sources from using the command.
type Command struct{}

// Allow returns false if the cmd.Source passed is not of the type *player.Player.
func (Command) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}
