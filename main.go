package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", rootHandle)

	err := http.ListenAndServe(":3000", mux)
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
