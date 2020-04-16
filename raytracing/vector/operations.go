package vector

// Add adds two vectors
// Example: Add({1, 2, 0}, {1, 0, 1}) = {2, 0, 1}
func Add(this Vector3, vectors ...Vector3) Vector3 {
	result := this.Copy()
	for _, vector := range vectors {
		result.Add(vector)
	}
	return result
}

// Sub subs two or more vectors
// Example: Sub({{1, 2, 0}, 1, 0, 1}) = {0, 2, -1}
func Sub(this Vector3, vectors ...Vector3) Vector3 {
	result := this.Copy()
	for _, vector := range vectors {
		result.Sub(vector)
	}
	return result
}

// Mul multiplies a vector with a constant factor
// Example: Mul({1, 2, 0}, 2) = {2, 4, 0}
func Mul(vector Vector3, factors ...float64) Vector3 {
	result := vector.Copy()
	for _, factor := range factors {
		result.Mul(factor)
	}
	return result
}

// MulVector multiplies two vector with their respective positions
// Example: Mul({1, 2, 0}, {1, 3, 1}) = {1, 6, 0}
func MulVector(this Vector3, other Vector3) Vector3 {
	return New(this.X()*other.X(), this.Y()*other.Y(), this.Z()*other.Z())
}

// Div divides a vector with a constant factor
// Example: Div({4, 2, 0}, 2) = {2, 1, 0}
func Div(vector Vector3, factor float64) Vector3 {
	result := vector.Copy()
	result.Div(factor)
	return result
}

// Dot calculates the dot product of two vectors
// This refers to the sum of the products of the corresponding entries
// of the two sequences of numbers
// Example: Dot({2, 1, 2}, {0, 1, 3}) = 2*0 + 1*1 + 2*3 = 7
func Dot(this Vector3, other Vector3) float64 {
	x := this.X() * other.X()
	y := this.Y() * other.Y()
	z := this.Z() * other.Z()
	return x + y + z
}

// Cross calculates the cross product of two vectors
// This refers to a vector c that is perpendicular (orthogonal)
// to both given vectors a and b
// Example: Cross({2,3,4}, {5,6,7}) = {-3, 6, -3}
func Cross(this Vector3, other Vector3) Vector3 {
	x := this.Y()*other.Z() - this.Z()*other.Y()
	y := this.Z()*other.X() - this.X()*other.Z()
	z := this.X()*other.Y() - this.Y()*other.X()
	return New(x, y, z)
}

// Unit calculates the unit vector of given vector
// A unit vector is a vector of length 1
func Unit(this Vector3) Vector3 {
	copy := this.Copy()
	copy.Div(copy.Length())
	return copy
}
