package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var pesan = "Welcome"
	w.Write([]byte(pesan))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var pesan = "Halo world"
	w.Write([]byte(pesan))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/halo", handlerHello)
	http.HandleFunc("/halo/world", handlerHello)

	var address = "localhost:8080"
	fmt.Printf("server started at %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
