package demo

import (
	"github.com/jeinfeldt/raytracer/raytracing/camera"
	"github.com/jeinfeldt/raytracer/raytracing/object"
	"github.com/jeinfeldt/raytracer/raytracing/renderer"
	"github.com/jeinfeldt/raytracer/raytracing/scene"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

// NewSimpleRenderer factory for new ppm renderer (basic tutorial render)
func NewSimpleRenderer(width, height int) renderer.Renderer {
	camera := NewSimpleCamera(width, height)
	scene := NewSimpleScene()
	return renderer.New(width, height, camera, scene)
}

// NewSimpleCamera factory for new camera
func NewSimpleCamera(width, height int) camera.Camera {
	lookfrom := vector.New(3, 3.0, 2)
	lookat := vector.New(0.0, 0.0, -1.0)
	vup := vector.New(0.0, 1.0, 0.0)
	focus := vector.Sub(lookfrom, lookat)
	focusDist := focus.Length()
	aperture := 2.0
	ratio := float64(width) / float64(height)
	return camera.New(lookfrom, lookat, vup, 20, ratio, aperture, focusDist)
}

// NewSimpleBackground factory for background
func NewSimpleBackground() object.GradientBackground {
	return object.NewBackground(vector.New(0.5, 0.7, 1.0),
		vector.New(1.0, 1.0, 1.0))
}

// NewSimpleScene factory for scene
func NewSimpleScene() scene.Scene {
	return scene.New(NewSimpleWorld(), NewSimpleBackground())
}

// NewSimpleWorld factory for world
func NewSimpleWorld() object.World {
	world := object.NewWorld()

	// red sphere in the middle
	matLamb := object.NewLambertian(vector.NewRed())
	red := object.NewSphere(vector.New(0, 0, -1), 0.5, &matLamb)
	world.Add(&red)

	// metal sphere at the side
	matMetal := object.NewMetal(vector.New(0.7, 0.6, 0.5), 0.3)
	metal := object.NewSphere(vector.New(1, 0, -1), 0.5, &matMetal)
	world.Add(&metal)

	// glass sphere at the side
	matGlass := object.NewDielectric(1.5)
	glass := object.NewSphere(vector.New(-1, 0, -1), 0.5, &matGlass)
	world.Add(&glass)

	// green sphere representing underground
	matGround := object.NewLambertian(vector.New(0.8, 0.8, 0))
	underground := object.NewSphere(vector.New(0, -100.5, -1),
		100, &matGround)
	world.Add(&underground)
	return world
}
