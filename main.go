package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {

	templ, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
}

func asciiHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	txt := req.FormValue("text")
	banners := req.FormValue("banner")

	if txt == "" || banners == "" {
		http.Error(w, "Empty values", http.StatusBadRequest)
		return
	}

	bannerPath := fmt.Sprintf("banners/%s.txt", banners)
	loaded, err := LoadBanner(bannerPath)
	if err != nil {
		http.Error(w, "error loading banner", http.StatusNotFound)
		return
	}

	output := BuildArt(txt, loaded)
	fmt.Fprint(w, output)

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
