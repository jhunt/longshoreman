package main

import (
	"net/http"
)

func main() {

	/* start here */

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
