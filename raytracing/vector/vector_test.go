package vector_test

import (
	"testing"

	"github.com/jeinfeldt/raytracer/raytracing/vector"
	"github.com/stretchr/testify/assert"
)

func TestVector3_Basics(t *testing.T) {
	t.Run("test Add()", func(t *testing.T) {
		v := vector.New(1, 0, 4)
		v.Add(vector.New(1, 1, 1))
		assert.Equal(t, 2.0, v.X())
		assert.Equal(t, 1.0, v.Y())
		assert.Equal(t, 5.0, v.Z())
	})

	t.Run("test Sub()", func(t *testing.T) {
		v := vector.New(1, 0, 4)
		v.Sub(vector.New(1, 1, 1))
		assert.Equal(t, 0.0, v.X())
		assert.Equal(t, -1.0, v.Y())
		assert.Equal(t, 3.0, v.Z())
	})

	t.Run("test Mul()", func(t *testing.T) {
		v := vector.New(1, 0, 4)
		v.Mul(2)
		assert.Equal(t, 2.0, v.X())
		assert.Equal(t, 0.0, v.Y())
		assert.Equal(t, 8.0, v.Z())
	})

	t.Run("test Div()", func(t *testing.T) {
		v := vector.New(4, 2, 0)
		v.Div(2)
		assert.Equal(t, 2.0, v.X())
		assert.Equal(t, 1.0, v.Y())
		assert.Equal(t, 0.0, v.Z())
	})

	t.Run("test Copy()", func(t *testing.T) {
		v := vector.New(4, 2, 0)
		copy := v.Copy()
		assert.Equal(t, 4.0, copy.X())
	})
}

func TestVector3_Output(t *testing.T) {
	t.Run("test String()", func(t *testing.T) {
		v := vector.New(2, 1, 2)
		assert.Equal(t, "{X:2.000000, Y:1.000000, Z:2.000000}", v.String())
	})
}

func TestVector3_Features(t *testing.T) {
	t.Run("test Length()", func(t *testing.T) {
		v := vector.New(2, 1, 2)
		assert.Equal(t, 3.0, v.Length())
	})

	t.Run("test LengthSquared()", func(t *testing.T) {
		v := vector.New(2, 1, 2)
		assert.Equal(t, 9.0, v.LengthSquared())
	})
}
