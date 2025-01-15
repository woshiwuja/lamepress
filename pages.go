package main

import (
	"fmt"
	"net/http"
	"os"
)

func GetPagesTitles(path string) (titles []string, err error) {
	titlesEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, title := range titlesEntries {
		titles = append(titles, title.Name())
	}
	fmt.Println(titles)
	return titles, nil
}

func RegisterPagesPath(pages []string, mux *http.ServeMux) {
	for _, page := range pages {
		mux.HandleFunc("/"+page, pageHandler(page))
	}
}

func pageHandler(page string) http.HandlerFunc {
	fmt.Println("got request for page:", page)
	pageContent, err := os.ReadFile("./static/pages/" + page + "/page.html")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) { w.Write(pageContent) }
}
