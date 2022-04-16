package canny

import (
	"fmt"
	"github.com/hhharrisonnn/canny-edge-detector/getpix"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/user"
)

// Return points relative to midpoint of x kernel
func sobelX(i, j int) float64 {
	// Matrix for the x direction
	xMat := [3][3]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	return xMat[i+1][j+1]
}

// Return points relative to midpoint of y kernel
func sobelY(i, j int) float64 {
	// Matrix for the y direction
	yMat := [3][3]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	return yMat[i+1][j+1]
}

func SobelConvolution() {
	imageIndex, imgWidth, imgHeight := getpix.GetGrey("./img/gaussianBlur.png")

	// Stores final image values
	// Must be one pixel shorter on each side, otherwise the sobel kernels would be out of bound
	newImage := image.NewGray((image.Rectangle{image.Point{1, 1}, image.Point{imgWidth - 1, imgHeight - 1}}))
	// Iterate over image pixels, get surrounding 3x3 area, convolve with each sobel kernels
	for j := 1; j < imgHeight-1; j++ {
		for i := 1; i < imgWidth-1; i++ {
			// Sobel convolution for x direction
			Gx := imageIndex[i-1][j+1]*sobelX(-1, 1) +
				imageIndex[i][j+1]*sobelX(0, 1) +
				imageIndex[i+1][j+1]*sobelX(1, 1) +
				imageIndex[i-1][j]*sobelX(-1, 1) +
				imageIndex[i][j]*sobelX(0, 0) +
				imageIndex[i+1][j]*sobelX(1, 0) +
				imageIndex[i-1][j-1]*sobelX(-1, -1) +
				imageIndex[i+1][j-1]*sobelX(1, -1) +
				imageIndex[i+1][j-1]*sobelX(1, -1)

			// Sobel convolution for y direction
			Gy := imageIndex[i-1][j+1]*sobelX(-1, 1) +
				imageIndex[i][j+1]*sobelY(0, 1) +
				imageIndex[i+1][j+1]*sobelY(1, 1) +
				imageIndex[i-1][j]*sobelY(-1, 1) +
				imageIndex[i][j]*sobelY(0, 0) +
				imageIndex[i+1][j]*sobelY(1, 0) +
				imageIndex[i-1][j-1]*sobelY(-1, -1) +
				imageIndex[i+1][j-1]*sobelY(1, -1) +
				imageIndex[i+1][j-1]*sobelY(1, -1)

			// Get magnitude of gradients for the two directions
			G := math.Sqrt(math.Pow(Gx, 2) + math.Pow(Gy, 2))

			// Magnitude of gradients go to the final image
			greyColour := color.Gray{uint8(G)}
			newImage.Set(i, j, greyColour)
		}
	}

	// Encode the image
	userPath, _ := user.Current()
	newFi, err := os.Create(userPath.HomeDir + "/canny-edge-detector/img/sobel.png")
	if err != nil {
		fmt.Printf("Failed to create %s: %s", newFi, err)
		panic(err.Error())
	}
	defer newFi.Close()
	png.Encode(newFi, newImage)
}
