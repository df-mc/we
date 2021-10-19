package worldedit

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Positions struct {
	Pos1 *mgl64.Vec3
	Pos2 *mgl64.Vec3
}

func (p *Positions) MaxXYZ() (float64, float64, float64) {
	return math.Max(p.Pos1.X(), p.Pos2.X()), math.Max(p.Pos1.Y(), p.Pos2.Y()), math.Max(p.Pos1.Z(), p.Pos2.Z())
}
func (p *Positions) MinXYZ() (float64, float64, float64) {
	return math.Min(p.Pos1.X(), p.Pos2.X()), math.Min(p.Pos1.Y(), p.Pos2.Y()), math.Min(p.Pos1.Z(), p.Pos2.Z())
}
func (p *Positions) BlocksCoordinatesBetween() []mgl64.Vec3 {
	var coords []mgl64.Vec3

	maxX, maxY, maxZ := p.MaxXYZ()
	minX, minY, minZ := p.MinXYZ()

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				coords = append(coords, mgl64.Vec3{x, y, z})
			}
		}
	}
	return coords
}
