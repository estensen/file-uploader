package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
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
	fmt.Println("Uploading File...")

	r.ParseMultipartForm(10 << 20) // 10 MB (n << x is n times 2, x times)
	file, handler, err := r.FormFile("selectedImage")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	filename := handler.Filename

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadPath := filepath.Join("uploads", filename)
	ioutil.WriteFile(uploadPath, fileBytes, 0644)
	fmt.Printf("Successfully Uploaded %+v\n", handler.Filename)
}
