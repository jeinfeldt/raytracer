package output

import (
	"image"
	"image/png"
	"os"
)

type (
	// PNGWriter writes an image to a png file
	PNGWriter struct {
		width, height int
	}
)

// NewPNG is a factory method to create a new ppm writer
func NewPNG(width, height int) *PNGWriter {
	return &PNGWriter{
		width:  width,
		height: height,
	}
}

// Writes writes image to png file
func (writer *PNGWriter) Write(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// make sure to close file
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// write png
	err = png.Encode(file, img)
	return err
}
