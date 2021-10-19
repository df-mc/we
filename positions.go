package worldedit

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Positions struct {
	Pos1 *mgl64.Vec3
	Pos2 *mgl64.Vec3
}

func (p *Positions) MaxX() float64 { return biggerFloat(p.Pos1.X(), p.Pos2.X()) }
func (p *Positions) MaxY() float64 { return biggerFloat(p.Pos1.Y(), p.Pos2.Y()) }
func (p *Positions) MaxZ() float64 { return biggerFloat(p.Pos1.Z(), p.Pos2.Z()) }

func (p *Positions) MinX() float64 { return smallerFloat(p.Pos1.X(), p.Pos2.X()) }
func (p *Positions) MinY() float64 { return smallerFloat(p.Pos1.Y(), p.Pos2.Y()) }
func (p *Positions) MinZ() float64 { return smallerFloat(p.Pos1.Z(), p.Pos2.Z()) }

func (p *Positions) BlocksCoordinatesBetween() []mgl64.Vec3 {
	var coords []mgl64.Vec3
	for x := p.MinX(); x <= p.MaxX(); x++ {
		for y := p.MinY(); y <= p.MaxY(); y++ {
			for z := p.MinZ(); z <= p.MaxZ(); z++ {
				coords = append(coords, mgl64.Vec3{x, y, z})
			}
		}
	}
	return coords
}

func biggerFloat(num1, num2 float64) float64 {
	if num1 >= num2 {
		return num1
	}
	return num2
}
func smallerFloat(num1, num2 float64) float64 {
	if biggerFloat(num1, num2) == num1 {
		return num2
	}
	return num1
}
