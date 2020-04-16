package object

import (
	"math"

	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// Sphere describes a sphere on the scene
	Sphere struct {
		center   vector.Vector3
		radius   float64
		material Material
	}
)

// NewSphere is a factory method to create a new sphere
func NewSphere(center vector.Vector3, radius float64, m Material) Sphere {
	return Sphere{
		center:   center,
		radius:   radius,
		material: m,
	}
}

// Hit indicates if and where the sphere was hit
func (s *Sphere) Hit(r Ray, tmin, tmax float64, record *HitRecord) bool {
	oc := vector.Sub(r.Origin(), s.center)
	direction := r.Direction()
	a := direction.LengthSquared()
	halfB := vector.Dot(oc, r.Direction())
	c := oc.LengthSquared() - s.radius*s.radius
	discriminant := halfB*halfB - a*c
	// evaluate discriminant
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		temp := (halfB*-1 - root) / a
		if temp < tmax && temp > tmin {
			*record = *s.updateRecord(record, r, temp)
			return true
		}
		temp = (halfB*-1 + root) / a
		if temp < tmax && temp > tmin {
			*record = *s.updateRecord(record, r, temp)
			return true
		}
	}
	return false
}

// updateRecord updates the hit record with the hit data
func (s *Sphere) updateRecord(record *HitRecord, r Ray, t float64) *HitRecord {
	record.Time = t
	record.Point = r.At(t)
	record.Material = s.material
	outward := vector.Div(vector.Sub(record.Point, s.center), s.radius)
	record.SetFaceNormal(r, outward)
	return record
}
