package brush

// Shape represents a shape in which world editing actions may be performed. Any Shape may be combined with an
// act.Action to produce a specific result.
type Shape interface {
	// Inside checks if a specific X, Y and Z is within the shape with centre position (cx, cy, cz). If this is
	// the case, Inside returns true.
	Inside(cx, cy, cz, x, y, z int) bool
	// Dim returns the dimensions of the shape in blocks.
	Dim() [3]int
}
