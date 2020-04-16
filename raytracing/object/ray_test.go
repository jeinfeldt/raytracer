package object_test

import (
	"testing"

	"github.com/jeinfeldt/raytracer/raytracing/object"
	"github.com/jeinfeldt/raytracer/raytracing/vector"
	"github.com/stretchr/testify/assert"
)

func TestRay(t *testing.T) {
	t.Run("test At()", func(t *testing.T) {
		// ray has origin in zero coordinate
		ray := object.NewRay(vector.NewEmpty(), vector.New(1, 0, 0))
		actual := ray.At(0)
		assert.Equal(t, vector.NewEmpty(), actual)

		actual = ray.At(5)
		assert.Equal(t, vector.New(5, 0, 0), actual)

		actual = ray.At(-5)
		assert.Equal(t, vector.New(-5, 0, 0), actual)
	})
}
