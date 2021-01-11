package main

import (
	"log"
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

// WailsInit => init runtime for files structs
func (f *File) WailsInit(runtime *wails.Runtime) error {
	// set runtime
	f.runtime = runtime

	return nil
}

//WailsInit => init runtime for API strutcts
func (a *API) WailsInit(runtime *wails.Runtime) error {
	// set runtime
	a.runtime = runtime
	a.logger = a.runtime.Log.New("VT-Check API file watcher...")

	a.filename = APICONFIGFILENAME

	// create a blank file if doesn't exist
	a.CreateFile()

	// start watching the api file and return it
	return a.StartWatcher()
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	// create instance of API
	api, err := NewAPI()
	if err != nil {
		log.Fatal(err)
	}

	// create instance of File
	file := NewFileHandler()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  950,
		Height: 600,
		Title:  "vt-check",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(api)
	app.Bind(file)
	app.Run()
}

// LogRes - logs
func LogRes(content string){
	fmt.Println(content)
}
