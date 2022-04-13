package main

import (
	"fmt" // Standard package for Go formatting
	greyscale "github.com/hhharrisonnn/canny-edge-detector/src"
	"io/ioutil" // Package for reading/writing files
	"log"       // For logging errors
	"net/http"  // Anything HTTP related - start web servers, handling requests
	"os"
	"os/user"
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

// Receive input from the menu and return an image
func postImage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	menuValue := r.FormValue("steps")

	if menuValue == "Greyscale" {
		path := "img/grayscale.png"
		fmt.Fprintf(w, "<h1>Greyscale image</h1>")
		fmt.Fprintf(w, "<img src=%q>", path)
	}
}

func main() {
	// If there's anything in the img directory, remove it
	files, _ := ioutil.ReadDir("./img/")
	userPath, _ := user.Current()
	if len(files) != 0 {
		for _, f := range files {
			os.Remove(userPath.HomeDir + "/canny-edge-detector/img/" + f.Name())
		}
	}

	if len(files) == 0 {
		http.HandleFunc("/upload", uploadFile)
	}

	// Handler function for request
	http.HandleFunc("/", postImage)
	// Handle requests for 'img' directory
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	// Starts simple web server
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
