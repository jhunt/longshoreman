package main

import (
	// You can put other (potentially missing) imports here
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// The front-end expects `GET /ping` to return the text "pong".
		// It uses this response payload to make sure that the backend API
		// is functional.
		//
		// To make that happen, take a look at the `fmt` package documentation
		// (available online at https://golang.org/pkg/fmt/), and keep in mind
		// that w (an http.ResponseWriter) implements the io.Writer interface.
		//
	})

	// You'll need to add a second HandleFunc() call to handle requests to
	// the `/run` endpoint.  A few notes on that implementation:
	//
	// 1) You should check to make sure that the request is a POST; check
	//    the documentation for the http.Request type for ideas.
	//
	// 2) You need some way of reading the body of the request.  There's a
	//    package called `io/ioutil` that may have something of interest...
	//
	// 3) Having verified the request type and extracted the request body,
	//    your next step is to call sandbox.Run().  Don't forget to add the
	//    "github.com/jhunt/sandbox" to your imports!
	//

	http.Handle("/", http.FileServer(http.Dir("static")))
	fmt.Printf("longshoreman up and running at http://localhost:8184\n")
	fmt.Printf("  go check your browser!\n")
	http.ListenAndServe(":8080", nil)
}
