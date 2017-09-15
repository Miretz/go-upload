package main

import (
	"os"
	"io"
	"net/http"
)

func UploadFileIntoFS(r *http.Request) (string, error) {

	//parse form and get file
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		return "", err
	}
	defer file.Close()

	//create a new file and open it
	newFile, err := os.OpenFile(uploadDir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	//copy bytes to the new file
	io.Copy(newFile, file)

	return handler.Filename, nil
}

func DeleteFromFS(path string){
	os.Remove("." + path)
}
