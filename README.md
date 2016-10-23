# FakeLink

FakeLink is a small backend that provides FLaaS (Fake Links as a Service), based on the [Open Graph protocol](http://ogp.me/#types)


## Usage (via Docker)

A fully automated development environment for this project can be set up with Docker, Rocker and Docker-Compose. Here's how:

* `bin/docker-build.sh` creates two images: 
   - __devlucky/fakelink-dev__, for development purposes
   - __devlucky/fakelink__, for production purposes (lightweight deployable Go image)
* `bin/docker-run.sh` runs the whole project, including its dependencies, and serves it on 8080
* `bin/docker-run.sh /bin/bash`, where /bin/bash can be replaced by any other possible command, executes the command on the project's containers (including dependencies) in an interactive way.
