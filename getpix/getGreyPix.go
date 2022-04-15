package getpix

import "image"

func GetGrey(imgData image.Image) ([][]float64, int, int) {
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
