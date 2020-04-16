package demo

import (
	"fmt"
	"os"

	"github.com/jeinfeldt/raytracer/raytracing/output"
)

const (
	fileName = "scene.png"
)

// Run runs the demo with the default renderer, scene
// and camera as described by the tutorial
// https://raytracing.github.io/books/RayTracingInOneWeekend.html
// returns the output ppm image as string
func Run(width, height int) {
	fmt.Println("starting...")
	renderer := NewRandomRenderer(width, height)
	img := renderer.Render()
	out := output.NewPNG(width, height)
	err := out.Write(fileName, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("all done! created %q", fileName))
}
