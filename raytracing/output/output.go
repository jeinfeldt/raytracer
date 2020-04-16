package output

import (
	"image"
)

type (
	// ImageWriter is an abstraction how to write an image to file
	ImageWriter interface {
		Write(path string, img image.Image) error
	}
)
