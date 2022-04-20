package canny

import (
	"github.com/hhharrisonnn/canny-edge-detector/imgcreation"
	"image"
	"image/color"
)

func DoubleThreshold(highThresholdRatio float64, lowThresholdRatio float64) {
	imageIndex, imgWidth, imgHeight := imgcreation.GetGrey("./img/nonmaxsup.png")

	var strong float64 = 255
	var weak float64 = 100

	highThreshold := strong * highThresholdRatio
	lowThreshold := highThreshold * lowThresholdRatio

	newImage := image.NewGray((image.Rectangle{image.Point{1, 1}, image.Point{imgWidth - 1, imgHeight - 1}}))

	for j := 1; j < imgHeight-1; j++ {
		for i := 1; i < imgWidth-1; i++ {
			// 100% an edge
			if imageIndex[i][j] >= highThreshold {
				imageIndex[i][j] = strong
			}
			// Weak
			if (imageIndex[i][j] < highThreshold) && (imageIndex[i][j] > lowThreshold) {
				imageIndex[i][j] = weak
			}
			// Not an edge
			if imageIndex[i][j] <= lowThreshold {
				imageIndex[i][j] = 0
			}

			greyColour := color.Gray{uint8(imageIndex[i][j])}
			newImage.Set(i, j, greyColour)
		}
	}

	imgcreation.Encode(newImage, "doublethreshold")
}
