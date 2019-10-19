package main

import (
	"log"
	"net/http"

	"github.com/spf13/afero"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Llongfile)
	log.Println("Hello World: 16")
}

func main() {
	// native os backend
	osFs := afero.NewOsFs()
	setting, err := osFs.Open("setting.toml")
	if err != nil {
		log.Panicln(err)
	}
	defer setting.Close()

	// wrap any of the existing backends
	httpFs := afero.NewHttpFs(osFs)
	fileserver := http.FileServer(httpFs.Dir("."))
	http.Handle("/", fileserver)
	_ = http.ListenAndServe(":8080", nil)
	for {}
}
