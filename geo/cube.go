package geo

// Cube represents a cubical object, which has sides of equal length and width.
type Cube struct {
	R int
}

// Inside checks if a specific point is within the cube with the centre coordinates passed.
func (c Cube) Inside(cx, cy, cz, x, y, z int) bool {
	return x <= cx+c.R && x >= cx-c.R && y <= cy+c.R && y >= cy-c.R && z <= cz+c.R && z >= cz-c.R
}

// Dim returns the width, height and length of the cube.
func (c Cube) Dim() [3]int {
	return [3]int{c.R*2 + 1, c.R*2 + 1, c.R*2 + 1}
}
