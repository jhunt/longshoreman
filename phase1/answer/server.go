package main

import (
	"fmt"
	"github.com/jhunt/sandbox"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}

		b, _ := ioutil.ReadAll(r.Body)
		_, output, err := sandbox.Run(sandbox.Options{
			Image:   "filefrog/sandbox",
			Command: string(b),
		})
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "failed: %s\n", err)
			return
		}
		fmt.Fprintf(w, "%s", output)
	})

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
