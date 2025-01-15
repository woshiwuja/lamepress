package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

func GetDirPages(path string) (pages []string, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		println(file.Name())
		page := file.Name()
		pages = append(pages, page)
	}
	println(pages)
	return pages, nil
}

func RegisterPagesPath(pages []string, mux *http.ServeMux) {
	fmt.Println("got request for page")
	for _, page := range pages {
		pageFunc := func(w http.ResponseWriter, r *http.Request) {
			pageFile, err := os.ReadFile("./static/pages/" + page)
			if err != nil {
				panic(err)
			}
			content := bytes.NewReader(pageFile)
			if err != nil {
				panic(err)
			}
			http.ServeContent(w, r, page, time.Now(), content)
		}
		mux.HandleFunc("/pages/"+page, pageFunc)
	}
}
