package canny

import (
	"github.com/hhharrisonnn/canny-edge-detector/imgcreation"
	"image"
	"image/color"
)

func Hysteresis() {
	imageIndex, imgWidth, imgHeight := imgcreation.GetGrey("./img/doublethreshold.png")

	var strong float64 = 255
	var weak float64 = 100

	newImage := image.NewGray((image.Rectangle{image.Point{1, 1}, image.Point{imgWidth - 1, imgHeight - 1}}))

	for j := 1; j < imgHeight-1; j++ {
		for i := 1; i < imgWidth-1; i++ {
			// If weak pixel has a strong pixel around it, it will also be strong
			if imageIndex[i][j] == weak {
				if (imageIndex[i-1][j+1] == strong) || (imageIndex[i][j+1] == strong) || (imageIndex[i+1][j+1] == strong) ||
					(imageIndex[i-1][j] == strong) || (imageIndex[i][j] == strong) || (imageIndex[i+1][j] == strong) ||
					(imageIndex[i-1][j-1] == strong) || (imageIndex[i][j-1] == strong) || (imageIndex[i+1][j-1] == strong) {
					imageIndex[i][j] = strong
				} else {
					imageIndex[i][j] = 0
				}
			}
			greyColour := color.Gray{uint8(imageIndex[i][j])}
			newImage.Set(i, j, greyColour)
		}
	}

	imgcreation.Encode(newImage, "hysteresis")
}
