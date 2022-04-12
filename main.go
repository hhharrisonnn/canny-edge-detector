//package main
//
//import (
//	"fmt"   // Standard package for formatting
//	"image" // Package for reading images
//	"image/color"
//	"image/png"
//	"os" // Package for handling files
//)
//
//func Greyscale() {
//	// Open image file
//	fi := "./img/test.png"
//	inputImg, err := os.Open(fi)
//	// If there is an error that exists then print it
//	if err != nil {
//		fmt.Printf("Failed to open %s: %s", fi, err)
//		panic(err.Error())
//	}
//	defer inputImg.Close() // Close file either when the function is finished or there's a panic
//
//	// Decodes image to get its values
//	imgData, _, err := image.Decode(inputImg)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// Greyscale function
//	// Get dimensions of image
//	imgBound := imgData.Bounds()
//	imgWidth, imgHeight := imgBound.Max.X, imgBound.Max.Y
//	// Return a grey image with the given bounds
//	greyscale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{imgWidth, imgHeight}})
//	// Loop over every pixel in the image
//	for x := 0; x < imgWidth; x++ {
//		for y := 0; y < imgHeight; y++ {
//			imgColour := imgData.At(x, y)                                               // Convert values at a given point to 16-bit per channel RGBA
//			r, g, b, _ := imgColour.RGBA()                                              // Stores the separate values of RGBA
//			Y := uint16((0.3 * float64(r)) + (0.59 * float64(g)) + (0.11 * float64(b))) // Weighted values of RGB added
//			greyColour := color.Gray{uint8(Y >> 8)}                                     // 8 bits are discarded from the 16-bit weighted values
//			greyscale.Set(x, y, greyColour)                                             // Set pixels as the new grey colour
//		}
//	}
//
//	// Encode the image
//	newFi, err := os.Create("greyscale.png")
//	if err != nil {
//		fmt.Printf("Failed to create %s: %s", newFi, err)
//		panic(err.Error())
//	}
//	defer newFi.Close()
//	png.Encode(newFi, greyscale)
//}

package main

import (
	"fmt"       // Standard package for Go formatting
	"io/ioutil" // Package for reading/writing files
	"log"       // For logging errors
	"net/http"  // Anything HTTP related - start web servers, handling requests
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)                   // Declares maximum file size
	file, fileHeader, err := r.FormFile("inputFile") // Returns first file from 'myFile' located in index.html
	if err != nil {
		fmt.Println("Error retrieving form from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", fileHeader.Filename) // Get file name
	fmt.Printf("File size: %+v\n", fileHeader.Size)         // Get file size
	fmt.Printf("Header: %+v\n", fileHeader.Header)          // Get file header

	// Create a temporary file in the 'img' directory, generating a random file name with the *
	tempFile, err := ioutil.TempFile("img", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// Puts all data of the uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes) // Write the byte array into the temporary file

	reDir := "javascript:history.back()"
	fmt.Fprintf(w, "<a href=%q>Click on this to see the steps!</a>", reDir)
}

// Starts simple web server
func main() {
	http.HandleFunc("/upload", uploadFile)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
