package demo

import (
	"github.com/jeinfeldt/raytracer/raytracing/output"
)

// Run runs the demo with the default renderer, scene
// and camera as described by the tutorial
// https://raytracing.github.io/books/RayTracingInOneWeekend.html
// returns the output ppm image as string
func Run(width, height int) {
	renderer := NewRandomRenderer(width, height)
	pixels := renderer.Render()
	out := output.NewPPM(pixels, width, height)
	out.Write("", nil)
}
