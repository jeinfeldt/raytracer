package demo

import (
	"github.com/jeinfeldt/raytracer/raytracing/camera"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

// NewDefaultCamera factory for new camera
func NewDefaultCamera(width, height int) camera.Camera {
	lookfrom := vector.New(13, 2.0, 3)
	lookat := vector.New(0.0, 0.0, 0.0)
	vup := vector.New(0.0, 1.0, 0.0)
	focusDist := 10.0
	aperture := 0.1
	ratio := float64(width) / float64(height)
	return camera.New(lookfrom, lookat, vup, 20, ratio, aperture, focusDist)
}
