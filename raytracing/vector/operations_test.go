package vector_test

import (
	"testing"

	"github.com/jeinfeldt/raytracer/raytracing/vector"
	"github.com/stretchr/testify/assert"
)

func TestOps(t *testing.T) {
	t.Run("test Add()", func(t *testing.T) {
		actual := vector.Add(vector.New(1, 0, 4), vector.New(1, 1, 1))
		assert.Equal(t, 2.0, actual.X())
		assert.Equal(t, 1.0, actual.Y())
		assert.Equal(t, 5.0, actual.Z())
	})

	t.Run("test Sub()", func(t *testing.T) {
		actual := vector.Sub(vector.New(1, 0, 4), vector.New(1, 1, 1))
		assert.Equal(t, 0.0, actual.X())
		assert.Equal(t, -1.0, actual.Y())
		assert.Equal(t, 3.0, actual.Z())
	})

	t.Run("test Mul()", func(t *testing.T) {
		actual := vector.Mul(vector.New(1, 0, 4), 2)
		assert.Equal(t, 2.0, actual.X())
		assert.Equal(t, 0.0, actual.Y())
		assert.Equal(t, 8.0, actual.Z())
	})

	t.Run("test Div()", func(t *testing.T) {
		actual := vector.Div(vector.New(4, 2, 0), 2)
		assert.Equal(t, 2.0, actual.X())
		assert.Equal(t, 1.0, actual.Y())
		assert.Equal(t, 0.0, actual.Z())
	})

	t.Run("test Dot()", func(t *testing.T) {
		this := vector.New(2, 1, 2)
		other := vector.New(0, 1, 3)
		actual := vector.Dot(this, other)
		assert.Equal(t, 7.0, actual)
	})

	t.Run("test Cross()", func(t *testing.T) {
		this := vector.New(2, 3, 4)
		other := vector.New(5, 6, 7)
		actual := vector.Cross(this, other)
		assert.Equal(t, vector.New(-3, 6, -3), actual)
	})

	t.Run("test Unit()", func(t *testing.T) {
		this := vector.New(0, 0, 9)
		assert.Equal(t, vector.New(0, 0, 1), vector.Unit(this))
	})

	t.Run("test MulVector()", func(t *testing.T) {
		this := vector.New(1, 2, 0)
		other := vector.New(1, 3, 1)
		actual := vector.MulVector(this, other)
		assert.Equal(t, vector.New(1, 6, 0), actual)
	})
}
