package geo

import (
	"github.com/df-mc/we/brush"
)

// init registers all brush.Shape implementations in the geo package.
func init() {
	brush.RegisterShape("Ball", func(r int) brush.Shape { return Ball{R: r} })
	brush.RegisterShape("Cube", func(r int) brush.Shape { return Cube{R: r} })
}
