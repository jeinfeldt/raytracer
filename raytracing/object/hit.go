package object

import (
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// HitRecord indicates where a Hittable object was hit
	HitRecord struct {
		Time          float64
		Point, Normal vector.Vector3
		FrontFace     bool
		Material      Material
	}

	// Hittable indicates that the implementing struct can be hit by
	// a sphere
	Hittable interface {
		Hit(Ray, float64, float64, *HitRecord) bool
	}
)

// SetFaceNormal update FrontFace and Normal information on record
func (h *HitRecord) SetFaceNormal(r Ray, outwardNormal vector.Vector3) {
	h.FrontFace = vector.Dot(r.Direction(), outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
		return
	}
	h.Normal = vector.Mul(outwardNormal, -1)
}
