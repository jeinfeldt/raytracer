package object

import (
	"math"

	"github.com/jeinfeldt/raytracer/raytracing/util"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// Material indicates a certain material which scatters a ray
	Material interface {
		Scatter(r Ray, record *HitRecord, attenuation *vector.Vector3, scattered *Ray) bool
	}

	// Lambertian represents lambertian (diffuse) material
	Lambertian struct {
		albedo vector.Vector3
	}

	// Metal represents reflecting metal material
	Metal struct {
		albedo vector.Vector3
		fuzz   float64
	}

	// Dielectric represents glass that always refracts
	Dielectric struct {
		refIndex float64
	}
)

// NewLambertian is a factory method to create new lambertian material
func NewLambertian(albedo vector.Vector3) Lambertian {
	return Lambertian{
		albedo: albedo,
	}
}

// NewMetal is a factory method to create new metal material
func NewMetal(albedo vector.Vector3, fuzz float64) Metal {
	return Metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

// NewDielectric is a factory method to create new dielectric material
func NewDielectric(refIndex float64) Dielectric {
	return Dielectric{
		refIndex: refIndex,
	}
}

// Scatter function for lambertian (diffuse) material
func (l *Lambertian) Scatter(r Ray, record *HitRecord, attenuation *vector.Vector3, scattered *Ray) bool {
	direction := vector.Add(record.Normal, vector.RandomUnit())
	*scattered = NewRay(record.Point, direction)
	*attenuation = l.albedo
	return true
}

// Scatter function for metal (reflecting) material
func (m *Metal) Scatter(r Ray, record *HitRecord, attenuation *vector.Vector3, scattered *Ray) bool {
	reflected := reflect(vector.Unit(r.Direction()), record.Normal)
	direction := vector.Add(reflected, vector.Mul(vector.RandomUnit(), m.fuzz))
	*scattered = NewRay(record.Point, direction)
	*attenuation = m.albedo
	return vector.Dot(scattered.Direction(), record.Normal) > 0
}

// Scatter function for dielectric (refracting) material
func (d *Dielectric) Scatter(r Ray, record *HitRecord, attenuation *vector.Vector3, scattered *Ray) bool {
	*attenuation = vector.New(1, 1, 1)
	etaiOverEtat := d.refIndex
	if record.FrontFace {
		etaiOverEtat = 1.0 / d.refIndex
	}
	direction := vector.Unit(r.Direction())
	cosTheta := math.Min(vector.Dot(vector.Mul(direction, -1.0), record.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	// reflection
	if etaiOverEtat*sinTheta > 1.0 {
		reflected := reflect(direction, record.Normal)
		*scattered = NewRay(record.Point, reflected)
		return true
	}
	// schlick
	reflectProb := schlick(cosTheta, etaiOverEtat)
	if util.RandFloat() < reflectProb {
		reflected := reflect(direction, record.Normal)
		*scattered = NewRay(record.Point, reflected)
		return true
	}
	// refcration
	refracted := refract(direction, record.Normal, etaiOverEtat)
	*scattered = NewRay(record.Point, refracted)

	return true
}

// reflect is a helper function to calculate reflection
func reflect(v vector.Vector3, n vector.Vector3) vector.Vector3 {
	return vector.Sub(v, vector.Mul(n, vector.Dot(v, n)*2))
}

// refract is a helper function to calculate refraction
func refract(uv vector.Vector3, n vector.Vector3, etaiOverEtat float64) vector.Vector3 {
	cosTheta := vector.Dot(vector.Mul(uv, -1), n)
	rOutParallel := vector.Mul(vector.Add(uv, vector.Mul(n, cosTheta)), etaiOverEtat)
	rOutPerp := vector.Mul(n, math.Sqrt(1.0-rOutParallel.LengthSquared())*-1)
	result := vector.Add(rOutParallel, rOutPerp)
	return result
}

// Schlick approximation is a helper to introduce more realistic glass
// reflection
func schlick(cosine, refIndex float64) float64 {
	r0 := (1 - refIndex) / (1 + refIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
