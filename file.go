package main

import (
	"fmt"

	"github.com/wailsapp/wails"
)

// File => handles file to be uploaded for checking
type File struct {
	filename string
	runtime  *wails.Runtime
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
		fmt.Println(filename)
		return filename
	}

	return ""
}
