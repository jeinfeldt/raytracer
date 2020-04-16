package vector

import (
	"math"

	"github.com/jeinfeldt/raytracer/raytracing/util"
)

const (
	// toHex used to transform values to hex
	toHex float64 = 0xffff
)

// helper variables for standard colours

// NewBlack is a factory method to create a new vector representing black
func NewBlack() Vector3 {
	return NewEmpty()
}

// NewWhite is a factory method to create a new vector representing white
func NewWhite() Vector3 {
	return New(1, 1, 1)
}

// NewRed is a factory method to create a new vector representing red
func NewRed() Vector3 {
	return New(1, 0, 0)
}

// NewGreen is a factory method to create a new vector representing green
func NewGreen() Vector3 {
	return New(1, 0, 0)
}

// NewBlue is a factory method to create a new vector representing blue
func NewBlue() Vector3 {
	return New(0, 0, 1)
}

// RGBA to implement color interface from image package
// transforms values to rgb compatible values
func (v *Vector3) RGBA() (r, g, b, a uint32) {
	// Sqrt() for gamma-2 correction
	// color needs to be in hex
	r = uint32(util.Clamp(math.Sqrt(v.x), 0, 0.999) * toHex)
	g = uint32(util.Clamp(math.Sqrt(v.y), 0, 0.999) * toHex)
	b = uint32(util.Clamp(math.Sqrt(v.z), 0, 0.999) * toHex)
	a = uint32(toHex)
	return
}
