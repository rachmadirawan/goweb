package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	//message := "Hello World"
	//w.Write([]byte(message))

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Non Found File", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":    "Belajar Golang",
		"containt": "Golang WEB DTS PROA 3",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Non Found File", http.StatusInternalServerError)
		return
	}
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	message := "About Page"
	w.Write([]byte(message))
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	message := "Profile Page"
	w.Write([]byte(message))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"ID":      "1",
		"Product": "Mikrotik RB 450G x4",
		"Price":   1800000,
		"Stock":   5,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Page Not Found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Page Not Found", http.StatusInternalServerError)
		return
	}
}
