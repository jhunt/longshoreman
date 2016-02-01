Longshoreman - Phase 2
======================

Now it's time to extend our webserver to utilize goroutines in a
worker pool architecture.  Rather than have a `POST /run` block
waiting for the command to exit, we'll submit a job to our
workers, and return an identifier to the client.  The client can
then take that token and ask for the status of the command via a
new `GET /output` endpoint.

So our new API design looks like this:

    C: GET /ping
    S: 200 OK
       pong

(this is unchanged)

    C: POST /run
       echo "desired output"
    S: 200 OK
       1db89b70-d0f9-4632-9cec-3b51e0c5b560

The `/run` interface now returns a UUID that represents the job
executing the command.

    C: GET /output/1db89b70-d0f9-4632-9cec-3b51e0c5b560
    S: 204 No Content

       ... time passes ...

    C: GET /output/1db89b70-d0f9-4632-9cec-3b51e0c5b560
    S: 200 OK
       desired output

The new `/output` interface takes a UUID (in the URL path) and
returns either a 204 No Content if the job is still executing, or
a 200 OK (with the output) if the job has completed.

Getting Started
---------------

You are free to start with the `server.go` template, but you'll
probably be better off copying your implementation from phase 1
into this directory.  Running it is the same as before:

    go run server.go

And then visit http://localhost:8184 in your browser.

Next Steps
----------

Completing this phase will require that you do the following:

1. Design a custom, struct-based type
2. Create new goroutines
3. Pass messages around via channels

Objectives
----------

Hopefully, by completing this phase, you'll gain an appreciation
for goroutines and concurrency.
