package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails"
)

// APICONFIGFILENAME -> default filename where to store the user's VirusTotal API_KEY
var APICONFIGFILENAME = "vtcheck-api.json"

// API => for handling the api keys needed for file upload and check
type API struct {
	filename string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
	watcher  *fsnotify.Watcher
}

// NewAPI attemps to create a new instance of the struct
func NewAPI() (*API, error) {
	// create new api instance
	instance := &API{}

	// set the filename for the instance
	instance.filename = APICONFIGFILENAME

	// create the config file
	// if it doesn't exist
	// on load / start-up
	instance.CreateFile()

	// return the instance api
	return instance, nil
}

// Load loads the api from the config file
func (a *API) Load() (string, error) {
	content, err := ioutil.ReadFile(a.filename)
	if err != nil {
		err = fmt.Errorf("Unable to load the config file: %s", a.filename)
	}

	return string(content), err
}

// Save saves a new api to the config file
func (a *API) Save(api string) error {
	return ioutil.WriteFile(a.filename, []byte(api), 0755)
}

// CreateFile creates an initial config file if
// no config file has been detected
func (a *API) CreateFile() {
	if _, err := os.Stat(a.filename); os.IsNotExist(err) {
		// create a blank config file
		ioutil.WriteFile(a.filename, []byte("{}"), 0755)
	}
}

// StartWatcher => watches file changes
// made into the main config file
func (a *API) StartWatcher() error {
	a.logger.Info("Starting Watcher ... ")

	watcher, err := fsnotify.NewWatcher()
	a.watcher = watcher
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					a.logger.Infof("modified file: %s", event.Name)
					a.runtime.Events.Emit("datamodified") // emit datamodified, the frontend will reload after getting this message
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				a.logger.Error(err.Error())
			}
		}
	}()

	// add the api file to the watcher
	err = a.watcher.Add(a.filename)
	if err != nil {
		return err
	}

	return nil
}
