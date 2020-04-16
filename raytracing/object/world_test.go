package object_test

import (
	"testing"

	"github.com/jeinfeldt/raytracer/raytracing/vector"

	"github.com/jeinfeldt/raytracer/raytracing/object"
	"gotest.tools/assert"
)

func TestWorld(t *testing.T) {
	t.Run("test Add() and Clear()", func(t *testing.T) {
		world := object.NewWorld()
		// ensure empty
		assert.Equal(t, 0, len(world))

		// add something
		sphere1 := object.NewSphere(vector.Vector3{}, 0, nil)
		world.Add(&sphere1)
		sphere2 := object.NewSphere(vector.Vector3{}, 0, nil)
		world.Add(&sphere2)
		assert.Equal(t, 2, len(world))

		// ensure clear
		world.Clear()
		assert.Equal(t, 0, len(world))
	})
}
