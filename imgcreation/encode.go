package imgcreation

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/user"
)

func Encode(image *image.Gray, imgName string) {
	// Encode the image
	userPath, _ := user.Current()
	newFi, err := os.Create(userPath.HomeDir + "/canny-edge-detector/img/" + imgName + ".png")
	if err != nil {
		fmt.Printf("Failed to create %s: %s", newFi, err)
		panic(err.Error())
	}
	defer newFi.Close()
	png.Encode(newFi, image)
}
