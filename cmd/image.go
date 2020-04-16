package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jeinfeldt/raytracer/demo"
	"github.com/spf13/cobra"
)

const (
	// fallback width of scene
	fallbackWidth = 200
	// fallback height of scene
	fallbackHeight = 100
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "This command enables you to create an image",
	Long:  `Image will be created using raytracing, provide width and height of your image`,
	Run: func(cmd *cobra.Command, args []string) {
		// just single argument, error
		if len(args) == 1 {
			fmt.Println("Please run without arguments for fallback")
			fmt.Println("Or run with width and height")
			os.Exit(1)
		}
		// no arguments, run with default
		width := fallbackWidth
		height := fallbackHeight
		// width and height given, pass to demo
		if len(args) == 2 {
			width, _ = strconv.Atoi(args[0])
			height, _ = strconv.Atoi(args[1])
		}
		// render scene - background is a gradient from blue to white
		// with a multiple spheres
		demo.Run(width, height)
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
