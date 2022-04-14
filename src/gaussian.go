package canny

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/user"
)

func GaussianKernel(i, j int8, sigma float64) float64 {
	normalFunc := 1 / (2 * math.Pi * math.Pow(sigma, 2)) // Normal function

	// 5x5 matrices used for the calculation of the kernel
	xMat := [5][5]float64{
		{-2, -2, -2, -2, -2},
		{-1, -1, -1, -1, -1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
	} // Matrix for the x direction

	yMat := [5][5]float64{
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
		{-2, -1, 0, 1, 2},
	} // Matrix for the y direction

	xyMat := [5][5]float64{} // Initialise slice for result of exponential

	// Gaussian kernel calculation
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			xMatSquare := math.Pow(xMat[i][j], 2) // Square elements from x matrix

			yMatSquare := math.Pow(yMat[i][j], 2) // Square elements from y matrix

			xyMatDiv := -((xMatSquare + yMatSquare) / (2 * math.Pow(sigma, 2))) // This is what the exponential will be raised to

			xyMatExp := math.Exp(xyMatDiv) // Exponential of xyMatDiv

			xyMatFinal := xyMatExp * normalFunc // Normalise

			xyMat[i][j] = xyMatFinal // Put values into matrix
		}
	}

	return xyMat[i+2][j+2] // Make middle of the kernel (0, 0)
}

func GaussianConvolution(sigma float64) {
	// Open greyscale.png
	fi := "./img/grayscale.png"
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

	// Stores final image values
	newImage := image.NewGray((image.Rectangle{image.Point{0, 0}, image.Point{imgWidth, imgHeight}}))
	// Iterate over pixels, get surrounding pixels in 5x5 area, and convolve them
	for j := 2; j < imgHeight-2; j++ {
		for i := 2; i < imgWidth-2; i++ {
			// Convolution
			result := imageIndex[i-2][j+2]*GaussianKernel(-2, 2, sigma) +
				imageIndex[i-1][j+2]*GaussianKernel(-1, 2, sigma) +
				imageIndex[i][j+2]*GaussianKernel(0, 2, sigma) +
				imageIndex[i+1][j+2]*GaussianKernel(1, 2, sigma) +
				imageIndex[i+2][j+2]*GaussianKernel(2, 2, sigma) +
				imageIndex[i-2][j+1]*GaussianKernel(-2, 1, sigma) +
				imageIndex[i-1][j+1]*GaussianKernel(-1, 1, sigma) +
				imageIndex[i][j+1]*GaussianKernel(0, 1, sigma) +
				imageIndex[i+1][j+1]*GaussianKernel(1, 1, sigma) +
				imageIndex[i+2][j+1]*GaussianKernel(2, 1, sigma) +
				imageIndex[i-2][j]*GaussianKernel(-2, 0, sigma) +
				imageIndex[i-1][j]*GaussianKernel(-1, 0, sigma) +
				imageIndex[i][j]*GaussianKernel(0, 0, sigma) +
				imageIndex[i+1][j]*GaussianKernel(1, 0, sigma) +
				imageIndex[i+2][j]*GaussianKernel(2, 0, sigma) +
				imageIndex[i-2][j-1]*GaussianKernel(-2, -1, sigma) +
				imageIndex[i-1][j-1]*GaussianKernel(-1, -1, sigma) +
				imageIndex[i][j-1]*GaussianKernel(0, -1, sigma) +
				imageIndex[i+1][j-1]*GaussianKernel(1, -1, sigma) +
				imageIndex[i+2][j-1]*GaussianKernel(2, -1, sigma) +
				imageIndex[i-2][j-2]*GaussianKernel(-2, -2, sigma) +
				imageIndex[i-1][j-2]*GaussianKernel(-1, -2, sigma) +
				imageIndex[i][j-2]*GaussianKernel(0, -2, sigma) +
				imageIndex[i+1][j-2]*GaussianKernel(1, -2, sigma) +
				imageIndex[i+2][j-2]*GaussianKernel(2, -2, sigma)
			greyColour := color.Gray{uint8(result)}
			newImage.Set(i, j, greyColour)
		}
	}

	// Encode the image
	userPath, _ := user.Current()
	newFi, err := os.Create(userPath.HomeDir + "/canny-edge-detector/img/gaussianBlur.png")
	if err != nil {
		fmt.Printf("Failed to create %s: %s", newFi, err)
		panic(err.Error())
	}
	defer newFi.Close()
	png.Encode(newFi, newImage)
}