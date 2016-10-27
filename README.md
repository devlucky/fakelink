# FakeLink

[![Build Status](https://travis-ci.org/devlucky/fakelink.svg?branch=master)](https://travis-ci.org/devlucky/fakelink)
[![Coverage Status](https://coveralls.io/repos/github/devlucky/fakelink/badge.svg)](https://coveralls.io/github/devlucky/fakelink)

FakeLink is a small backend that provides FLaaS (Fake Links as a Service), based on the [Open Graph protocol](http://ogp.me/#types)


## Usage (via Docker)

A fully automated development environment for this project can be set up with Docker, Rocker and Docker-Compose. Here's how:

* `bin/docker-build.sh` creates two images: 
   - __devlucky/fakelink-dev__, for development purposes
   - __devlucky/fakelink__, for production purposes (lightweight deployable Go image)
* `bin/docker-run.sh` runs the whole project, including its dependencies, and serves it on 8080
* `bin/docker-run.sh bash`, where _bash_ can be replaced by any other possible command, executes the command on the project's containers (including dependencies) in an interactive way.


## HTTP

The application exposes the following endpoints:

* `GET /random` Returns the HTML for a random, public link
* `GET /links/:slug` Returns the HTML for a particular link, identified by its slug
* `POST /links` Takes a _multipart/form-data_ payload with two keys:
    - a file "image", to upload
    - a field "json" with the following structure:

```
{
    "link": {
        "private": false,
        "values": {
            "title": "A title for my fake link",
            "description": "..."
            
            # Other OpenGraph fields. See src/templates package 
            # to understand the accepted values and they way 
            # they will be used
        }
    }
}
```
