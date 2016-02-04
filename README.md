Longshoreman
============

An exercise in writing real-world Go applications


- **Phase 1** - In which we set up a working HTTP REST API.
- **Phase 2** - In which we use goroutines for concurrency.
- **Phase 3** - In which we mitigate DoS via a worker pool.

HINTS are included for the stuck.

Getting Started
---------------

Before you can start hacking, you'll need to spin up the
Longshoreman development environment, a Vagrant box.

You'll need Vagrant installed.

    brew cask install virtualbox
    brew cask install vagrant

Then, spin up the Vagrant machine (from this directory)

    vagrant up

That may take a few minutes, as it installs a bunch of stuff
you'll need.  Oncce it's up, you can SSH into it:

    vagrant ssh

This will drop you into a shell as the `vagrant` user.  In your
home directory, you'll see a `go/` directory (your `$GOPATH`), and
multiple `phaseN` directories, one for each phase of this
exercise.

Finally, you'll need to install the following libraries, which are
dependencies of the project (just how this is done is part of your
training!):

- https://github.com/jhunt/sandbox
- https://github.com/pborman/uuid


Running the Demo
----------------

To see a complete implementation of Longshoreman in action, run
phase3's demo script:

    ./phase3/demo

then go visit http://localhost:8184 on your laptop.


Working Through The Exercises
-----------------------------

Each phase of the exercise is self-contained inside of its
directory, including the static assets for the front-end web
interface.  Inside each phase directory, you'll also find a
README.md file that contains instructions for the phase, and a
HINTS file for when you get stuck.

Happy Hacking!
