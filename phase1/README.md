Longshoreman - Phase 1
======================

In this iteration of the project, your primary goal is to get a
simple web server up, using the power of Go and `net/http`.  It
must implement the following REST API:

    C: GET /ping
    S: 200 OK
       pong

The `/ping` interface is used by the front-end to ensure that it
is in fact talking to a Longshoreman REST API.  If the endpoint
doesn't properly respond with the 4 characters p-o-n-g, the
front-end code will balk.

    C: POST /run
       echo "desired output"
    S: 200 OK
       desired output

The `/run` interface is used by the front-end to submit commands.
In this phase, the commands will be run synchronously, and the
output will be returned in the immediate response.

The heavy lifting of actually running commands will be done by the
`sandbox` library ([available here][sandbox]).  You'll want to use
the **filefrog/sandbox** image:

Getting Started
---------------

To get you started, there is a `server.go` template that contains
a (very) short, but still functional webserver, running on port
8184.  You should be able to run it:

    go run server.go

And then visit http://localhost:8184 in your browser.

Next Steps
----------

Completing this phase will require that you do the following:

1. Dispatch URLs using `net/http`
2. Read the full body of an HTTP request
3. Install and use the `sandbox` library

Objectives
----------

Hopefully, this phase will teach you the basics of `net/http`, and
provide a hands-on experience writing a real Go web app!


[sandbox]: http://github.com/jhunt/sandbox
