package canny

import (
	"github.com/hhharrisonnn/canny-edge-detector/imgcreation"
	"image"
	"image/color"
	"math"
)

// Return points from the x kernel
func sobelX(i, j int) float64 {
	// Matrix for the x direction
	xMat := [3][3]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	return xMat[i][j]
}

// Return points from the y kernel
func sobelY(i, j int) float64 {
	// Matrix for the y direction
	yMat := [3][3]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	return yMat[i][j]
}

func SobelConvolution() (*image.Gray, [][]float64) {
	imageIndex, imgWidth, imgHeight := imgcreation.GetGrey("./img/gaussian.png")

	// Initialise slice to store the angle of each gradient
	theta := make([][]float64, imgWidth-1)
	for i := range theta {
		theta[i] = make([]float64, imgHeight-1)
	}

	// Stores final image values
	// Must be one pixel shorter on each side, otherwise the sobel kernels would be out of bound
	newImage := image.NewGray((image.Rectangle{image.Point{1, 1}, image.Point{imgWidth - 1, imgHeight - 1}}))
	// Iterate over image pixels, get surrounding 3x3 area, convolve with each sobel kernels
	for j := 1; j < imgHeight-1; j++ {
		for i := 1; i < imgWidth-1; i++ {
			// Sobel convolution for x direction
			Gx := imageIndex[i-1][j+1]*sobelX(0, 0) +
				imageIndex[i][j+1]*sobelX(1, 0) +
				imageIndex[i+1][j+1]*sobelX(2, 0) +
				imageIndex[i-1][j]*sobelX(0, 1) +
				imageIndex[i][j]*sobelX(1, 1) +
				imageIndex[i+1][j]*sobelX(2, 1) +
				imageIndex[i-1][j-1]*sobelX(0, 2) +
				imageIndex[i+1][j-1]*sobelX(1, 2) +
				imageIndex[i+1][j-1]*sobelX(2, 2)

			// Sobel convolution for y direction
			Gy := imageIndex[i-1][j+1]*sobelY(0, 0) +
				imageIndex[i][j+1]*sobelY(1, 0) +
				imageIndex[i+1][j+1]*sobelY(2, 0) +
				imageIndex[i-1][j]*sobelY(0, 1) +
				imageIndex[i][j]*sobelY(1, 1) +
				imageIndex[i+1][j]*sobelY(2, 1) +
				imageIndex[i-1][j-1]*sobelY(0, 2) +
				imageIndex[i+1][j-1]*sobelY(1, 2) +
				imageIndex[i+1][j-1]*sobelY(2, 2)

			// Get magnitude of gradients for the two directions
			G := math.Abs(Gx) + math.Abs(Gy)

			// Direction of gradient
			theta[i][j] = math.Atan2(Gy, Gx) * 180 / math.Pi

			// Magnitude of gradients go to the final image
			greyColour := color.Gray{uint8(G)}
			newImage.Set(i, j, greyColour)
		}
	}
	return newImage, theta
}

func Sobel() {
	newImage, _ := SobelConvolution()
	imgcreation.Encode(newImage, "sobel")
}
