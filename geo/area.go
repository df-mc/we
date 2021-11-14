package geo

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
)

type Area struct {
	Min, Max cube.Pos
}

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

func (a Area) String() string {
	return fmt.Sprintf("%v-%v", a.Min, a.Max)
}

func (a Area) Dx() int {
	return a.Max[0] - a.Min[0] + 1
}

func (a Area) Dy() int {
	return a.Max[1] - a.Min[1] + 1
}

func (a Area) Dz() int {
	return a.Max[2] - a.Min[2] + 1
}

func (a Area) Range(f func(x, y, z int)) {
	for x := a.Min[0]; x <= a.Max[0]; x++ {
		for y := a.Min[1]; y <= a.Max[1]; y++ {
			for z := a.Min[2]; z <= a.Max[2]; z++ {
				f(x, y, z)
			}
		}
	}
}
