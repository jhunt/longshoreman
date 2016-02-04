package main

import (
	"net/http"
)

type Job struct {
	Command string
	Running bool
	Output  string
}

// Our worker() function will be given a channel, via which it will
// receive Job pointers for processing...
func worker(in chan *Job) {
	for {
		// Receive a Job from the channel, run it, record the
		// output and mark it as completed.
	}
}

func main() {
	jobs := make(map[string]*Job)

	// We'll use this channel to communicate with out workers.
	ch := make(chan *Job)

	// Spin some worker goroutines

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

		// Now that we have our new job, we need to somehow communicate
		// the job details to one of our available workers

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
