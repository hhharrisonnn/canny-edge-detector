package canny

import (
	"github.com/hhharrisonnn/canny-edge-detector/imgcreation"
	"image"
	"image/color"
)

func NonMaxSuppression() {
	_, theta := SobelConvolution()

	imageIndex, imgWidth, imgHeight := imgcreation.GetGrey("./img/sobel.png")

	newImage := image.NewGray((image.Rectangle{image.Point{1, 1}, image.Point{imgWidth - 1, imgHeight - 1}}))

	for j := 1; j < imgHeight-1; j++ {
		for i := 1; i < imgWidth-1; i++ {
			angle := theta[i][j]

			var a int
			var b int

			if (0 <= angle && angle < 180) || (157.5 <= angle && angle < 180) { // 0 degrees
				a = int(imageIndex[i][j+1])
				b = int(imageIndex[i][j-1])
			} else if 22.5 <= angle && angle < 67.5 { // 45 degrees
				a = int(imageIndex[i+1][j-1])
				b = int(imageIndex[i-1][j+1])
			} else if 67.5 <= angle && angle < 112.5 { // 90 degrees
				a = int(imageIndex[i+1][j])
				b = int(imageIndex[i-1][j])
			} else if 112.5 <= angle && angle < 157.5 { // 135 degrees
				a = int(imageIndex[i-1][j-1])
				b = int(imageIndex[i+1][j+1])
			}

			if (int(imageIndex[i][j]) <= a) && (int(imageIndex[i][j])) <= b {
				// If the current pixel is smaller than the two other pixels, it is suppressed
				imageIndex[i][j] = 0
			} else {
				// Otherwise, it is kept
				imageIndex[i][j] *= 1
			}

			greyColour := color.Gray{uint8(imageIndex[i][j])}
			newImage.Set(i, j, greyColour)
		}
	}

	imgcreation.Encode(newImage, "nonmaxsup")
}
