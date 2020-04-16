package scene

import (
	"math"

	"github.com/jeinfeldt/raytracer/raytracing/object"
	"github.com/jeinfeldt/raytracer/raytracing/util"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

type (
	// Camera defines the eye looking at a scene
	Camera struct {
		origin          vector.Vector3
		lowerLeftCorner vector.Vector3
		horizontal      vector.Vector3
		vertical        vector.Vector3
		u               vector.Vector3
		v               vector.Vector3
		w               vector.Vector3
		lensRadius      float64
	}
)

// NewCamera factory method to fetch a new camera
func NewCamera(
	lookfrom, lookat, vup vector.Vector3,
	vfov, aspect, aperture, focusDist float64,
) Camera {
	// helper
	theta := util.DegreesToRadians(vfov)
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	// calculate camera data
	origin := lookfrom.Copy()
	lensRadius := aperture / 2
	w := vector.Unit(vector.Sub(origin, lookat))
	u := vector.Unit(vector.Cross(vup, w))
	v := vector.Cross(w, u)
	lowerLeftCorner := vector.Sub(
		origin, vector.Mul(u, halfWidth, focusDist),
		vector.Mul(v, halfHeight, focusDist),
		vector.Mul(w, focusDist))
	horizontal := vector.Mul(u, 2, halfWidth, focusDist)
	vertical := vector.Mul(v, 2, halfHeight, focusDist)

	// init camera
	return Camera{
		origin:          origin,
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
		lensRadius:      lensRadius,
		w:               w,
		u:               u,
		v:               v,
	}
}

// Ray creates a ray determined by given parameters
func (c *Camera) Ray(s, t float64) object.Ray {
	rd := vector.Mul(vector.RandomUnitDisk(), c.lensRadius)
	offset := vector.Add(vector.Mul(c.u, rd.X()), vector.Mul(c.v, rd.Y()))
	sum := vector.Add(c.lowerLeftCorner, vector.Mul(c.horizontal, s), vector.Mul(c.vertical, t))
	return object.NewRay(vector.Add(c.origin, offset), vector.Sub(sum, c.origin, offset))
}
