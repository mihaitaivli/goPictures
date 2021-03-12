package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandle)
	http.ListenAndServe(":3000", nil)
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my wonderful site</h1>")
}
