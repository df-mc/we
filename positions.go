package worldedit

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// Positions is a type that contains two vec3 values which are used to know which blocks to act on.
type Positions struct {
	Pos1 mgl64.Vec3
	Pos2 mgl64.Vec3
}

// NewPositions returns a *Positions with the two vec3 values passed
func NewPositions(pos1, pos2 mgl64.Vec3) *Positions { return &Positions{Pos1: pos1, Pos2: pos2} }

// MaxXYZ returns the biggest value of X, Y, Z in the two vec3 values.
func (p *Positions) MaxXYZ() (float64, float64, float64) {
	return math.Max(p.Pos1.X(), p.Pos2.X()), math.Max(p.Pos1.Y(), p.Pos2.Y()), math.Max(p.Pos1.Z(), p.Pos2.Z())
}

// MinXYZ returns the smallest value of X, Y, Z in the two vec3 values.
func (p *Positions) MinXYZ() (float64, float64, float64) {
	return math.Min(p.Pos1.X(), p.Pos2.X()), math.Min(p.Pos1.Y(), p.Pos2.Y()), math.Min(p.Pos1.Z(), p.Pos2.Z())
}

// BlocksCoordinatesBetween returns the blocks coordinates between the two vec3 values.
// This will let us know which coordinate to act on.
func (p *Positions) BlocksCoordinatesBetween() (coords []mgl64.Vec3) {

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
