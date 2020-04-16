package demo

// Run runs the demo with the default renderer, scene
// and camera as described by the tutorial
// https://raytracing.github.io/books/RayTracingInOneWeekend.html
// returns the output ppm image as string
func Run(width, height int) string {
	renderer := NewRandomRenderer(width, height)
	return renderer.Render()
}
