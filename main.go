package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandlerIndex) //membuat routing index root
	mux.HandleFunc("/about", handler.AboutPage)
	mux.HandleFunc("/profil", handler.ProfilePage)
	mux.HandleFunc("/product", handler.ProductHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) //link untuk asset agar bisa di gunakan di layout.html

	log.Println("Starting Port :8888") //memberikan pesan di dalam terminal

	err := http.ListenAndServe( "8888", mux) //perintah ini untuk menjalankan web server
	log.Fatal(err)

}
