package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello hell2!")
}
func main() {
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Single page:", r.URL.String())
	})
	http.HandleFunc("/pages/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Login page:", r.URL.String())
	})
	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Multiple pages:", r.URL.String())
	})
	http.HandleFunc("/", handler)

	fmt.Println("staring server at")
	http.ListenAndServe(":8081", nil)

}
