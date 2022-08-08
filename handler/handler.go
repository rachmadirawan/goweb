package handler

import "net/http"

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	w.Write([]byte(message))
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	message := "About Page"
	w.Write([]byte(message))
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	message := "Profile Page"
	w.Write([]byte(message))
}
