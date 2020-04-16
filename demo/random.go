package demo

import (
	"github.com/jeinfeldt/raytracer/raytracing/object"
	"github.com/jeinfeldt/raytracer/raytracing/renderer"
	"github.com/jeinfeldt/raytracer/raytracing/scene"
	"github.com/jeinfeldt/raytracer/raytracing/util"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
)

// NewRandomRenderer factory for new ppm renderer
func NewRandomRenderer(width, height int) renderer.Renderer {
	camera := NewDefaultCamera(width, height)
	scene := NewRandomScene()
	return renderer.New(width, height, camera, scene)
}

// NewRandomScene factory for scene
func NewRandomScene() scene.Scene {
	return scene.New(NewRandomWorld(), NewSimpleBackground())
}

// NewRandomWorld factory for world
func NewRandomWorld() object.World {
	world := object.NewWorld()
	// grey sphere representing underground
	m := object.NewLambertian(vector.New(0.5, 0.5, 0.5))
	underground := object.NewSphere(vector.New(0, -1000, 0), 1000, &m)
	world.Add(&underground)
	// create random spheres
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := util.RandFloat()
			center := vector.New(
				float64(a)+0.9*util.RandFloat(),
				0.2,
				float64(b)+0.9*util.RandFloat(),
			)
			length := vector.Sub(center, vector.New(4, 0.2, 0))
			if length.Length() > 0.9 {
				if chooseMat < 0.6 {
					// diffuse
					albedo := vector.MulVector(vector.Random(), vector.Random())
					m := object.NewLambertian(albedo)
					sphere := object.NewSphere(center, 0.2, &m)
					world.Add(&sphere)
					continue
				}
				if chooseMat < 0.9 {
					// metal
					albedo := vector.Random()
					fuzz := 0.2
					m := object.NewMetal(albedo, fuzz)
					sphere := object.NewSphere(center, 0.2, &m)
					world.Add(&sphere)
					continue
				}
				// glass
				m := object.NewDielectric(1.5)
				sphere := object.NewSphere(center, 0.2, &m)
				world.Add(&sphere)
			}
		}
	}
	// add large spheres
	matGlass := object.NewDielectric(1.5)
	glass := object.NewSphere(vector.New(0, 1, 0), 1.0, &matGlass)
	world.Add(&glass)

	matLamb := object.NewLambertian(vector.NewRed())
	red := object.NewSphere(vector.New(-4, 1, 0), 1.0, &matLamb)
	world.Add(&red)

	matMetal := object.NewMetal(vector.New(0.7, 0.6, 0.5), 0.0)
	metal := object.NewSphere(vector.New(4, 1, 0), 1.0, &matMetal)
	world.Add(&metal)
	return world
}
