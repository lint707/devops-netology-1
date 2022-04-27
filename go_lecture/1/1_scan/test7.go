package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Name:", h.Name, "URL:", r.URL.String())
}
func (h *Handler) MyAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "MyAPI Name:", h.Name, "URL:", r.URL.String())
}
func main() {
	testHandler := &Handler{Name: "test"}
	http.Handle("/test/", testHandler)

	rootHandler := &Handler{Name: "root"}
	http.Handle("/", rootHandler)

	apiHandler := &Handler{Name: "API"}
	http.HandleFunc("/api/", apiHandler.MyAPI)

	// uh:=&Handler{Name: "User"}
	// http.HandleFunc("/api/user/login", uh.Login)
	// http.HandleFunc("/api/user/logout", uh.Logout)
	// http.HandleFunc("/api/user/signup", uh.Signup)

	fmt.Println("staring server at 8084")
	http.ListenAndServe(":8084", nil)

}
