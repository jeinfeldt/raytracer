package object

import (
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// GradientBackground defines a background consisting of a
	// gradient
	GradientBackground struct {
		from, to vector.Vector3
	}
)

// NewBackground is a factory method creating a new GradientBackground
func NewBackground(from, to vector.Vector3) GradientBackground {
	return GradientBackground{
		from: from,
		to:   to,
	}
}

// Colour returns the colour of the background depending
// on where the ray hits
func (g *GradientBackground) Colour(r Ray) vector.Vector3 {
	unitDir := vector.Unit(r.Direction())
	weight := 0.5*unitDir.Y() + 1.0
	from := vector.Mul(g.from, weight)
	to := vector.Mul(g.to, (1.0 - weight))
	result := vector.Add(to, from)
	return result
}
