package main

import (
	"fmt"   // Standard package for formatting
	"image" // Package for reading images
	"image/color"
	"image/png"
	"os" // Package for handling files
)

func main() {
	// Open image file
	fi := "./img/test.png"
	inputImg, err := os.Open(fi)
	// If there is an error that exists then print it
	if err != nil {
		fmt.Printf("Failed to open %s: %s", fi, err)
		panic(err.Error())
	}
	defer inputImg.Close() // Close file either when the function is finished or there's a panic

	// Decodes image to get its values
	imgData, _, err := image.Decode(inputImg)
	if err != nil {
		panic(err.Error())
	}

	// Greyscale function
	// Get dimensions of image
	imgBound := imgData.Bounds()
	imgWidth, imgHeight := imgBound.Max.X, imgBound.Max.Y
	// Return a grey image with the given bounds
	greyscale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{imgWidth, imgHeight}})
	// Loop over every pixel in the image
	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			imgColour := imgData.At(x, y)                                               // Convert values at a given point to 16-bit per channel RGBA
			r, g, b, _ := imgColour.RGBA()                                              // Stores the separate values of RGBA
			Y := uint16((0.3 * float64(r)) + (0.59 * float64(g)) + (0.11 * float64(b))) // Weighted values of RGB added
			greyColour := color.Gray{uint8(Y >> 8)}                                     // 8 bits are discarded from the 16-bit weighted values
			greyscale.Set(x, y, greyColour)                                             // Set pixels as the new grey colour
		}
	}

	// Encode the image
	newFi, err := os.Create("greyscale.png")
	if err != nil {
		fmt.Printf("Failed to create %s: %s", newFi, err)
		panic(err.Error())
	}
	defer newFi.Close()
	png.Encode(newFi, greyscale)
}
