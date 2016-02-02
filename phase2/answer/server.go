package main

import (
	"fmt"
	"github.com/jhunt/sandbox"
	"github.com/pborman/uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type Job struct {
	Command string
	Running bool
	Output  string
}

func main() {
	jobs := make(map[string]*Job)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}

		b, _ := ioutil.ReadAll(r.Body)
		job := &Job{
			Command: string(b),
			Running: true,
		}
		id := uuid.NewRandom()
		jobs[id.String()] = job

		go func() {
			_, output, _ := sandbox.Run(sandbox.Options{
				Image:   "filefrog/sandbox",
				Command: string(job.Command),
			})
			job.Output = output
			job.Running = false
		}()

		fmt.Fprintf(w, "%s", id)
	})

	http.HandleFunc("/output/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}

		id := strings.Split(r.URL.Path, "/")[2]
		if job, ok := jobs[id]; ok {
			if job.Running {
				w.WriteHeader(204)
				return
			}

			fmt.Fprintf(w, "%s", job.Output)
			return
		}

		w.WriteHeader(404)
		fmt.Fprintf(w, "uuid %s not found", id)
	})

	http.Handle("/", http.FileServer(http.Dir("static")))
	fmt.Printf("longshoreman up and running at http://localhost:8184\n")
	fmt.Printf("  go check your browser!\n")
	http.ListenAndServe(":8080", nil)
}
