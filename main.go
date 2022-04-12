package main

import (
	"fmt" // Standard package for Go formatting
	greyscale "github.com/hhharrisonnn/canny-edge-detector/src"
	"io"
	"io/ioutil" // Package for reading/writing files
	"log"       // For logging errors
	"net/http"  // Anything HTTP related - start web servers, handling requests
	"os"
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

	greyscale.Greyscale() // Activate Greyscale function after receiving the image
}

// Function to check if a directory is empty
func emptyDir(dirName string) bool {
	file, err := os.Open(dirName)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Readdirnames(1)
	if err == io.EOF {
		return true
	}
	return false
}

func main() {
	if emptyDir("./img/") == true { // If directory is empty
		http.HandleFunc("/upload", uploadFile)
	}
	if emptyDir("./img/") == false { // If directory is not empty
		greyscale.Greyscale() // Activates Greyscale function
	}

	// Starts simple web server
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
