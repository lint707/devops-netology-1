package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL:", r.URL.String())
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// muxUser := http.NewServeMux()
	// muxUser.HandleFunc("/", handler)
	// mux.Handle("/users", muxUser)

	//http.ListenAndServe(":8082",nil) -> http.DefaultServerMux + http.DefaultServer
	//http.ListenAndServe(":8082",mux) -> mux + http.DefaultServer

	server := http.Server{
		Addr:         ":8082",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("staring server at: 8082")
	server.ListenAndServe()

}
