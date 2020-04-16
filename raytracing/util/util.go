package util

import (
	"math"
	"math/rand"
)

const (
	// Pi represents our custom pi
	Pi = 3.1415926535897932385
)

var (
	// Infinity cannot be instantiated as constant
	Infinity = math.Inf(1)
)

// Clamp clamps the value x to the range [min,max]
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// RandFloat calculates a random float between 0 and 1
func RandFloat() float64 {
	return rand.Float64()
}

// RandClampFloat calculates a random float between min and max
func RandClampFloat(min, max float64) float64 {
	return Clamp(rand.Float64(), min, max)
}

// DegreesToRadians transfers degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return (degrees * Pi) / 180
}
