package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	setupRoutes()
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Running on port 8080")
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading Image\n")

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("selectedImage")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
