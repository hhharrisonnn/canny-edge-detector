package canny

import "math"

func GaussianKernel(sigma float64) [5][5]float64 {
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
	return xyMat
}
