package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails"
)

// api of VT
var virustotalAPI = "https://www.virustotal.com/api/v3/files"

// File => handles file to be uploaded for checking
type File struct {
	filename string
	runtime  *wails.Runtime
}

// ReturnResponse constructs a return response
type ReturnResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Response   string `json:"response"`
}
type response struct {
	data map[string]string
}

// NewFileHandler attemps to create an instance of File struct
func NewFileHandler() *File {
	// create a new instance
	instance := &File{}

	return instance
}

// SelectFileUpload pop-ups a dialog for file selection...
func (f *File) SelectFileUpload() string {
	filename := f.runtime.Dialog.SelectFile()

	if len(filename) > 0 {
		// set filename on the struct
		f.filename = filename

		// return the filename
		return filename
	}

	return ""
}

// FileUpload uploads the set filename to the api
// It returns the status code and the response string
// based from: https://gist.github.com/andrewmilson/19185aab2347f6ad29f5
func (f *File) FileUpload(filename, apiKey string) string {
	// open the file (change to f.filename)
	file, _ := os.Open(filename)
	defer file.Close()

	fmt.Println(filename)

	// create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// create new request
	req, _ := http.NewRequest("POST", "https://www.virustotal.com/api/v3/files", body)

	// set headers
	req.Header.Add("Content-Type", writer.FormDataContentType()) // multipart/form-data
	req.Header.Add("x-apikey", apiKey)                           // virustotal apikey

	// new client
	client := &http.Client{}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read content from resp
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	returnResp := &ReturnResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Response:   string(content),
	}

	// convert to string the response
	out, _ := json.Marshal(returnResp)

	return string(out)
}
