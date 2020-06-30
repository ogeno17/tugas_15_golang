package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar Guys?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i := 1; i <= 100; i++ {
			fmt.Fprintln(w, i)
		}
	})

	http.HandleFunc("/index", index)
	fmt.Println("Start web server")
	http.ListenAndServe(":8080", nil)
}
