package getpix

import (
	"fmt"
	"image"
	"os"
)

func GetGrey(fi string) ([][]float64, int, int) {
	// Open greyscale.png
	inputImg, err := os.Open(fi)
	if err != nil {
		fmt.Printf("Failed to open %s: %s", fi, err)
		panic(err.Error())
	}
	defer inputImg.Close()

	imgData, _, err := image.Decode(inputImg)
	if err != nil {
		panic(err.Error())
	}

	// Get image dimensions
	imgBound := imgData.Bounds()
	imgWidth, imgHeight := imgBound.Max.X, imgBound.Max.Y

	// Make 2D slice with dimensions imgWidth x imgHeight
	imageIndex := make([][]float64, imgWidth)
	for i := range imageIndex {
		imageIndex[i] = make([]float64, imgHeight)
	}

	// Iterate over image to get only grey values
	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			imgColour := imgData.At(x, y)
			pixelGreyValue, _, _, _ := imgColour.RGBA()
			Y := uint8(pixelGreyValue)
			greyColour := float64(Y)
			imageIndex[x][y] = greyColour // Add greyColour to imageIndex
		}
	}

	return imageIndex, imgWidth, imgHeight
}
