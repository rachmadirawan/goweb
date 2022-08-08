package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	/*
		cobaHandler := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ini adalah handler Coba")) //function expression
		}
	*/

	mux.HandleFunc("/", handler.HandlerIndex) //membuat routing index root
	mux.HandleFunc("/about", handler.AboutPage)
	mux.HandleFunc("/profil", handler.ProfilePage)

	//mux.HandleFunc("/coba", cobaHandler)
	/*
		mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Page Test")) //anonymouse function
		})
	*/

	log.Println("Starting Port :8888") //memberikan pesan di dalam terminal

	err := http.ListenAndServe(":8888", mux) //perintah ini untuk menjalankan web server
	log.Fatal(err)

}
