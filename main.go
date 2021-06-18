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

func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my wonderful site</h1>")
}

func contact (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>To contact us please send an email to <a href=\"mailto:support@localhost.com\">support@localhost.com</a></p>")
}