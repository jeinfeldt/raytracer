package scene

import (
	"github.com/jeinfeldt/raytracer/raytracing/object"
	"github.com/jeinfeldt/raytracer/raytracing/util"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// Scene represents a scene for rendering including a world
	// with objects and a background
	Scene struct {
		world      object.World
		background object.GradientBackground
	}
)

// New is a factory method for a scene consisting of a world
// of objects and a background
func New(w object.World, b object.GradientBackground) Scene {
	return Scene{
		world:      w,
		background: b,
	}
}

// Colour returns the colour as a vector depending on which pixel
// the given ray hits
func (s *Scene) Colour(r object.Ray, depth int) vector.Vector3 {
	record := &object.HitRecord{}

	// recursion anchor, if we hit max no more light is gathered
	if depth <= 0 {
		return vector.NewBlack()
	}

	// guard, in case anything in the world is hit return corresponding colour
	isHit := s.world.Hit(r, 0.001, util.Infinity, record)
	if isHit {
		scattered := object.NewEmptyRay()
		attenuation := vector.NewBlack()

		// calculate ray bounce
		if record.Material.Scatter(r, record, &attenuation, &scattered) {
			return vector.MulVector(attenuation, s.Colour(scattered, depth-1))
		}
		return vector.NewBlack()
	}

	// otherwise draw background
	return s.background.Colour(r)
}
