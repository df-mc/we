package geo

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
)

// Area contains the points with Min.X <= X <= Max.X, Min.Y <= Y <= Max.Y,
// and Min.Z <= Z <= Max.Z
// It is well-formed if Min.X <= Max.X and likewise for Y and Z. Points are
// always well-formed. An area's methods always return well-formed outputs
// for well-formed inputs.
type Area struct {
	Min, Max cube.Pos
}

// NewArea is shorthand for Area{Pos(x0, y0, z0), Pos(x1, y1, z0)}. The returned
// Area has minimum and maximum coordinates swapped if necessary so that
// it is well-formed.
func NewArea(x0, y0, z0, x1, y1, z1 int) Area {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	if z0 > z1 {
		z0, z1 = z1, z0
	}
	return Area{Min: cube.Pos{x0, y0, z0}, Max: cube.Pos{x1, y1, z1}}
}

// String returns a string representation of the Area like "(1,2,3)-(4,5,6)".
func (a Area) String() string {
	return fmt.Sprintf("%v-%v", a.Min, a.Max)
}

// Dx returns the Area's width.
func (a Area) Dx() int {
	return a.Max[0] - a.Min[0] + 1
}

// Dy returns the Area's height.
func (a Area) Dy() int {
	return a.Max[1] - a.Min[1] + 1
}

// Dz returns the Area's length.
func (a Area) Dz() int {
	return a.Max[2] - a.Min[2] + 1
}

// Range iterates over all points where Min.X <= X <= Max.X, Min.Y <= Y <= Max.Y,
// and Min.Z <= Z <= Max.Z and calls f for every X, Y and Z.
func (a Area) Range(f func(x, y, z int)) {
	for x := a.Min[0]; x <= a.Max[0]; x++ {
		for y := a.Min[1]; y <= a.Max[1]; y++ {
			for z := a.Min[2]; z <= a.Max[2]; z++ {
				f(x, y, z)
			}
		}
	}
}
