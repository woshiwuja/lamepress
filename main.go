package main

import (
	"bytes"
	"net/http"
	"os"
	"time"
)

const SRV_PORT string = ":8080"

func getIndex(w http.ResponseWriter, r *http.Request) {
	indexFile, err := os.ReadFile("./static/index.html")
	if err != nil {
		panic(err)
	}
	println(indexFile)
	content := bytes.NewReader(indexFile)
	if err != nil {
		panic(err)
	}
	http.ServeContent(w, r, "index.html", time.Now(), content)
}
func getCss(w http.ResponseWriter, r *http.Request) {
	cssFile, err := os.ReadFile("./static/css/style.css")
	if err != nil {
		panic(err)
	}
	content := bytes.NewReader(cssFile)
	if err != nil {
		panic(err)
	}
	http.ServeContent(w, r, "style.css", time.Now(), content)
}

func main() {
	mux := http.NewServeMux()
	pages, errDir := GetDirPages("./static/pages/")
	if errDir != nil {
		println("dir not found")
	}

	mux.HandleFunc("/", getIndex)
	mux.HandleFunc("/index", getIndex)
	mux.HandleFunc("/index.html", getIndex)
	mux.HandleFunc("/css/style.css", getCss)
	RegisterPagesPath(pages, mux)
	err := http.ListenAndServe(SRV_PORT, mux)
	if err != nil {
		panic(err)
	}
}
