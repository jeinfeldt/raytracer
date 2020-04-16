package scene

import (
	"image"

	"github.com/cheggaaa/pb"
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
	// Renderer renders a scene
	Renderer struct {
		width, height int
		scene         Scene
		camera        Camera
	}
)

// NewRenderer factory method to create a new renderer to render an image
// with given width, height camera and scene
func NewRenderer(width, height int, c Camera, s Scene) Renderer {
	return Renderer{
		width:  width,
		height: height,
		camera: c,
		scene:  s,
	}
}

// Render renders the given scene as a ppm file
func (renderer *Renderer) Render() image.Image {
	// create image
	height := renderer.height
	width := renderer.width
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// render scene
	bar := pb.StartNew(height)
	for y := 0; y < height; y++ {
		bar.Increment()
		for x := 0; x < width; x++ {
			colour := renderer.smoothColour(x, y)
			colour.Div(SamplesPerPixel)
			img.Set(x, height-y-1, &colour)
		}
	}
	bar.Finish()

	// return image data
	return img
}

// shootRay shoots a ray through given pixels and returns corresponding colour
func (renderer *Renderer) shootRay(x, y int) vector.Vector3 {
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
func (renderer *Renderer) smoothColour(x, y int) vector.Vector3 {
	colour := vector.NewEmpty()
	// anti aliasing
	for s := 0; s < SamplesPerPixel; s++ {
		colour.Add(renderer.shootRay(x, y))
	}
	return colour
}
