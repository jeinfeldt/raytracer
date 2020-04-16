package renderer

import (
	"github.com/cheggaaa/pb"
	"github.com/jeinfeldt/raytracer/raytracing/camera"
	"github.com/jeinfeldt/raytracer/raytracing/scene"
	"github.com/jeinfeldt/raytracer/raytracing/util"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

const (
	// SamplesPerPixel for anti aliasing
	SamplesPerPixel = 100
	// MaxDepth max depth for recursion (ray bounce limit)
	MaxDepth = 50
)

type (
	// PPMRenderer renders a scene
	PPMRenderer struct {
		width, height int
		scene         scene.Scene
		camera        camera.Camera
	}
)

// New factory method to create a new renderer to render an image
// with given width, height camera and scene
func New(width, height int, c camera.Camera, s scene.Scene) PPMRenderer {
	return PPMRenderer{
		width:  width,
		height: height,
		camera: c,
		scene:  s,
	}
}

// Render renders the given scene as a ppm file
func (renderer *PPMRenderer) Render() []string {
	height := renderer.height
	width := renderer.width

	// render scene
	pixels := make([]string, 0)
	bar := pb.StartNew(height)
	for y := height - 1; y >= 0; y-- {
		bar.Increment()
		for x := 0; x < width; x++ {
			colour := renderer.smootColour(x, y)
			pixels = append(pixels, colour.WriteColour(SamplesPerPixel))
		}
	}
	bar.Finish()

	// return image data
	return pixels
}

// shootRay shoots a ray through given pixels and returns corresponding colour
func (renderer *PPMRenderer) shootRay(x, y int) vector.Vector3 {
	height := renderer.height
	width := renderer.width
	scene := renderer.scene
	u := (float64(x) + util.RandFloat()) / float64(width)
	v := (float64(y) + util.RandFloat()) / float64(height)
	ray := renderer.camera.Ray(u, v)
	colour := scene.Colour(ray, MaxDepth)
	return colour
}

// smoothColour smoothes pixel colour for anti aliasing (smooth edges)
// this is done by shooting a ray multiple times add a pixel
// and fetching the aggregated colour
func (renderer *PPMRenderer) smootColour(x, y int) vector.Vector3 {
	colour := vector.Vector3{}
	// anti aliasing
	for s := 0; s < SamplesPerPixel; s++ {
		colour = vector.Add(colour, renderer.shootRay(x, y))
	}
	return colour
}
