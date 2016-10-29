#!/usr/bin/env bash

export DOCKER_GO_VERSION=${DOCKER_GO_VERSION:-"1.7"}
export DOCKER_REGISTRY=${DOCKER_REGISTRY:-""}
export DOCKER_IMAGE=${DOCKER_IMAGE:-"devlucky/fakelink"}
export DOCKER_ENVIRONMENT=${DOCKER_ENVIRONMENT:-""}
export DOCKER_TAG=${DOCKER_TAG:-"local"}

rocker build