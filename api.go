package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails"
)

// API => for handling the api keys needed for file upload and check
type API struct {
	filename string
	runtime  *wails.Runtime
	watcher  *fsnotify.Watcher
}

// NewAPI attemps to create a new instance of the struct
func NewAPI() (*API, error) {
	// create new api instance
	instance := &API{}

	return instance, nil
}

// Load loads the api from the config file
func (a *API) Load() {

}

// Save saves a new api to the config file
func (a *API) Save() {

}

// CreateFile creates an initial config file if
// no config file has been detected
func (a *API) CreateFile() {

}
