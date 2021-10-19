package worldedit

import (
	"errors"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type World struct {
	*world.World
	pos *Positions
}

func (w *World) Positions() *Positions { return w.pos }

func (w *World) SetPos1(pos mgl64.Vec3) {
	w.Positions().Pos1 = &pos
}
func (w *World) SetPos2(pos mgl64.Vec3) {
	w.Positions().Pos2 = &pos
}

func (w *World) SetBlock(b world.Block) error {
	positions := w.Positions()
	if positions.Pos1 == nil || positions.Pos2 == nil {
		return errors.New("could not complete SetBlock action: pos1 or pos2 is nil")
	}
	coords := positions.BlocksCoordinatesBetween()
	for _, i := range coords {
		w.World.SetBlock(cube.PosFromVec3(i), b)
	}
	return nil
}
