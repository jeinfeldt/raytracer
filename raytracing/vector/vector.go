package vector

import (
	"fmt"
	"math"

	"github.com/jeinfeldt/raytracer/raytracing/util"
)

const (
	// colourTransformation used to translate decimal values to colour values
	colourTransformation float64 = 256
)

type (
	// Vector3 represent a vector with three coordinates
	Vector3 struct {
		x, y, z float64
	}
)

// New is a factory method to create a new vector with three positions
func New(x, y, z float64) Vector3 {
	return Vector3{x: x, y: y, z: z}
}

// NewEmpty is a factory method to create a new vector with (0, 0, 0)
func NewEmpty() Vector3 {
	return New(0, 0, 0)
}

// X returns the first position of the vector
// Example: {1, 2, 0}.X() = 1
func (v *Vector3) X() float64 {
	return v.x
}

// Y returns the second position of the vector
// Example: {1, 2, 0}.Y() = 2
func (v *Vector3) Y() float64 {
	return v.y
}

// Z returns the third position of the vector
// Example: {1, 2, 0}.Z() = 0
func (v *Vector3) Z() float64 {
	return v.z
}

// Add adds two vectors
// Example: {1, 2, 0}.Add({1, 0, 1}) = {2, 0, 1}
func (v *Vector3) Add(other Vector3) {
	v.x += other.X()
	v.y += other.Y()
	v.z += other.Z()
}

// Sub subs two vectors
// Example: {1, 2, 0}.Sub({1, 0, 1}) = {0, 2, -1}
func (v *Vector3) Sub(other Vector3) {
	v.x -= other.X()
	v.y -= other.Y()
	v.z -= other.Z()
}

// Mul multiplies a vector with a constant factor
// Example: {1, 2, 0}.Mul({2}) = {2, 4, 0}
func (v *Vector3) Mul(factor float64) {
	v.x *= factor
	v.y *= factor
	v.z *= factor
}

// Div divides a vector with a constant factor
// Example: {4, 2, 0}.Div({2}) = {2, 1, 0}
func (v *Vector3) Div(factor float64) {
	v.x /= factor
	v.y /= factor
	v.z /= factor
}

// Length returns the length of the vector
// The length of a vector is the square root of the sum
// of the squares of the horizontal and vertical components
// Example: {2, 1, 2}.length = Sqrt(2*2 + 1*1 + 2*2) = 3
func (v *Vector3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the length of the vector squared
// This equals the sum
// of the squares of the horizontal and vertical components
// Example: {2, 1, 2}.length = 2*2 + 1*1 + 2*2 = 9
func (v *Vector3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// WriteColour writes the translated [0,255] value of each color component.
func (v *Vector3) WriteColour(samples int) string {
	// Divide the color total by the number of samples and gamma-correct
	// for a gamma value of 2.0.
	template := "%d %d %d\n"
	scale := 1.0 / float64(samples)
	r := math.Sqrt(v.x * scale)
	g := math.Sqrt(v.y * scale)
	b := math.Sqrt(v.z * scale)
	x := int32(util.Clamp(r, 0, 0.999) * colourTransformation)
	y := int32(util.Clamp(g, 0, 0.999) * colourTransformation)
	z := int32(util.Clamp(b, 0, 0.999) * colourTransformation)
	return fmt.Sprintf(template, x, y, z)
}

// String output for vector
// Example: Vector3{1, 2, 3}.String() = {X:1, Y:2, Z:3}
func (v *Vector3) String() string {
	template := "{X:%f, Y:%f, Z:%f}"
	return fmt.Sprintf(template, v.x, v.y, v.z)
}

// Copy returns a new reference to an equal vector
func (v *Vector3) Copy() Vector3 {
	return Vector3{v.x, v.y, v.z}
}

// Random returns a randomely initialised vector
func Random() Vector3 {
	return Vector3{
		x: util.RandFloat(),
		y: util.RandFloat(),
		z: util.RandFloat()}
}

// RandomClamp returns a randomely initialised vector with
// min and max values
func RandomClamp(min, max float64) Vector3 {
	return Vector3{
		x: util.RandClampFloat(min, max),
		y: util.RandClampFloat(min, max),
		z: util.RandClampFloat(min, max)}
}

// RandomUnit returns a randomely initialised vector within unit sphere
// meaning length 1
func RandomUnit() Vector3 {
	a := util.RandClampFloat(0, 2*util.Pi)
	z := util.RandClampFloat(-1, 1)
	r := math.Sqrt(1 - z*z)
	return Vector3{
		x: r * math.Cos(a),
		y: r * math.Sin(a),
		z: z,
	}
}

// RandomUnitDisk returns a randomely initialised vector within unit disk
func RandomUnitDisk() Vector3 {
	for {
		p := Vector3{
			x: util.RandClampFloat(-1, 1),
			y: util.RandClampFloat(-1, 1),
			z: 0,
		}
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}
