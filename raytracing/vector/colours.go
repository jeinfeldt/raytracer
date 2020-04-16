package vector

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
