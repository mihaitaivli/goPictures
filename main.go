package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)


	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic("Something's wrong, exiting...")
	}
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	path := r.URL.Path
	if path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my wonderful site</h1>")
	} else if path == "/contact"{
		fmt.Fprint(w, "<p>To contact us please send an email to <a href=\"mailto:support@localhost.com\">support@localhost.com</a></p>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h3>Not found</h3>")
	}
}

func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my wonderful site</h1>")
}

func contact (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>To contact us please send an email to <a href=\"mailto:support@localhost.com\">support@localhost.com</a></p>")
}