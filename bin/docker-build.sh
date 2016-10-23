#!/usr/bin/env bash

export DOCKER_REGISTRY=${DOCKER_REGISTRY:-""}
export DOCKER_IMAGE=${DOCKER_IMAGE:-"devlucky/fakelink"}
export DOCKER_TAG=${DOCKER_TAG:-"local"}

rocker build