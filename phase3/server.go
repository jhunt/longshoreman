package main

import (
	"net/http"
)

func main() {

	/* start here */

	http.Handle("/", http.FileServer(http.Dir("static")))
	fmt.Printf("longshoreman up and running at http://localhost:8184\n")
	fmt.Printf("  go check your browser!\n")
	http.ListenAndServe(":8080", nil)
}
