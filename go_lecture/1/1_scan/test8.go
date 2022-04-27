package main

import (
	"context"
	"fmt"
	"net/http"
)

// type Handler struct {
// 	Name string
//}

func runServer(addr string) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Addr:", addr, "URL:", r.URL.String())
	})

	mux.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		newAddr := ":" + r.FormValue("port")
		go runServer(newAddr)
		fmt.Fprintln(w, "Start new at", newAddr)

	})

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		server.Shutdown(context.Background())
	})

	fmt.Println("staring server at:", addr)
	server.ListenAndServe()
}

func main() {
	go runServer(":8081")
	runServer(":8080")
}
