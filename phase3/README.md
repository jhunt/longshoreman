Longshoreman - Phase 2
======================

We've got concurrency down pat, but we have a potential denial of
service waiting in the wings; there is no upper bound on the
number of goroutines (and by extension, containers) our web API
creates.  This is bad.

We'll fix that in this phase by using a fixed number of worker
goroutines (pre-spun) and use a channel to communicate with them.

The API design won't change, but we do need to move the bulk of
our `/run` handler into a separate function that can be run from
multiple goroutines concurrently.  We'll control our capacity by
only creating a fixed number of such goroutines.

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

1. Create a channel
2. Hook multiple goroutines up to that channel
3. Rewire the `/run` handler

Objectives
----------

Hopefully, by completing this phase, you'll gain an even deeper
appreciation for Go's concurrency primitives (channels and
goroutines).
