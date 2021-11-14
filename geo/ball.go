package geo

// Ball represents the space bounded by a sphere. It is a spherical object with a specific radius.
type Ball struct {
	R int
}

// Inside checks if a specific point is within the ball at the centre x, y and z passed.
func (b Ball) Inside(cx, cy, cz, x, y, z int) bool {
	dx, dy, dz := x-cx, y-cy, z-cz
	return dx*dx+dy*dy+dz*dz <= b.R*b.R
}

// Dim returns the width, height and length of the ball.
func (b Ball) Dim() [3]int {
	return [3]int{b.R*2 + 1, b.R*2 + 1, b.R*2 + 1}
}
