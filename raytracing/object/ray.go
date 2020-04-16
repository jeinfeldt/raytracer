package object

import (
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// Ray represents a function with an origin and a direction
	Ray struct {
		origin, direction vector.Vector3
	}
)

// NewEmptyRay creates a new ray based on an (0, 0, 0)
// and direction
func NewEmptyRay() Ray {
	return Ray{
		origin:    vector.NewEmpty(),
		direction: vector.NewEmpty(),
	}
}

// NewRay creates a new ray based on an origin point and a
// direction vector
func NewRay(origin vector.Vector3, direction vector.Vector3) Ray {
	return Ray{
		origin:    origin,
		direction: direction,
	}
}

// At returns the point at which the ray would pass given parameter t
func (r *Ray) At(param float64) vector.Vector3 {
	return vector.Add(r.origin, vector.Mul(r.direction, param))
}

// Direction returns a rays direction vector
func (r *Ray) Direction() vector.Vector3 {
	return r.direction
}

// Origin returns a rays origin point
func (r *Ray) Origin() vector.Vector3 {
	return r.origin
}
