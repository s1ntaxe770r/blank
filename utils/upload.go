package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
)

type Undelivered struct {
	Status string `json:"status"`
	Err    error  `json:"err"`
}

type Parcel struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

func UploadFile(filename string, server string, creds Credentials) (error, string) {

	//Open file
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(color.RedString(err.Error()))
		os.Exit(1)
	}

	// prepare request
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return err, "error writing to buffer"
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		return err, " "
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err, " "
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(server+"/push", contentType, bodyBuf)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()
	var uploadedfile Parcel
	json.NewDecoder(resp.Body).Decode(&uploadedfile)
	parsedurl, _ := url.Parse(server)
	filelocation := parsedurl.Scheme + "://" + uploadedfile.URL
	return nil, filelocation
}
