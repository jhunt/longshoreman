package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jhunt/sandbox"
	"github.com/pborman/uuid"
)

// We'll use this Job struct to keep track of the commands
// that we are running on behalf of our clients.
//
type Job struct {
	// For each Job, we need to track the Output of the command,
	// and whether or not it is still Running.
	//
	// Go ahead and set up those struct members...
}

func main() {
	// We'll use this map to keep track of our Jobs.
	jobs := make(map [string] *Job)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}

		b, _ := ioutil.ReadAll(r.Body)
		// In Phase 1, we just called sandbox.Run(), and waited for the
		// command to exit.  This time around, we want to create a Job...

		// ... generate a new UUID ...
		id := uuid.NewRandom()

		// and store the Job in the map for later reference

		// and call sanbox.Run() from a goroutine.
		go func() {
			// Anonymous functions can come in real handy when you're
			// spinning up goroutines...
		}()

		// Finally, our API design dictates that we print the UUID in
		// our response payload, so the client knows what it is.
	})

	// Finally, we need to finish the `GET /output/:uuid` handler.
	http.HandleFunc("/output/", func(w http.ResponseWriter, r *http.Request) {
		// First, check that this is a GET request.
		// Everything else is a 405 Method Not Allowed

		id := strings.Split(r.URL.Path, "/")[2]
		// Look up the job by its UUID.
		//
		// If it exists, and is still running, we need to return a
		// 204 No Content response, with no payload.
		//
		// If it exists, but has completed, return the output of the
		// job as a 200 OK

		// Otherwise, it's a 404 not found!
		w.WriteHeader(404)
		fmt.Fprintf(w, "uuid %s not found", id)
	})

	http.Handle("/", http.FileServer(http.Dir("static")))
	fmt.Printf("longshoreman up and running at http://localhost:8184\n")
	fmt.Printf("  go check your browser!\n")
	http.ListenAndServe(":8080", nil)
}
