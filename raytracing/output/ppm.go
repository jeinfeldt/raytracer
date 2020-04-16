package output

import (
	"fmt"
	"image"
	"strings"
)

const (
	// Format standard format for ppm
	Format = "P3"
	// MaxColour value for ppm format
	MaxColour = 255
)

type (
	// PPMWriter an image to a ppm file
	PPMWriter struct {
		pixels        []string
		width, height int
	}
)

// NewPPM is a factory method to create a new ppm writer
func NewPPM(pixels []string, width, height int) *PPMWriter {
	return &PPMWriter{
		pixels: pixels,
		width:  width,
		height: height,
	}
}

// Writes writes image to ppm file
func (ppm *PPMWriter) Write(_ string, _ image.Image) error {
	header := ppm.header(Format, ppm.width, ppm.height, MaxColour)
	data := strings.Join(ppm.pixels, "")
	data = strings.TrimSuffix(data, "%")
	fmt.Print(header + data)
	return nil
}

// header helper function to build the header for the ppm file
func (ppm *PPMWriter) header(format string, width, height, maxColour int) string {
	template := "%s\n%d %d\n%d\n"
	return fmt.Sprintf(template, format, width, height, maxColour)
}
