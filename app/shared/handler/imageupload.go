package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
//UploadFile function return name of image
func UploadFile(w http.ResponseWriter, r *http.Request) (image string) {
	//fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		image = "nofile"
		return image
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	// pattern := handler.Filename
	tempFile, err := ioutil.TempFile("static/image", "*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	//fmt.Fprintf(w, "Successfully Uploaded File\n")
	image = tempFile.Name()
	image = "../../" + image[0:(len(image) - 4)]

	return image
}
	

//DeleteImage Function delete old avatar
func DeleteImage(path string) {
	var err = os.Remove(path)
	if isError(err) { 
		return 
	}
}
